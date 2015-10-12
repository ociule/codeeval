package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "math"

func solve(in string) int {
	l := len(in)
	doubleWildcard := 0

	for i := 0; i < l/2; i++ {
		a, b := in[i], in[l/2+i]

		//fmt.Println(i, string(a), string(b))
		if a != b {
			// At least one should be *
			if a != '*' && b != '*' {
				return 0
			}
		}
		if a == '*' && b == '*' {
			doubleWildcard += 1
		}
	}
	if doubleWildcard > 0 {
		return int(math.Pow(2, float64(doubleWildcard)))
	} else {
		return 1
	}
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
		fmt.Println(solve(line))
	}
}
