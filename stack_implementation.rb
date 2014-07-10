def transform ints
    stack = []
    out = []
    ints.each {|i| stack.push i}
    (0..stack.length - 1).each do |i|
        ii = stack.pop
        if i % 2 == 0
            out.push ii
        end
    end
    return out.join " "
end

File.open(ARGV[0]).each_line do |line|
    split = line.strip.split
    split = split.map {|c| c.to_i}
    puts transform split 
end
