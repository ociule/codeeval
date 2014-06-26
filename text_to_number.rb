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

    'zero' => 0,
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

PARSE_ORDER = ["under_hundred", "hundred", "thousand", "million"]

def get_next_smaller what
    iw = PARSE_ORDER.index what
    if not iw.nil? and iw > 0
        PARSE_ORDER[iw - 1]
    else
        nil
    end
end

def parse what, text
    #puts "parsing %s #{text}" % what
    iw = text.index(what)
    outNum = 0
    if not iw.nil?
        prefix = text[0, iw]
        prefixNum = parse get_next_smaller(what), prefix
        outNum = VALUES[what] * prefixNum
    else
        iw = -1
    end
    if what == "under_hundred"
        text.each {|token| outNum += VALUES[token]}
    else
        suffix = text[iw + 1, text.length]
        suffixNum = parse get_next_smaller(what), suffix
        outNum += suffixNum
    end
    return outNum
end

def transform line
    out = 0
    lineSplit = line.split
    negative = lineSplit[0] == "negative"
    lineSplit = lineSplit[1, lineSplit.length] if negative else lineSplit
    
    out = parse "million", lineSplit

    out * (negative ? -1 : 1)
end

File.open(ARGV[0]).each_line do |line|
    puts transform line
end
