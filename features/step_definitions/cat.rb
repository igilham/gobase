require 'open3'

def call_cat_file(file)
  Open3.popen3 "bin/cat #{file}"
end

Given /^the input ([^ ]+)$/ do | file |
  @file = file
end

Given /^an empty file$/ do
  @file = 'resources/empty.txt'
end

When /^I cat the file$/ do
  @stdin, @stdout, @stderr, @wait_thr = call_cat_file @file
end

Then /^I should see the content of the file$/ do
  expect(@stdout.read).to eq(IO.read(@file))
end

Then /^there should be no error prints$/ do
  expect(@stderr.read).to eq("")
  expect(@wait_thr.value).to eq(0)
end

Then /^I should see nothing$/ do
  expect(@stderr.read).to eq("")
end
