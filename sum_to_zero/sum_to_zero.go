package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func sumToZero(x []int) int {
    countz := 0
    for i := 0; i < len(x); i++ {
        for j := i + 1; j < len(x); j++ {
            for k := j + 1; k < len(x); k++ {
                for l := k + 1; l < len(x); l++ {
                    sum := x[i] + x[j] + x[k] + x[l]
                    if sum == 0 {
                        countz += 1
                    }
                }
            }
        }
    }
    return countz
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testString := scanner.Text()
		testSplit := strings.Split(testString, ",")
        test := make([]int, 0, len(testSplit))
        for _, t := range testSplit {
            x, _ := strconv.Atoi(t)
            test = append(test, x)
        }
		fmt.Println(sumToZero(test))
	}
}
