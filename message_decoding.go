package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "errors"
import "strconv"
import "math"

func findHeaderAndMessage(line string) (string, string, error) {
    start := 1
    for start < len(line) {
        ixFirst01 := start + strings.IndexAny(line[start:], "01")
        sec01 := line[ixFirst01 + 1]
        third01 := line[ixFirst01 + 1]
        if (sec01 == '0' || sec01 == '1') && (third01 == '0' || third01 == '1') {
            return line[0:ixFirst01], line[ixFirst01:], nil
        }
        start = ixFirst01 + 2
    }
    return "", "", errors.New("Couldn't find end of header")
}

// Generates all binary keys of length 7, exluding the invalid ones (all 1's)
func genKeys() {
    length := 7

    for i := 1; i <= length; i++ {
        fmt.Println("All keys of len", i)
        limit := math.Pow(2, int64(i))
        for j := 0; j <= limit - 1; j++ {
            fmt.Println(strconv.FormatInt(j, 2))
        }
    }
}

func headerToMap(header string) map[string]string {
    mapping := make(map[string]string)
    for ix, char := range header {
        fmt.Println(ix, char)
    }
    return mapping
}

func decodeMessage(header, encoded string) {
    terminator := "1111111"
    pos := 0
    for pos < len(encoded) { // Main loop that parses segments
        key_length64, _ := strconv.ParseInt(encoded[pos:pos+3], 2, 0)
        key_length := int(key_length64)
        if key_length == 0 { // 000 or message end
            break
        }
        fmt.Println("Segment with key length", key_length)
        pos += 3
        for true {
            key := encoded[pos:pos+key_length]
            pos += key_length
            if key == terminator[:key_length] { // 111... is this a segment end ?
                fmt.Println("Segment terminated")
                break // Go back to beginning of main loop, parse next segment
            }
            fmt.Println(key)
        }
    }
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    genKeys()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
        header, encoded, _ := findHeaderAndMessage(line)
        fmt.Println(header, encoded)
        decodeMessage(header, encoded)
        /**
		if still_tests {
			if strings.HasPrefix(line, "END OF INPUT") {
				still_tests = false
				continue
			}
			tests[line] = true
			tests_sorted = append(tests_sorted, line)
		} else {
			word := line
			for test, _ := range tests {
				if !present(sn[test], word) && Distance(word, test) == 1 {
					//fmt.Println(word, test)
					to_explore[word] = true
					sn[test] = append(sn[test], word)
				}
			}
		}*/
	}
}
