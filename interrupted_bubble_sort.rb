
def inter_bb numbers, iter
    iter = [iter, numbers.length].min
    nIters = 0
    if numbers.length < 2
        return numbers
    end
    while nIters < iter do
        currentPos = 0
        while currentPos < numbers.length - 1
            #puts "Comparing #{numbers[currentPos]} with #{numbers[currentPos+1]}"
            if numbers[currentPos] > numbers[currentPos+1]
                #puts "Swapping"
                numbers[currentPos], numbers[currentPos + 1] = numbers[currentPos + 1], numbers[currentPos]
            end
            currentPos += 1
        end
        #puts numbers.join(" ")
        nIters += 1
    end
    numbers
end

def solve line 
    ss = line.split "|"
    numbers = ss[0].strip.split.collect{|x| x.to_i}
    iter = ss[1].to_i
    inter_bb(numbers, iter)
end

File.open(ARGV[0]).each_line do |line|
    puts (solve line.strip).join " "
end
