package main

import "fmt"
import "strconv"
import "os"

func is_prime(n int) bool {
    for i := 2; i < n; i++ {
        if n%i == 0 { 
            return false
        }
    }
    return true
}

func is_palindrome(n int) bool {
    sn := strconv.Itoa(n)
    for i := 0; i < len(sn) / 2; i++ {
        if sn[i] != sn[len(sn)-i-1] {
            return false
        }
    }
    return true
}

func main() {
    const max = 1000
	for i := max - 1; i > 1; i--  {
        if is_prime(i) && is_palindrome(i) {
            fmt.Println(i)
            break
        }
	}
}
