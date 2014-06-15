package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func sumToZero(c chan int, x []int) {
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
    c <- countz
}


// This is a concurrent solution that runs sumToZero for each test in a separate goroutine
// It's useless on codeeval, because codeeval limits GOMAXPROCS to 1. Check it out with runtime.GoMaxProcs(0)
func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    tests := make([][]int, 0)
    c := make(chan int)
	for scanner.Scan() {
		testString := scanner.Text()
		testSplit := strings.Split(testString, ",")
        test := make([]int, 0, len(testSplit))
        for _, t := range testSplit {
            x, _ := strconv.Atoi(t)
            test = append(test, x)
        }
        tests = append(tests, test)
        go sumToZero(c, test)
	}
    for _, _ = range tests {
        countz := <-c
        fmt.Println(countz)
    }
}
