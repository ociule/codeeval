package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func uniqueElements(in string) string {
	inSplit := strings.Split(in, ",")
	presenceCheck := make(map[string]bool, len(inSplit))
	outSlice := make([]string, 0, len(inSplit))
	for _, elem := range inSplit {
		_, present := presenceCheck[elem]
		if !present {
			presenceCheck[elem] = true
			outSlice = append(outSlice, elem)
		}
	}

	return strings.Join(outSlice, ",")
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
		fmt.Println(uniqueElements(line))
	}
}
