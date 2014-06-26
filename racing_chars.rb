class Track
    Checkpoint = "C"
    Gate = "_"
    Straight = "|"
    Left = "/"
    Right = "\\"

    def initialize
        @track = []
    end
    def push(line)
        @track.push(line)
    end

    def checkpoint(lineNumber = 0)
        @track[lineNumber].index Checkpoint
    end

    def gate(lineNumber = 0)
        @track[lineNumber].index Gate 
    end

    def solve
        previous = checkpoint || gate
        @track.each_with_index do |line, i|
            current = checkpoint(i) || gate(i)
            if previous == current
                line[current] = Straight
            elsif previous < current
                line[current] = Right
            else
                line[current] = Left
            end
            previous = current
            puts line
        end
    end
    
    def to_s
        @track.to_s
    end
end

track = Track.new
File.open(ARGV[0]).each_line do |line|
    track.push(line.strip)
end
track.solve
