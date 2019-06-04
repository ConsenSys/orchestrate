Feature: deploy an ERC20 contract
  As an external developer
  I want to deploy a contract

  Scenario: Create an instance of ERC20
    Given I have the following envelope:
      | chainId | from                                       | contractName | methodSignature        | gas     |
      | 888     | 0x7E654d251Da770A068413677967F6d3Ea2FeA9E4 | SimpleToken        | constructor() | 2000000 |
    When I send these envelopes to CoreStack
    Then CoreStack should receive them
    Then the tx-crafter should set the data
    Then the tx-nonce should set the nonce
    Then the tx-signer should sign
    Then the tx-sender should send the tx
    Then the tx-listener should catch the tx
    Then the tx-decoder should decode
