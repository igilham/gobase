Feature: exercise the echo application

    Scenario: echo prints text
      When I run `echo "hello world"`
      Then the stdout should contain exactly "hello world\n"

    Scenario: echo -n suppresses newline
      When I run `echo -n "hello world"`
      Then the stdout should contain exactly "hello world"
    
    Scenario: echo arg1 arg2 arg3 prints multiple arguments
      When I run `echo arg1 arg2 arg3`
      Then the stdout should contain exactly "arg1 arg2 arg3\n"
    
    Scenario: echo -n arg1 arg2 arg3 prints multiple arguments and suppresses newline
      When I run `echo -n arg1 arg2 arg3`
      Then the stdout should contain exactly "arg1 arg2 arg3"

    Scenario: echo without arguments prints a newline
      When I run `echo`
      Then the stdout should contain exactly "\n"
