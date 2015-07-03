package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

const ARROW_LEFT = ">>-->"
const ARROW_RIGHT = "<--<<"

func solve(in string) int {
	count := 0
	min_length := len(ARROW_LEFT)
	for ixr, r := range in {
		if len(in)-ixr < min_length {
			break
		}
		//fmt.Println(">", len(in)-ixr, min_length, in[ixr:ixr+min_length])
		if (r == rune(ARROW_LEFT[0]) && in[ixr:ixr+min_length] == ARROW_LEFT) || (r == rune(ARROW_RIGHT[0]) && in[ixr:ixr+min_length] == ARROW_RIGHT) {
			count += 1
		}
	}

	return count
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
		fmt.Println(solve(line))
	}
}
