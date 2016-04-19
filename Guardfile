# A sample Guardfile
# More info at https://github.com/guard/guard#readme

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
  #watch(/(.*).go/) do |m|
  #  puts "---"
  #  puts `go build -o rfk-graph ./graph `
  #end
  watch(/.*\/.*.rs/) do |m|
    puts
    puts "-"*80
    puts
    cmd = 'cd src/server;RUST_BACKTRACE=1 cargo run --verbose -- --limit 3'
    puts "#{cmd}"
    puts `#{cmd}`
  end
end

# vim: ft=ruby
