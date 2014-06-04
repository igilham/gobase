require 'open3'

When(/^I run basename with a file argument$/) do
  @input = 'src/basename/basename.go'
  @stdin, @stdout, @stderr = Open3.popen3("bin/basename \"#{@input}\"")
end

Then(/^it should print the basename of the argument$/) do
  expect(@stdout.read).to eq("basename.go\n")
end
