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
	for scanner.Scan() {
		lineSplit := strings.Fields(scanner.Text())
        rev := make([]string, len(lineSplit))
        for i := len(lineSplit) - 1; i >= 0; i-- {
            rev[i] = lineSplit[len(lineSplit) - i - 1]
        }
        fmt.Println(strings.Join(rev, " "))
	}
}
