@public-tx
Feature: Send transfer transaction
  As an external developer
  I want to process a single transfer transaction

  Background:
    Given I have the following tenants
      | alias   | tenantID        |
      | tenant1 | {{random.uuid}} |
    And I register the following chains
      | alias | Name                | URLs                       | Headers.Authorization    |
      | besu  | besu-{{scenarioID}} | {{global.nodes.besu.URLs}} | Bearer {{tenant1.token}} |
      | geth  | geth-{{scenarioID}} | {{global.nodes.geth.URLs}} | Bearer {{tenant1.token}} |
    And I have created the following accounts
      | alias    | ID              | Headers.Authorization    |
      | account1 | {{random.uuid}} | Bearer {{tenant1.token}} |
      | account2 | {{random.uuid}} | Bearer {{tenant1.token}} |
    Given I sign the following transactions
      | alias     | ID              | Value              | Gas   | To           | privateKey                                 | ChainUUID     | Headers.Authorization    |
      | txFaucet1 | {{random.uuid}} | 100000000000000000 | 21000 | {{account1}} | {{global.nodes.besu.fundedPrivateKeys[0]}} | {{besu.UUID}} | Bearer {{tenant1.token}} |
      | txFaucet2 | {{random.uuid}} | 100000000000000000 | 21000 | {{account2}} | {{global.nodes.geth.fundedPrivateKeys[0]}} | {{geth.UUID}} | Bearer {{tenant1.token}} |
    Then I track the following envelopes
      | ID               |
      | {{txFaucet1.ID}} |
      | {{txFaucet2.ID}} |
    Given I set the headers
      | Key           | Value                    |
      | Authorization | Bearer {{tenant1.token}} |
      | Content-Type  | application/json         |
    When I send "POST" request to "{{global.tx-scheduler}}/transactions/send-raw" with json:
  """
{
    "chain": "besu-{{scenarioID}}",
    "params": {
        "raw": "{{txFaucet1.Raw}}"
    },
    "labels": {
    	"scenario.id": "{{scenarioID}}",
    	"id": "{{txFaucet1.ID}}"
    }
}
      """
    Then the response code should be 202
    When I send "POST" request to "{{global.tx-scheduler}}/transactions/send-raw" with json:
  """
{
    "chain": "geth-{{scenarioID}}",
    "params": {
        "raw": "{{txFaucet2.Raw}}"
    },
    "labels": {
    	"scenario.id": "{{scenarioID}}",
    	"id": "{{txFaucet2.ID}}"
    }
}
      """
    Then the response code should be 202
    Then Envelopes should be in topic "tx.decoded"

  Scenario: Send transfer transaction
    Given I register the following alias
      | alias | value              |
      | to1   | {{random.account}} |
      | to2   | {{random.account}} |
    When I send envelopes to topic "tx.crafter"
      | ID              | ChainName           | From         | Gas   | To      | Value     | Headers.Authorization    |
      | {{random.uuid}} | besu-{{scenarioID}} | {{account1}} | 21000 | {{to1}} | 500000000 | Bearer {{tenant1.token}} |
      | {{random.uuid}} | geth-{{scenarioID}} | {{account2}} | 21000 | {{to2}} | 400000000 | Bearer {{tenant1.token}} |
    Then Envelopes should be in topic "tx.crafter"
    Then Envelopes should be in topic "tx.signer"
    And Envelopes should have the following fields
      | Nonce |
      | 0     |
      | 0     |
    Then Envelopes should be in topic "tx.sender"
    And Envelopes should have the following fields
      | Raw | TxHash |
      | ~   | ~      |
      | ~   | ~      |
    Then Envelopes should be in topic "tx.decoded"
    And Envelopes should have the following fields
      | Receipt.Status |
      | 1              |
      | 1              |
    When I send "POST" request to "{{global.chain-registry}}/{{besu.UUID}}" with json:
      """
      {
        "jsonrpc": "2.0",
        "method": "eth_getBalance",
        "params": [
          "{{to1}}",
          "latest"
        ],
        "id": 1
      }
      """
    Then the response code should be 200
    And Response should have the following fields
      | result     |
      | 0x1dcd6500 |
    When I send "POST" request to "{{global.chain-registry}}/{{geth.UUID}}" with json:
      """
      {
        "jsonrpc": "2.0",
        "method": "eth_getBalance",
        "params": [
          "{{to2}}",
          "latest"
        ],
        "id": 1
      }
      """
    Then the response code should be 200
    And Response should have the following fields
      | result     |
      | 0x17d78400 |
