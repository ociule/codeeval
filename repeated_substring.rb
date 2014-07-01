
def longest_common_prefix(a, b)
    s = [a.length, b.length].min
    (0..s - 1).each do |i|
        if a[i] != b[i]
            s = i
            break
        end
    end
    a[0, s]
end

def gen_suffix_array line
    l = Array.new line.length - 1
    (0..line.length - 1).each do |i|
        l[i] = line[i, line.length]
    end
    l.sort!
end

def gen_preffix_array line
    l = Array.new line.length - 1
    (0..line.length - 1).each do |i|
        l[i] = line[0, line.length - i]
    end
    l.sort!
end

def does_overlap substr, str
    is = str.index(substr)
    return str[is + substr.length, str.length].index(substr) == nil
end

def longest a, b
    a.length > b.length ? a : b
end

def transform line
    suffixes = gen_suffix_array line
    
    candidates = []
    (0..line.length - 2).each do |i|
        a = suffixes[i]
        b = suffixes[i + 1]
        lcp = longest_common_prefix(a, b)
        if not lcp.empty? and not lcp.strip.empty?
            candidates.push [lcp, a, b]
            (gen_preffix_array(lcp) - [lcp]).each {|p| candidates.push [p, a, b]}
        end
    end
    
    # Filter candidates
    candidates.select! {|c| !does_overlap c[0], longest(c[1], c[2]) }
    max_len = candidates.map{|a| a[0].length}.reduce{|a, b| [a, b].max}
    candidates.select! {|c| c[0].length == max_len }
    candidates.sort_by! {|c| line.index c[0]}
    candidates[0].nil? ? "NONE" : candidates[0][0]
end

File.open(ARGV[0]).each_line do |line|
    puts transform(line.strip)
end
