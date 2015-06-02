package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func uniqueElements(in string) string {
	inSplit := strings.Fields(in)
	outSlice := make([]string, 0, len(inSplit))
	for _, word := range inSplit {
		swapped := word[len(word)-1:len(word)] + word[1:len(word)-1] + word[0:1]
		outSlice = append(outSlice, swapped)
	}

	return strings.Join(outSlice, " ")
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
		fmt.Println(uniqueElements(line))
	}
}
