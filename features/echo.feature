Feature: exercise the echo application

    Scenario: echo prints text
      When I run echo with a single quoted argument
      Then it should print the input with a newline

    Scenario: echo -n suppresses newline
      When I run echo -n with a single quoted argument
      Then it should print the input without a newline
    
    Scenario: echo arg1 arg2 arg3 prints multiple arguments
      When I run echo with three arguments
      Then it should print the input with a newline
    
    Scenario: echo -n arg1 arg2 arg3 prints multiple arguments and suppresses newline
      When I run echo -n with three arguments
      Then it should print the input without a newline

    Scenario: echo without arguments prints a newline
      When I run echo without arguments
      Then it should print a newline
