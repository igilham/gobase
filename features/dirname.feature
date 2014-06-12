Feature: exercise dirname application

  Scenario: run dirname with no arguments
    When I run `dirname`
    Then the stdout should not contain anything
    And the stderr should contain "not enough arguments"

  Scenario Outline: dirname prints the dirname of its argument
    When I run `dirname <argument>`
    Then the stdout should contain exactly "<dirname>\n"
    Examples:
      | argument | dirname  |
      | .        | .        |
      | ..       | .        |
      | a        | .        |
      | ./a      | .        |
      | ./a/b    | ./a      |
      | /        | /        |
      | //       | /        |
      | a/       | .        |
      | a/b      | a        |
      | a//      | .        |
      | a//b     | a        |
      | a///b    | a        |
