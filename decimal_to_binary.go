package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

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
		fmt.Println(strconv.FormatUint(uint64(n), 2))
	}
}
