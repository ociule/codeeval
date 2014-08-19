package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func toBinString(in uint) string {
	return strconv.FormatUint(uint64(in), 2)
}

func bitPos(n, p1, p2 uint) bool {
	//fmt.Println(toBinString(n), p1, p2)

	bp1 := n >> (p1 - 1) & 1
	bp2 := n >> (p2 - 1) & 1
	return bp1 == bp2
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
		var n, p1, p2 uint
		fmt.Sscanf(line, "%d,%d,%d", &n, &p1, &p2)
		//fmt.Println(n, p1, p2)
		fmt.Println(bitPos(n, p1, p2))
	}
}
