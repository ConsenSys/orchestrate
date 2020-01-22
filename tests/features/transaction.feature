@public-tx
Feature: Multiple transactions
  As an external developer
  I want to process multiple transactions

  Scenario: Send transactions
    Given I register the following contract
      | name         | artifacts        | tenantid                             |
      | SimpleToken  | SimpleToken.json | f30c452b-e5fb-4102-a45d-bc00a060bcc6 |
    And I have deployed contract "token"
      | chain.nodeName | from                                       | contract.name | method.sig    | tx.gas  | tenantid                             |
      | geth           | 0x7E654d251Da770A068413677967F6d3Ea2FeA9E4 | SimpleToken   | constructor() | 2000000 | f30c452b-e5fb-4102-a45d-bc00a060bcc6 |
    When I send envelopes to topic "tx.crafter"
      | chain.nodeName | from                                       | tx.to | method.sig                | args                                           | tenantid                             |
      | geth           | 0x7E654d251Da770A068413677967F6d3Ea2FeA9E4 | token | transfer(address,uint256) | 0xdbb881a51cd4023e4400cef3ef73046743f08da3,0x1 | f30c452b-e5fb-4102-a45d-bc00a060bcc6 |
      | geth           | 0x7E654d251Da770A068413677967F6d3Ea2FeA9E4 | token | transfer(address,uint256) | 0x6009608a02a7a15fd6689d6dad560c44e9ab61ff,0x2 | f30c452b-e5fb-4102-a45d-bc00a060bcc6 |
    Then Envelopes should be in topic "tx.crafter"
    Then Envelopes should be in topic "tx.nonce"
    Then Envelopes should be in topic "tx.signer"
    And Envelopes should have nonce set
    Then Envelopes should be in topic "tx.sender"
    And Envelopes should have raw and hash set
    Then Envelopes should be in topic "tx.decoder"
    Then Envelopes should be in topic "tx.decoded"
    And Envelopes should have log decoded
