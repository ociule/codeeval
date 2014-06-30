def transform line
    seq = line.strip.split(',').map {|i| i.to_i}
    freqs = Hash.new(0)
    max_elem = nil
    max_freq = 0

    seq.each do |i|
        freqs[i] += 1
        if freqs[i] > max_freq
            max_freq = freqs[i]
            max_elem = i
        end
    end

    if max_freq > seq.length / 2
        return max_elem
    else
        return 'None'
    end
end

File.open(ARGV[0]).each_line do |line|
    puts transform line
end
