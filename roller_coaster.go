package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "unicode"

func roller(in string) string {
	outSlice := make([]string, 0, len(in))

	startCase := true

	for _, char := range in {
		if unicode.IsLetter(char) {
			var charC string
			if startCase {
				charC = strings.ToUpper(string(char))
			} else {
				charC = strings.ToLower(string(char))
			}
			outSlice = append(outSlice, charC)
			startCase = !startCase
		} else {
			outSlice = append(outSlice, string(char))
		}
	}
	return strings.Join(outSlice, "")
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
		fmt.Println(roller(line))
	}
}
