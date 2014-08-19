package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sum uint = 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var n uint
		fmt.Sscanf(line, "%d", &n)
		sum += n
	}
	fmt.Println(sum)
}
