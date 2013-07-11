require 'open3'

When /^I run echo with a single quoted argument$/ do
  @input = 'hello, world!'
  @stdin, @stdout, @stderr = Open3.popen3("bin/echo \"#{@input}\"")
end

When /^I run echo with three arguments$/ do
  @input = 'goodbye cruel world'
  @stdin, @stdout, @stderr = Open3.popen3("bin/echo #{@input}")
end

When /^I run echo -n with a single quoted argument$/ do
  @input = 'hello, world!'
  @stdin, @stdout, @stderr = Open3.popen3("bin/echo -n #{@input}")
end

When /^I run echo -n with three arguments$/ do
  @input = 'goodbye cruel world'
  @stdin, @stdout, @stderr = Open3.popen3("bin/echo -n #{@input}")
end

When /^I run echo without arguments$/ do
  @stdin, @stdout, @stderr = Open3.popen3("bin/echo")
end

Then /^it should print the input with a newline$/ do 
  @stdout.read.should == "#{@input}\n"
end

Then /^it should print the input without a newline$/ do 
  @stdout.read.should == @input
end

Then /^it should print a newline$/ do 
  @stdout.read.should == "\n"
end
