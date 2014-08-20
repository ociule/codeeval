package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func countOnes(in string) (out uint) {
	for _, char := range in {
		if char == '1' {
			out += 1
		}
	}
	return
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
		var n uint
		fmt.Sscanf(line, "%d", &n)
		//fmt.Println(n, p1, p2)
		binStr := strconv.FormatUint(uint64(n), 2)
		fmt.Println(countOnes(binStr))
	}
}
