require 'open3'

When(/^I run basename with an ([^ ]*)$/) do | argument |
  @argument = argument
  @stdin, @stdout, @stderr = Open3.popen3("bin/basename \"#{@argument}\"")
end

Then(/^basename should print the ([^ ]+) of the argument$/) do | basename |
  expect(@stdout.read).to eq("#{basename}\n")
end
