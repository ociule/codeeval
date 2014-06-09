package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

var table = [...]int{
	0, 0, 0, 0, 1, 1, 1, 1, 6, 2, 4, 8, 1, 3, 9, 7, 6, 4, 6, 4,
	5, 5, 5, 5, 6, 6, 6, 6, 1, 7, 9, 3, 6, 8, 4, 2, 1, 9, 1, 9}

func get_last_digit_of_pow(base, pow int) int {
	return table[((base%10)<<2)+(pow&3)]
}

var cycle_len = [10]int{
	1, 1, 4, 4, 2, 1, 1, 4, 4, 2}

func digit_statistics(a, n int) [10]int {
	var stats [10]int
	cl := cycle_len[a%10]

	//fmt.Println(cl, n/cl, n%cl)
	// We could count the digit statistics by looping n times, but as they cycle every cl,
	// this is the equivalent:
	// Looping for a cycle (looping cl times)
	// Multiplying the statistics by n/cl, then looping again for n%cl
	for in := 1; in <= cl; in++ { // First let's run the stats for a cycle
		ld := get_last_digit_of_pow(a, in)
		//fmt.Println(a, in, ld)
		stats[ld] += 1
	}
	for ix, st := range stats { // Multiply each count by n/cl
		stats[ix] = st * (n / cl)
	}
	for in := 1; in <= n%cl; in++ { // Then loop again for n%cl
		ld := get_last_digit_of_pow(a, in)
		//fmt.Println(a, in, ld)
		stats[ld] += 1
	}
	return stats
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	s := make([]string, 10)
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		line := scanner.Text()
		an := strings.Split(line, " ")
		a, _ := strconv.Atoi(an[0])
		n, _ := strconv.Atoi(an[1])
		ds := digit_statistics(a, n)
		for ix, f := range ds {
			s[ix] = fmt.Sprintf("%d: %d", ix, f)
		}
		fmt.Println(strings.Join(s, ", "))
	}
}
