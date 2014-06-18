

def count_paths(step, cache)
    if step <= 0
        return 0
    elsif step == 1
        return 1
    elsif step == 2
        return 2
    else
        if cache[step] != nil
            return cache[step]
        else
            res = count_paths(step-1, cache) + count_paths(step-2, cache)
            cache[step] = res
            return res 
        end
    end
end

def count_paths2(n)
    n = n + 1
    phi = 1.61803398874989484820458683436563811772030917980576
    top = (phi ** n) - ((-phi) ** -n)
    raw = top / (5 ** 0.5)
    return (raw + 0.5).floor()
end

cache = {}
File.open(ARGV[0]).each_line do |line|
    line.strip!

    puts count_paths(line.to_i, cache)
end
