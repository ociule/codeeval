
class PriorityQueue
    def initialize n
        @n = n
        @objects = {}
        @length = 0
    end

    def _add obj, priority
        if @objects[priority].nil?
            @objects[priority] = [obj]
        else
            @objects[priority] << obj
        end
    end

    def add obj, priority
        if @length < @n
            _add obj, priority
            @length += 1
        else
            #puts "SHOULD P " + obj
            prev_priors = @objects.keys.sort
            if priority > prev_priors[0]
                #puts "YES %d %s" % [priority, prev_priors.to_s]
                # make place
                @objects[prev_priors[0]].shift
                if @objects[prev_priors[0]].length == 0
                    @objects.delete prev_priors[0]
                end
                _add obj, priority
            end
        end
    end

    def get
        @objects.values.flatten
    end
end


lines = nil
n = false 

File.open(ARGV[0]).each_line do |line|
    line.strip!
    if not n
        n = line.to_i
        lines = PriorityQueue.new n
    else
        lines.add line, line.length
    end
end

lines = lines.get.sort_by{|w| w.length}.reverse
lines.each {|l| puts l}
