
def step n
    sum = 0
    while n > 0
        highest_pow_10 = 10 ** (Math.log10 n).to_i
        highest_digit = n / highest_pow_10 
        #puts "hd = " + highest_digit.to_s + " " + highest_pow_10.to_s
        sum += highest_digit ** 2
        n = n % highest_pow_10
    end
    sum += n ** 2
    return sum 
end

def transform line
    m = line.to_i
    return 1 if m == 1
    happy = false
    already_seen = {}
    while already_seen[m].nil?
        already_seen[m] = true
        m = step m
        return 1 if m == 1
    end
    return 0
end

File.open(ARGV[0]).each_line do |line|
    puts transform line.strip
end
