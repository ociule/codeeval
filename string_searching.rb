
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
        l[i] = line[i, line.length - i]
    end
    l.sort!
end

SEARCH_ASTERISK = "^"

def has_asterisk substr
    apos = substr.index(SEARCH_ASTERISK)
    return apos != nil
end

def asterisk_search suffix, substr
    first_asterisk_pos = substr.index(SEARCH_ASTERISK)
    return false if substr[0, first_asterisk_pos - 1] != suffix[0, first_asterisk_pos - 1]
    next_suff = suffix[first_asterisk_pos - 1, suffix.length]
    next_substr = substr[first_asterisk_pos + 1, substr.length]
    return search next_suff, next_substr 
end

def transform line
    str, substr = line.split ","
    substr.sub! "\\*", "+"
    substr.sub! "*", SEARCH_ASTERISK
    # And baaack
    substr.sub! "+", "*"

    return search str, substr
end

def search str, substr
    sa = gen_suffix_array str
    ast = has_asterisk substr
    sa.each do |s|
        if s.length >= substr.length
            if has_asterisk substr
                return asterisk_search s, substr
            else
                return true if longest_common_prefix(s, substr) == substr
            end
        end
    end
    return false
end

File.open(ARGV[0]).each_line do |line|
    puts transform(line.strip)
end
