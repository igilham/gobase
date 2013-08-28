Feature: exercise basename application

  Scenario: basename prints the basename of its argument
    When I run basename with a file argument
    Then it should print the basename of the argument

