VALUES = {
    "million" => 10**6,
    "thousand" => 10**3,
    "hundred" => 100,
    "ninety" => 90,
    "eighty" => 80,
    "seventy" => 70,
    "sixty" => 60,
    "fifty" => 50,
    "forty" => 40,
    "thirty" => 30,
    "twenty" => 20,

    'one'  => 1,
    'two' => 2,
    'three' => 3,
    'four' => 4,
    'five' => 5,
    'six' => 6,
    'seven' => 7,
    'eight' => 8,
    'nine' => 9,
    'ten' => 10,
    'eleven' => 11,
    'twelve' => 12,
    'thirteen' => 13,
    'fourteen' => 14,
    'fifteen' => 15,
    'sixteen' => 16,
    'seventeen' => 17,
    'eighteen' => 18,
    'nineteen' => 19,
}

PARSE_ORDER = [10 ** 6, 1000, 100, 99]

def get_next_smaller number 
    iw = PARSE_ORDER.index number 
    if not iw.nil? and iw >= 0
        PARSE_ORDER[iw + 1]
    else
        nil
    end
end

def generate base, number
    #puts "generate %d %d" % [base, number]
    if base == 1000000 and number == 0 # special case for 0
        return ["zero"]
    end
    out = []
    if number / base > 0 and base != 99 
        prefixNum = number / base 
        prefix = generate get_next_smaller(base), prefixNum
        out += prefix if prefix != ['zero']
        out += [VALUES.key(base)]
    end

    if base == 99 
        if number > 20
            tens = number / 10
            rest = number % 10
            out += [VALUES.key(tens * 10)] if tens > 0
            out += [VALUES.key(rest)] if rest > 0
        else
            if number > 0
                out += [VALUES.key(number)]
            end
        end
    else
        suffixNum = number % base 
        suffix = generate get_next_smaller(base), suffixNum
        out += suffix
    end
    return out
end

def transform number
    generate 10 ** 6, number
end

File.open(ARGV[0]).each_line do |line|
    tokens = transform(line.strip.to_i) + ["dollars"]
    tokens = tokens.map {|t| t.capitalize }
    puts tokens.join("")
end
