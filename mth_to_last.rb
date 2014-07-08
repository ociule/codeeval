def transform line
    lineSplit = line.split
    m = lineSplit[-1].to_i
    lineSplit[-m-1]
end

File.open(ARGV[0]).each_line do |line|
    puts transform line.strip
end
