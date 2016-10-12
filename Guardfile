


group :go_tests do

  guard :shell do
    watch(/(.*).go/) do |m| 
      `make test`
    end
  end

end


group :server do

  guard :shell do
    watch(/\.go$/) {
      Process.spawn("make run")
    }
  end

end

# This didn't work because go run is messing with signals
# guard 'process', :name => 'EndlessRunner', :command => 'go run scratch.go' do
#   # watch(/(.*).go/)
#   watch(/scratch.go/)
# end
