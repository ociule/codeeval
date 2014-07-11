package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "errors"
import "strconv"
import "math"

const TERMINATOR = "1111111"

// Assumes the header does not contain 3 consecutive "0" or "1"'s
func findHeaderAndMessage(line string) (string, string, error) {
	start := 1
	for start < len(line) {
		ixFirst01 := start + strings.IndexAny(line[start:], "01")
		sec01 := line[ixFirst01+1]
		third01 := line[ixFirst01+1]
		if (sec01 == '0' || sec01 == '1') && (third01 == '0' || third01 == '1') {
			return line[0:ixFirst01], line[ixFirst01:], nil
		}
		start = ixFirst01 + 2
	}
	return "", "", errors.New("Couldn't find end of header")
}

// Generates all binary keys of length 7, exluding the invalid ones (all 1's)
func genKeys() []string {
	length := 7
	out := make([]string, 0, 247)

	for i := 1; i <= length; i++ {
		//fmt.Println("All keys of len", i)
		limit := int(math.Pow(2, float64(i)))
		for j := 0; j <= limit-1; j++ {
			ks := strconv.FormatInt(int64(j), 2)
			if ks != TERMINATOR[:i] {
				ks = strings.Repeat("0", i-len(ks)) + ks
				out = append(out, ks)
			}
		}
	}
	return out
}

func headerMapping() map[string]int {
	keys := genKeys()
	mapping := make(map[string]int)
	for ix, key := range keys {
		mapping[key] = ix
	}
	return mapping
}

func decodeMessage(header, encoded string, mapping map[string]int) string {
	pos := 0
	out := make([]string, 0)
	for pos < len(encoded) { // Main loop that parses segments
		key_length64, _ := strconv.ParseInt(encoded[pos:pos+3], 2, 0)
		key_length := int(key_length64)
		if key_length == 0 { // 000 or message end
			break
		}
		//fmt.Println("Segment with key length", key_length)
		pos += 3
		for true {
			key := encoded[pos : pos+key_length]
			pos += key_length
			if key == TERMINATOR[:key_length] { // 111... is this a segment end ?
				//fmt.Println("Segment terminated")
				break // Go back to beginning of main loop, parse next segment
			}
			out = append(out, string(header[mapping[key]]))
		}
	}
	return strings.Join(out, "")
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mapping := headerMapping()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		header, encoded, _ := findHeaderAndMessage(line)
		fmt.Println(decodeMessage(header, encoded, mapping))
	}
}
