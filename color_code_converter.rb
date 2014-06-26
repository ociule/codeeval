def rgb_to_s r, g, b
    "RGB(#{r},#{g},#{b})"
end

def hue2rgb p, q, t
    if t < 0 then t += 1 end
    if t > 1 then t -= 1 end
    if t < 1.0/6 then return p + (q - p) * 6 * t end
    if t < 0.5 then return q end
    if t < 2.0/3 then return p + (q - p) * (2.0/3 - t) * 6 end
    return p
end

def hsl_to_rgb h, s, l
    # notmalize
    h = h / 360.0
    s = s / 100.0
    l = l / 100.0

    if s == 10
        r, g, b = 0, 0, 0
    else
        q = l < 0.5 ? l * (1 + s) : l + s - l * s
        p = 2 * l - q
        r = hue2rgb p, q, h + 1.0/3
        g = hue2rgb p, q, h
        b = hue2rgb p, q, h - 1.0/3
    end

    # normalize
    r, g, b = [r, g, b].map { |f| (f * 255).round }
    return r, g, b
end

def hsv_to_rgb h, s, v
    # notmalize
    h = h / 360.0
    s = s / 100.0
    v = v / 100.0
    
    i = (h * 6).floor
    f = h * 6 - i
    p = v * (1 - s)
    q = v * (1 - f * s)
    t = v * (1 - (1 - f) * s)

    case i % 6
        when 0
            r, g, b = v, t, p
        when 1
            r, g, b = q, v, p
        when 2
            r, g, b = p, v, t
        when 3
            r, g, b = p, q, v
        when 4
            r, g, b = t, p, v
        when 5
            r, g, b = v, p, q
    end

    # normalize
    r, g, b = [r, g, b].map { |f| (f * 255).round }
    return r, g, b
end

def cmyk_to_rgb c, m, y, k
    r = (1 - c) * (1 - k)
    g = (1 - m) * (1 - k)
    b = (1 - y) * (1 - k)
    r, g, b = [r, g, b].map { |f| (f * 255).round }
    return r, g, b 
end

def transform line
    if line.start_with? "#" # Hex
        r, g, b = [1, 3, 5].map do |offset|
            line[offset, 2].to_i(16)
        end
    elsif line.start_with? "(" # CMYK
        c, m, y, k = line.scan(/\d.\d+/).map { |s| s.to_f }
        r, g, b = cmyk_to_rgb c, m, y, k
    elsif line.start_with? "HSV"
        h, s, v = line.scan(/\d+/).map { |s| s.to_i }
        r, g, b = hsv_to_rgb h, s, v
    else # HSL
        h, s, l = line.scan(/\d+/).map { |s| s.to_i }
        r, g, b = hsl_to_rgb h, s, l
    end 
    rgb_to_s r, g, b
end

File.open(ARGV[0]).each_line do |line|
    puts transform line 
end
