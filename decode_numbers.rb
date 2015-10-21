
def count(start, line)
    cutline = line[start..line.length]
    if cutline.length == 1
        return 1
    elsif cutline.length == 2
        if cutline.to_i <= 26
            return 2
        else
            return 1
        end
    else
        if cutline[0..1].to_i <= 26
            return count(start+1, line) + count(start+2, line)
        else
            return count(start+1, line)
        end
    end
end

def solve line
    count(0, line)
end

File.open(ARGV[0]).each_line do |line|
    puts solve line.strip 
end
