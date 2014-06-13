Feature: Exercise the cat tool

  Scenario: read an empty file
    Given an empty file named "foo"
    When I run `cat foo`
    Then the stdout should not contain anything
    And the stderr should not contain anything

  Scenario: read a simple file
    Given a file named "test.txt" with:
      """
      hello world
      """
    When I run `cat test.txt`
    Then the stdout should contain exactly "hello world"
    And the stderr should not contain anything

  Scenario: concatenate two files
    Given a file named "foo" with:
      """
      foo
      """
    And a file named "bar" with:
      """
      bar
      """
    When I run `cat foo bar`
    Then the stdout should contain exactly "foobar"
    And the stderr should not contain anything
