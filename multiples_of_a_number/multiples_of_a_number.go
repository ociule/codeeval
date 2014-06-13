package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
//import "math"

func linearDivision(x, n int) int {
    q := 0
    r := x - n * q
    for r >= n {
        q++
        r = x - n * q
    }
    return q
}

// What we want to compute is ceil(x / n) * n, but we can't use division, so we'll replace it with another algorithm.
func smallestMultiple(x, n int) int {
    quotient := linearDivision(x, n)
    if quotient * n < x {
        quotient ++
    }
    return quotient * n
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := scanner.Text()
		testSplit := strings.Split(test, ",")
        expected := -1
		xs, ns := testSplit[0], testSplit[1]
        x, _ := strconv.Atoi(xs)
        n, _ := strconv.Atoi(ns)
        if len(testSplit) == 3 {
        // expected exists
            es := testSplit[2]
            expected, _ = strconv.Atoi(es)
        }
        if expected >= 0 {
		    actual := smallestMultiple(x, n)
		    fmt.Println(actual, expected, actual == expected)
        } else {
		    fmt.Println(smallestMultiple(x, n))
        }
	}
}
