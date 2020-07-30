@tx-scheduler
Feature: Transaction Scheduler Idempotency
  As an external developer
  I want to use idempotency-key to interact with the transaction scheduler API

  Background:
    Given I have the following tenants
      | alias   | tenantID        |
      | tenant1 | {{random.uuid}} |
    Then I register the following contracts
      | name        | artifacts        | Headers.Authorization    |
      | SimpleToken | SimpleToken.json | Bearer {{tenant1.token}} |
    Then I register the following chains
      | alias  | Name                  | URLs                         | Headers.Authorization    |
      | besu_1 | besu_1-{{scenarioID}} | {{global.nodes.besu_1.URLs}} | Bearer {{tenant1.token}} |
      | besu_2 | besu_2-{{scenarioID}} | {{global.nodes.besu_1.URLs}} | Bearer {{tenant1.token}} |
    And I have created the following accounts
      | alias    | ID              | Headers.Authorization    |
      | account1 | {{random.uuid}} | Bearer {{tenant1.token}} |
      | account2 | {{random.uuid}} | Bearer {{tenant1.token}} |

  Scenario: Send twice same transaction using same idempotency key
    Given I register the following alias
      | alias          | value           |
      | idempotencykey | {{random.uuid}} |
      | deployTxOneID  | {{random.uuid}} |
      | deployTxTwoID  | {{random.uuid}} |
    Then I track the following envelopes
      | ID                |
      | {{deployTxOneID}} |
    Then  I set the headers
      | Key               | Value                    |
      | Authorization     | Bearer {{tenant1.token}} |
      | X-Idempotency-Key | {{idempotencykey}}       |
    When I send "POST" request to "{{global.tx-scheduler}}/transactions/deploy-contract" with json:
  """
{
    "chain": "besu_1-{{scenarioID}}",
    "params": {
        "from": "{{account1}}",
        "contractName": "SimpleToken"
    },
    "labels": {
    	"scenario.id": "{{scenarioID}}",
    	"id": "{{deployTxOneID}}"
    }
}
      """
    Then the response code should be 202
    Then  I set the headers
      | Key               | Value                    |
      | Authorization     | Bearer {{tenant1.token}} |
      | X-Idempotency-Key | {{idempotencykey}}       |
    When I send "POST" request to "{{global.tx-scheduler}}/transactions/deploy-contract" with json:
  """
{
    "chain": "besu_1-{{scenarioID}}",
    "params": {
        "from": "{{account1}}",
        "contractName": "SimpleToken"
    },
    "labels": {
    	"scenario.id": "{{scenarioID}}",
    	"id": "{{deployTxTwoID}}"
    }
}
      """
    Then the response code should be 202
    Then Envelopes should be in topic "tx.sender"
    Then Envelopes should be in topic "tx.decoded"
    When I send "GET" request to "{{global.tx-scheduler}}/transactions?idempotency_keys={{idempotencykey}}"
    Then the response code should be 200
    And Response should have the following fields
      | 0.idempotencyKey   | 0.schedule.jobs[0].status | 0.schedule.jobs[0].labels.id |
      | {{idempotencykey}} | MINED                     | {{deployTxOneID}}            |

  Scenario: Send twice different transaction using same idempotency key
    Given I register the following alias
      | alias          | value           |
      | idempotencykey | {{random.uuid}} |
      | deployTxOneID  | {{random.uuid}} |
      | deployTxTwoID  | {{random.uuid}} |
    Then I track the following envelopes
      | ID                |
      | {{deployTxOneID}} |
    Then  I set the headers
      | Key               | Value                    |
      | Authorization     | Bearer {{tenant1.token}} |
      | X-Idempotency-Key | {{idempotencykey}}       |
    When I send "POST" request to "{{global.tx-scheduler}}/transactions/deploy-contract" with json:
  """
{
    "chain": "besu_1-{{scenarioID}}",
    "params": {
        "from": "{{account1}}",
        "contractName": "SimpleToken"
    },
    "labels": {
    	"scenario.id": "{{scenarioID}}",
    	"id": "{{deployTxOneID}}"
    }
}
      """
    Then the response code should be 202
    Then  I set the headers
      | Key               | Value                    |
      | Authorization     | Bearer {{tenant1.token}} |
      | X-Idempotency-Key | {{idempotencykey}}       |
    When I send "POST" request to "{{global.tx-scheduler}}/transactions/deploy-contract" with json:
  """
{
    "chain": "besu_1-{{scenarioID}}",
    "params": {
        "oneTimeKey": true,
        "contractName": "SimpleToken"
    },
    "labels": {
    	"scenario.id": "{{scenarioID}}",
    	"id": "{{deployTxTwoID}}"
    }
}
      """
    Then the response code should be 409
    And Response should have the following fields
      | message                                                                                                          |
      | DB101@use-cases.send-tx: a transaction request with the same idempotency key and different params already exists |
    Then Envelopes should be in topic "tx.sender"
    Then Envelopes should be in topic "tx.decoded"
    When I send "GET" request to "{{global.tx-scheduler}}/transactions?idempotency_keys={{idempotencykey}}"
    Then the response code should be 200
    And Response should have the following fields
      | 0.idempotencyKey   | 0.schedule.jobs[0].status | 0.schedule.jobs[0].labels.id |
      | {{idempotencykey}} | MINED                     | {{deployTxOneID}}            |
  
# @TODO Restore once we can create ONLY MULTI-TENANCY scenarios
#  Scenario: Send twice same transaction using same idempotency key and different tenant
#    Given I register the following alias
#      | alias          | value           |
#      | idempotencykey | {{random.uuid}} |
#      | deployTxOneID  | {{random.uuid}} |
#      | deployTxTwoID  | {{random.uuid}} |
#    Then I track the following envelopes
#      | ID                |
#      | {{deployTxOneID}} |
#      | {{deployTxTwoID}} |
#    Then  I set the headers
#      | Key               | Value                    |
#      | Authorization     | Bearer {{tenant1.token}} |
#      | X-Tenant-ID       | {{tenant1.tenantID}}     |
#      | X-Idempotency-Key | {{idempotencykey}}       |
#    When I send "POST" request to "{{global.tx-scheduler}}/transactions/deploy-contract" with json:
#  """
#{
#    "chain": "besu_1-{{scenarioID}}",
#    "params": {
#        "from": "{{account1}}",
#        "contractName": "SimpleToken"
#    },
#    "labels": {
#    	"scenario.id": "{{scenarioID}}",
#    	"id": "{{deployTxOneID}}"
#    }
#}
#      """
#    Then the response code should be 202
#    Then  I set the headers
#      | Key               | Value                    |
#      | Authorization     | Bearer {{tenant1.token}} |
#      | X-Tenant-ID       | {{tenant1.tenantID}}     |
#      | X-Idempotency-Key | {{idempotencykey}}       |
#    When I send "POST" request to "{{global.tx-scheduler}}/transactions/deploy-contract" with json:
#  """
#{
#    "chain": "besu_2-{{scenarioID}}",
#    "params": {
#        "from": "{{account2}}",
#        "contractName": "SimpleToken"
#    },
#    "labels": {
#    	"scenario.id": "{{scenarioID}}",
#    	"id": "{{deployTxTwoID}}"
#    }
#}
#      """
#    Then the response code should be 202
#    Then Envelopes should be in topic "tx.decoded"
#    Then  I set the headers
#      | Key           | Value                    |
#      | Authorization | Bearer {{tenant1.token}} |
#      | X-Tenant-ID   | {{tenant1.tenantID}}     |
#    When I send "GET" request to "{{global.tx-scheduler}}/transactions?idempotency_keys={{idempotencykey}}"
#    Then the response code should be 200
#    And Response should have the following fields
#      | 0.idempotencyKey   | 0.schedule.jobs[0].status | 0.schedule.jobs[0].labels.id |
#      | {{idempotencykey}} | MINED                     | {{deployTxOneID}}            |
#    Then  I set the headers
#      | Key           | Value                    |
#      | Authorization | Bearer {{tenant1.token}} |
#      | X-Tenant-ID   | {{tenant1.tenantID}}     |
#    When I send "GET" request to "{{global.tx-scheduler}}/transactions?idempotency_keys={{idempotencykey}}"
#    Then the response code should be 200
#    And Response should have the following fields
#      | 0.idempotencyKey   | 0.schedule.jobs[0].status | 0.schedule.jobs[0].labels.id |
#      | {{idempotencykey}} | MINED                     | {{deployTxTwoID}}            |
