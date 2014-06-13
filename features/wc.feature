Feature: exercise wc

  Scenario: count newlines in empty file
    Given an empty file named "foo"
    When I run `wc -l foo`
    Then the stdout should contain "0 foo"

  Scenario: count words in empty file
    Given an empty file named "foo"
    When I run `wc -w foo`
    Then the stdout should contain "0 foo"

  Scenario: count bytes in empty file
    Given an empty file named "foo"
    When I run `wc -c foo`
    Then the stdout should contain "0 foo"

  Scenario: count haracters in empty file
    Given an empty file named "foo"
    When I run `wc -m foo`
    Then the stdout should contain "0 foo"

  Scenario: count everything in empty file
    Given an empty file named "foo"
    When I run `wc foo`
    Then the stdout should contain "0     0     0 foo"

  Scenario: count newlines in single line file
    Given a file named "foo" with:
      """
      hello world
      """
    When I run `wc -l foo`
    Then the stdout should contain "0 foo"

  Scenario: count words in single line file
    Given a file named "foo" with:
      """
      hello world
      """
    When I run `wc -w foo`
    Then the stdout should contain "1 foo"

  Scenario: count bytes in single line file
    Given a file named "foo" with:
      """
      hello world
      """
    When I run `wc -c foo`
    Then the stdout should contain "11 foo"

  Scenario: count characters in single line file
    Given a file named "foo" with:
      """
      hello world
      """
    When I run `wc -m foo`
    Then the stdout should contain "11 foo"

  Scenario: count everything in single line file
    Given a file named "foo" with:
      """
      hello world
      """
    When I run `wc foo`
    Then the stdout should contain "0     1    11 foo"

  Scenario: count newlines in multi-line file
    Given a file named "foo" with:
      """
      foo bar
      baz qux
      """
    When I run `wc -l foo`
    Then the stdout should contain "1 foo"

  Scenario: count words in multi-line file
    Given a file named "foo" with:
      """
      foo bar
      baz qux
      """
    When I run `wc -w foo`
    Then the stdout should contain "3 foo"

  Scenario: count bytes in multi-line file
    Given a file named "foo" with:
      """
      foo bar
      baz qux
      """
    When I run `wc -c foo`
    Then the stdout should contain "15 foo"

  Scenario: count characters in multi-line file
    Given a file named "foo" with:
      """
      foo bar
      baz qux
      """
    When I run `wc -m foo`
    Then the stdout should contain "15 foo"

  Scenario: count everything in multi-line file
    Given a file named "foo" with:
      """
      foo bar
      baz qux
      """
    When I run `wc foo`
    Then the stdout should contain "1     3    15 foo"
