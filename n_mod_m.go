package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

// N mod M
func mod(n, m int64) int64 {
	for n >= m {
		n = n - m
	}
	return n
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
		lineSplit := strings.Split(line, ",")
		n, _ := strconv.ParseInt(lineSplit[0], 10, 8)
		m, _ := strconv.ParseInt(lineSplit[1], 10, 8)
		fmt.Println(mod(n, m))
	}
}
