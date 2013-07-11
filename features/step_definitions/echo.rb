require 'open3'

When /^I run echo "(.+)"$/ do | text |
  @input = text
  @stdin, @stdout, @stderr = Open3.popen3("bin/echo #{text}")
end

When /^I run echo -n "(.+)"$/ do | text |
  @input = text
  @stdin, @stdout, @stderr = Open3.popen3("bin/echo -n #{text}")
end

Then /^I should see the input with a newline$/ do 
  @stdout.read.should == "#{@input}\n"
end

Then /^I should see the input without a newline$/ do 
  @stdout.read.should == @input
end