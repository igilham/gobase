require 'open3'

When(/^I run dirname with an ([^ ]*)$/) do | argument |
  @argument = argument
  @stdin, @stdout, @stderr = Open3.popen3("bin/dirname \"#{@argument}\"")
end

Then(/^dirname should print the ([^ ]+) of the argument$/) do | dirname |
  expect(@stdout.read).to eq("#{dirname}\n")
end
