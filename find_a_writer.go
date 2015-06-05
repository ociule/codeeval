package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func solve(in string) string {
	inSplit := strings.Split(in, "|")

	cypher := inSplit[0]
	keySplit := strings.Fields(strings.TrimSpace(inSplit[1]))

	outSlice := make([]string, 0, len(keySplit))
	for _, si := range keySplit {
		i, _ := strconv.ParseInt(si, 10, 8)
		outSlice = append(outSlice, cypher[i-1:i])
	}
	return strings.Join(outSlice, "")
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
