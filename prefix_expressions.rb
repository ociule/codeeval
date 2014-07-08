OPERATORS = ["+", "*", "/"]

def is_op s
    return OPERATORS.index(s) != nil
end

def transform line
    lineSplit = line.split
    intStack = []

    while lineSplit.length > 0
        p = lineSplit.pop
        if is_op p
            #puts "=== " + p + " " + intStack.to_s
            op1 = intStack.pop
            op2 = intStack.pop
            if p == "/"
                r = op1.to_f / op2
            else
                r = op1.send p, op2
            end
            intStack.push r
        else
            intStack.push p.to_i
        end
    end
    return intStack[0].round.to_i
end

File.open(ARGV[0]).each_line do |line|
    #puts line
    puts transform line.strip
end
