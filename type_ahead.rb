STRING = "Mary had a little lamb its fleece was white as snow And everywhere that Mary went the lamb was sure to go It followed her to school one day which was against the rule It made the children laugh and play to see a lamb at school And so the teacher turned it out but still it lingered near And waited patiently about till Mary did appear Why does the lamb love Mary so the eager children cry Why Mary loves the lamb you know the teacher did reply"


DEBUG = false 

def to_ngrams str, ngram_length
    ngrams = {}
    sstr = str.split
    (0..sstr.length - ngram_length).each do |i|
        words, last_word = sstr[i, ngram_length - 1], sstr[i+ngram_length - 1]
        words = words.join "-"
        if ngrams[words].nil?
            ngrams[words] = {last_word => 1}
        else
            if ngrams[words][last_word].nil?
                ngrams[words][last_word] = 1
            else
                ngrams[words][last_word] += 1
            end
        end
    end

    if ngram_length > 2 and DEBUG
        ngrams.each do |k, v|
            puts k, v if v.length > 1
        end
    end
    return ngrams
end

NGRAMS = {}

NGRAMS[2] = to_ngrams STRING, 2

def transform words, nl 
    if NGRAMS[nl].nil?
        NGRAMS[nl] = to_ngrams STRING, nl
    end

    stats = NGRAMS[nl][words.split.join("-")]
    sum = stats.values.reduce{|a, b| a + b}
    by_score = {}
    stats.each {|w, s| by_score[s] = by_score[s].nil? ? [w] : by_score[s] + [w] }
    ordered = []
    by_score.keys.sort.reverse.each do |score|
        ordered += by_score[score].sort
    end
    return ordered.map{|w| "%s,%.3f" % [w, stats[w]/sum.to_f]}.join ";" 
end

File.open(ARGV[0]).each_line do |line|
    ngram_length, words = line.strip.split(",")
    puts transform(words, ngram_length.to_i)
end
