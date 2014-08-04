package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

type RunePosCount struct {
    Pos, Count int
}

func getFirst(in string) string {
    nonRepeated := make(map[rune]RunePosCount, len(in))
    for ix, char := range in {
        val, ok := nonRepeated[char]
        if !ok {
            nonRepeated[char] = RunePosCount{ix, 1}
        } else {
            nonRepeated[char] = RunePosCount{val.Pos, val.Count + 1}
            //fmt.Println(string(char), ix, val.Count + 1)
        }
    }

    //fmt.Println(nonRepeated)
    lowestPos := len(in)
    out := rune(0)
    for char, posCount := range nonRepeated {
        if posCount.Count == 1 && posCount.Pos < lowestPos {
            out = char
            lowestPos = posCount.Pos
        }
    }
    return string(out)
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
		fmt.Println(getFirst(line))
	}
}
