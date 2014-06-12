Feature: exercise pwd

  Scenario: Run in current directory
    When I run `pwd`
    Then the stdout should contain "/Users/gilhai01/code/gobase"
