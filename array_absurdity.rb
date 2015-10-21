
def solve line 
    ss = line.split ";"
    n = ss[0].to_i
    numbers = ss[1].strip.split(",").collect{|x| x.to_i}

    sum = numbers.reduce{|x, y| x + y}

    # sum for the integers from 0 to n should be (n - 2) * (n - 1) / 2
    correct_sum = (n - 2) * (n - 1) / 2
    return sum - correct_sum
end

File.open(ARGV[0]).each_line do |line|
    puts solve line.strip
end
