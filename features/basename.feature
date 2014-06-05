Feature: exercise basename application

  Scenario Outline: basename prints the basename of its argument
    When I run basename with an <argument>
    Then basename should print the <basename> of the argument
    Examples:
      | argument | basename |
      |          | .        |
      | .        | .        |
      | ..       | ..       |
      | ./a      | a        |
      | ./a/b    | b        |
      | //       | /        |
      | a/       | a        |
      | a//      | a        |
      | a//b     | b        |
      | a///b    | b        |
