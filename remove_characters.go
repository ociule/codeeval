package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func parseTest(in string) (str string, chars map[rune]struct{}) {
    inSplit := strings.Split(in, ", ")
    str = inSplit[0]
    chars = make(map[rune]struct{}, len(inSplit[1]))
    for _, char := range inSplit[1] {
        chars[char] = struct{}{}
    }
	return
}

func remove(str string, chars map[rune]struct{}) string {
    //fmt.Println(str, chars)
    outSlice := make([]string, 0, len(str))
    for _, char := range str {
        _, present := chars[char]
        if !present {
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
		str, chars := parseTest(line)
		//fmt.Println(length, repChar, coded)
		fmt.Println(remove(str, chars))
	}
}
