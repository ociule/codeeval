DEBUG = false 

def parse_test line
    string, subs, expected = line.strip.split ";"
    subss = subs.split(",")
    subs = []
    (0..subss.length - 1).step 2 do |i|
        subs.push([subss[i], subss[i + 1]])
    end
    return string, subs, expected
end

def transform string, subs
    transformed = [false] * string.length
    subs.each do |sub|
        starting_ix = 0
        f, r = sub
        while starting_ix < string.length do
            ix = string.index(f, starting_ix)
            if not ix.nil?
                starting_ix = ix + r.length
                already_transformed = transformed[ix, f.length]
                if not already_transformed.any?
                    #puts "%s found %s at %d" % [string, f, ix]
                    # replace f with r
                    string = string[0, ix] + r + string[ix + f.length, string.length]
                    # mark as transformed
                    transformed = transformed[0, ix] + [true] * r.length + transformed[ix + f.length, transformed.length]
                    if string.length != transformed.length
                        raise "PANIC"
                    end
                    #ts = transformed.map{|v| v ? "1" : "0" }.join
                    #puts "%s %s" % [string, ts] 
                end
            else
                break # signal to move to next sub
            end
        end
    end
    return string
end

def transform2 string, subs
    transformed = [false] * string.length
    subs.each do |sub|
        starting_ix = 0
        f, r = sub
        rs = r.split("").map{|c| (c.to_i + 2).to_s}.join
        while starting_ix < string.length do
            ix = string.index(f, starting_ix)
            if not ix.nil?
                starting_ix = ix + r.length
                puts "%s found %s at %d -> %s, will search again from %d" % [string, f, ix, rs, starting_ix] if DEBUG
                # replace f with r
                string = string[0, ix] + rs + string[ix + f.length, string.length]
                # mark as transformed
                puts "%s" % [string] if DEBUG
            else
                break # signal to move to next sub
            end
        end
    end
    return string.split("").map{|c| (c.to_i > 1) ? c.to_i - 2 : c }.join
end

File.open(ARGV[0]).each_line do |line|
    string, substitutions, expected = parse_test line
    #p string, substitutions.flatten.join(",")
    actual = transform2(string, substitutions)
    puts "%s %s" % [actual, (not expected.nil? and actual == expected) ? "OK" : ""]
end
