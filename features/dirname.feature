Feature: exercise dirname application

  Scenario Outline: dirname prints the dirname of its argument
    When I run dirname with an <argument>
    Then dirname should print the <dirname> of the argument
    Examples:
      | argument | dirname  |
      |          | .        |
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
