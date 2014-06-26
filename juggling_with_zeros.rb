def transform line
    out = ""
    lineSplit = line.split
    (0..lineSplit.length - 1).step(2) do |i|
        group = lineSplit[i]
        if group.length == 1
            out += lineSplit[i+1]
        else
            out += "1" * lineSplit[i+1].length
        end
    end
    out.to_i(2)
end

File.open(ARGV[0]).each_line do |line|
    puts transform line
end
