Feature: exercise the echo application

  Scenario Outline: prints text
    When I run echo "<input>"
    Then I should see the input with a newline
    Examples:
    | input             |
    | hello world       |
    | que pasa, hombre? |


  Scenario Outline: suppress newline
    When I run echo -n "<input>"
    Then I should see the input without a newline
    Examples:
    | input             |
    | hello world       |
    | que pasa, hombre? |
