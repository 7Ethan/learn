def tell_the_true(options={})
    if options[:profession] == :lawyer
        'it could beleieved that this is certainly not false'
    else
        true
    end
end

puts tell_the_true({:profession=>:lawyer})