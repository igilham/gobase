Feature: exercise basename application

  Scenario Outline: basename prints the basename of its argument
    When I run `basename <argument>`
    Then the stdout should contain exactly "<basename>\n"
    Examples:
      | argument | basename |
      | .        | .        |
      | ..       | ..       |
      | ./a      | a        |
      | ./a/b    | b        |
      | //       | /        |
      | a/       | a        |
      | a//      | a        |
      | a//b     | b        |
      | a///b    | b        |
