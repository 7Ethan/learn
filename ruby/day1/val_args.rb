def Sample(*test)
  puts "The number of args is #{test.length}"
  for i in 0...test.length
    puts "#{i} : #{test[i]}"
  end
end

Sample "23e",12