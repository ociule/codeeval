package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"


func sumDigits(in uint) (out uint) {
    for in >= 10 {
        out += in % 10
        in = in / 10
    }
    out += in
    return out
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
		fmt.Println(sumDigits(n))
	}
}
