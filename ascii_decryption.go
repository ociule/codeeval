package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "errors"

type RunePosCount struct {
	Pos, Count int
}

func parseTest(in string) (int, string, string) {
	inSplit := strings.Split(in, "|")
	rLength, _ := strconv.ParseInt(strings.TrimSpace(inSplit[0]), 10, 64)
	length := int(rLength)
	repeatedChar := strings.TrimSpace(inSplit[1])
	codedSplit := strings.Fields(strings.TrimSpace(inSplit[2]))
	coded := make([]string, 0, len(codedSplit))
	for _, rawCode := range codedSplit {
		code, _ := strconv.ParseInt(rawCode, 10, 64)
		coded = append(coded, string(code))
	}
	return length, repeatedChar, strings.Join(coded, "")
}

func findRepeated(in string, length int) ([]string, error) {
	if len(in) <= length {
		return []string{}, errors.New("len(in) <= len_")
	}
	prefixSet := make(map[string]struct{}, len(in)-length)
	prefixes := make([]string, 0, 10)
	for ix := 0; ix < len(in)-length; ix++ {
		prefix := in[ix : ix+length]
		_, present := prefixSet[prefix]
		if present {
			//fmt.Println("found", prefix)
			prefixes = append(prefixes, prefix)
		}
		prefixSet[prefix] = struct{}{}
	}
	return prefixes, nil
}

func rot13(in string, delta int) string {
	out := make([]string, 0, len(in))
	for _, coded := range in {
		out = append(out, string(int(coded)-delta))
	}
	return strings.Join(out, "")
}

func stringIsAsciiPrintable(in string) bool {
	for _, char := range in {
		if int(char) < 32 || int(char) > 122 {
			return false
		}
	}
	return true
}

func decode(length int, repChar, coded string) string {
	repeated, _ := findRepeated(coded, length)
	for _, candidate := range repeated {
		codedRepChar := candidate[length-1]
		delta := codedRepChar - repChar[0]
		decoded := rot13(coded, int(delta))
		//fmt.Println(delta, candidate, codedRepChar, repChar, repChar[0], decoded, stringIsAsciiPrintable(decoded))
		if stringIsAsciiPrintable(decoded) {
			return decoded
		}
	}
	return ""
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		//fmt.Println(line)
		length, repChar, coded := parseTest(line)
		//fmt.Println(length, repChar, coded)
		fmt.Println(decode(length, repChar, coded))
	}
}
