package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func restitute(scrambled []string, orders []int) string {
	decyphered := make([]string, len(scrambled))
	for ix, or := range orders {
		decyphered[ix] = scrambled[or]
	}
	return strings.Join(decyphered, " ")
}

func FindInt(s []int, i int) int {
	for ix, val := range s {
		if i == val {
			return ix
		}
	}
	return -1
}

func fillOrders(orders []int, l int) []int {
	filled := make([]int, l)
	for i := 1; i <= l; i++ {
		pos := FindInt(orders, i)
		if pos >= 0 {
			filled[i-1] = pos
		} else {
			filled[i-1] = l - 1
		}
	}
	return filled
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
		tests := strings.Split(test, ";")
		orderss := strings.Fields(tests[len(tests)-1])
		orders := make([]int, len(orderss))
		for ix, or := range orderss {
			orders[ix], _ = strconv.Atoi(or)
		}
		scrambled := strings.Fields(strings.Join(tests[0:len(tests)-1], ";"))
		orders = fillOrders(orders, len(scrambled))
		fmt.Println(restitute(scrambled, orders))
	}
}
