package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func parseTest(in string) (n, m int) {
	_, _ = fmt.Sscanf(in, "%d %d", &n, &m)
	return
}

// Naive because I'm sure there's a better algo, linear (in n)
func naiveLockCount(n, m int) (count int) {
	locks := make([]bool, n)
	for i := 0; i < m-1; i++ {
		// First we close the even-numbered locks
		for pos := 1; pos < n; pos += 2 {
			locks[pos] = true
		}
		// Then we switch every 3rd lock
		for pos := 2; pos < n; pos += 3 {
			locks[pos] = !locks[pos]
		}
	}
	// Last pass, we switch the last door
	locks[n-1] = !locks[n-1]

	// Let's count
	for pos := 0; pos < n; pos++ {
		if !locks[pos] {
			count += 1
		}
	}

	return
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
		n, m := parseTest(line)
		fmt.Println(naiveLockCount(n, m))
	}
}
