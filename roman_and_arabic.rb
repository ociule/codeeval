ROMAN = {
    'I' => 1,
    'V' => 5,
    'X' => 10,
    'L' => 50,
    'C' => 100,
    'D' => 500,
    'M' => 1000,
}

def transform line 
    seq = line.strip

    sum = 0
    (0..seq.length-1).step(2) do |i|
        pair = seq[i,2]
        next_pair = seq[i+2,2]
        a, r = pair[0].to_i, ROMAN[pair[1]]
        m = 1
        nr = next_pair[1]
        if not nr.nil? and ROMAN[nr] > r
                m = -1
        end
        sum += (m * pair[0].to_i * ROMAN[pair[1]])
    end
    sum
end

File.open(ARGV[0]).each_line do |line|
    puts transform line
end
