require 'set'


start = [0, 0]

$visited = Set.new [start]
LIMIT = 19
$queue = [start]

def digits_sum n
    sum = 0
  while n > 0 do
    d = n / 10
    r = n % 10
    sum += r
    n -= r
    if d > 0
      n /= 10
    end
  end
    return sum
    sum = 0
    while n > 0
        highest_pow_10 = 10 ** (Math.log10 n).to_i
        highest_digit = n / highest_pow_10
        sum += highest_digit
        n = n % highest_pow_10
    end
    sum += n
    return sum
end

def accessible? p
    return digits_sum(p[0].abs) + digits_sum(p[1].abs) <= LIMIT
end

def test_accessible?
    puts "Testing"
    raise "ERROR" if !accessible? [-5, -7]
    raise "ERROR" if !accessible? [5, 7]
    raise "ERROR" if accessible? [59, 79]
    raise "ERROR" if accessible? [-59, -79]
end

def explore p
    x, y = p[0], p[1]
    if not $visited.include? p and accessible? p
            $queue << p
            $visited.add p
    end
end

if ARGV[0] == "--test"
    test_accessible?
else
    while $queue.length > 0
        x, y = $queue.shift
        explore [x + 1, y]
        explore [x - 1, y]
        explore [x, y + 1]
        explore [x, y - 1]
    end
    puts $visited.size
end

