Feature: Exercise the cat tool

  Scenario Outline: cat can read a file
    Given the input <file>
    When I cat the file
    Then I should see the content of the file
    And there should be no error prints
    Examples:
      | file |
      | resources/001.txt |

  Scenario: read an empty file
    Given an empty file
    When I cat the file
    Then I should see nothing
    And there should be no error prints

