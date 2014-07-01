
# Tortoise and hare algorithm
# We do not have access to the iterated function that geenrates the sequence, but directly to the sequence
# This means we'll use work directly on the sequence
def transform line
    i = 1
    tortoise = line[i]
    hare = line[i * 2]
    while tortoise != hare
        i += 1
        tortoise = line[i]
        hare = line[i * 2]
    end
    v = i

    mu = 0
    tortoise = line[mu]
    hare = line[v * 2 + mu]
    while tortoise != hare
        mu += 1
        tortoise = line[mu]
        hare = line[v * 2 + mu]
    end

    lam = 1
    hare = line[mu + lam]
    while tortoise != hare
        lam += 1
        hare = line[mu + lam]
    end
    #puts "v mu lam %d %d %d" % [v, mu, lam]
   
    line[mu, lam] 
end


def transform2 line
    f = Hash.new
    
    next_ = nil
    i = 0
    while f[next_].nil?
        this = line[i]
        next_ = line[i + 1]
        f[this] = next_
        i += 1
    end

    mu = line.index(next_)
    mulam = line[mu+1,line.length].index(next_) + mu+1

    lam = mulam - mu

    line[mu, lam]
end

File.open(ARGV[0]).each_line do |line|
    lineSplit = line.strip.split.map {|i| i.to_i }
    puts transform2(lineSplit).join " "
end
