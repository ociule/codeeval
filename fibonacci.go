package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func fib(n uint, cache []uint) uint {
	// This is the same as fibonacci(n), actually
	// We should do it in closed form
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	if cache[n] > 0 {
		return cache[n]
	} else {
		fibN := fib(n-2, cache) + fib(n-1, cache)
		cache[n] = fibN
		return fibN
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cache := make([]uint, 1000)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var n uint
		fmt.Sscanf(line, "%d", &n)
		fmt.Println(fib(n, cache))
	}
}
