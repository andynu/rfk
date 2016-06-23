# A sample Guardfile
# More info at https://github.com/guard/guard#readme
require 'rainbow'
## Uncomment and set this to only include directories you want to watch
# directories %w(app lib config test spec features) \
#  .select{|d| Dir.exists?(d) ? d : UI.warning("Directory #{d} does not exist")}

## Note: if you are using the `directories` clause above and you are not
## watching the project directory ('.'), then you will want to move
## the Guardfile to a watched dir and symlink it back, e.g.
#
#  $ mkdir config
#  $ mv Guardfile config/
#  $ ln -s config/Guardfile .
#
# and, you'll have to watch "config/Guardfile" instead of "Guardfile"

# Add files and commands to this file, like the example:
#   watch(%r{file/path}) { `command(s)` }
#
#guard :shell do
#  watch(/(.*).go/) {|m| 
#    test_file = m[1]+"_test.go"
#    dir = File.dirname(test_file)
#    if File.exist? test_file
#      cmd = "go test -cover ./#{dir}"
#      puts cmd
#      puts `#{cmd}`
#    end
#  }
#end

guard :shell do
  watch(/(.*).go/) do |m|
    puts "---"
    #cmd = "go build -o rfk-server ./server "
    cmd = "go run ./song_oracle/*.go"
    puts `#{cmd}`

    exit_code = $?.to_i
    if exit_code == 0
      puts Rainbow("OK").green
    else
      puts Rainbow("Error #{exit_code}").red
    end
  end
end

# vim: ft=ruby
