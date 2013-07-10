require 'open3'

def call_cat(file)
  Open3.popen3 "bin/cat #{file}"
end

Given /^the input ([^ ]+)$/ do | file |
  @file = file
end

Given /^an empty file$/ do
  @file = 'resources/empty.txt'
end

When /^I cat the file$/ do
  @stdin, @stdout, @stderr = call_cat @file
end

Then /^I should see the content of the file$/ do
  @stdout.read.should == IO.read(@file)
end

Then /^there should be no error prints$/ do
  @stderr.read.should == ""
end

Then /^I should see nothing$/ do
  @stderr.read.should == ""
end

