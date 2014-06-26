class Track
    attr_reader :lines, :curves
    def initialize description
        @lines = []
        @curves = []
        sdesc = description.split
        (0..sdesc.length - 1).step(2) do |step|
            @lines.push sdesc[step].to_f
            @curves.push sdesc[step + 1].to_i
        end
    end

    def race cars
        laps = {}
        cars.each do |car|
            laps[car.number] = car.race self
        end
        laps
    end
end

class Car
    attr_reader :number, :top_speed, :accel, :deccel
    def initialize description
        sdesc = description.split
        @number = sdesc[0].to_i
        @top_speed = sdesc[1].to_i / 3600.0 # convert top speed to miles per second
        @accel = sdesc[2].to_f
        @deccel = sdesc[3].to_f
    end

    def race track
        time = 0
        prev_speed = 0
        track.lines.each_with_index do |line, i|
            next_curve = track.curves[i]
            out_speed, line_time = race_line prev_speed, line, next_curve
            prev_speed = out_speed
            time += line_time
        end
        time
    end

    def race_line prev_speed, line_length, next_curve
        accel_time = (top_speed - prev_speed) / top_speed * accel
        accel_distance = (prev_speed + top_speed) / 2.0 * accel_time

        out_speed = (180 - next_curve) / 180.0 * top_speed
        deccel_time = (top_speed - out_speed) / top_speed * deccel
        deccel_distance = (out_speed + top_speed) / 2.0 * deccel_time

        top_speed_distance = line_length - accel_distance - deccel_distance
        top_speed_time = top_speed_distance / top_speed

        time = accel_time + top_speed_time + deccel_time
        #puts "racing #{number} line #{line_length} in #{time}s next curve #{next_curve}"

        return out_speed, time
    end
end

track = nil
cars = []
File.open(ARGV[0]).each_line do |line|
    if track.nil?
        track = Track.new(line)
    else
        cars.push Car.new(line)
    end
end

laps = track.race cars 
times = laps.values.sort
times.each do |time|
    puts "#{laps.key(time)} %.2f" % time
end
