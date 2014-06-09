package main

import "fmt"
import "log"
import "bufio"
import "os"
import "unicode"

func count(in string) string {
	upper_count := 0.0
	for _, char := range in {
		if unicode.IsUpper(char) {
			upper_count += 1
		}
	}
	len_in := float64(len(in))
	lower_perc := fmt.Sprintf("%2.2f", (len_in-upper_count)/len_in*100)
	upper_perc := fmt.Sprintf("%2.2f", (upper_count)/len_in*100)
	return fmt.Sprintf("lowercase: %s uppercase: %s", lower_perc, upper_perc)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(count(scanner.Text()))
	}
}
