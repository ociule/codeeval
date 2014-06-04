go build 
time ./skyscrapers skyscrapers.txt > skyscrapers.txt.actual
diff skyscrapers.txt.actual skyscrapers.txt.expected
