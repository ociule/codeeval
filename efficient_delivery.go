// This is the change-making problem

package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func parseTest(in string) ([]int, int) {
	inSplit := strings.Split(in, "), ")
	n64, _ := strconv.ParseInt(inSplit[1], 10, 32)
	tankersRaw := inSplit[0][1:]
	tankersSplit := strings.Split(tankersRaw, ",")
	tankers := make([]int, len(tankersSplit))
	for ix, tRaw := range tankersSplit {
		t64, _ := strconv.ParseInt(tRaw, 10, 32)
		tankers[ix] = int(t64)
	}

	return tankers, int(n64)
}

func mergeM(coins []int, m int, sol []int) []int {
	fmt.Println("Merging", m, coins[m-1], sol)
	if sol != nil {
		if len(sol) == len(coins) {
			sol[(m - 1)] += 1
		} else {
			c := 0
			for c < (len(sol) / len(coins)) {
				sol[len(coins)*c+m-1] += 1
				c += 1
			}
		}
		fmt.Println("Merging out", sol)
	}

	return sol
}

func count(coins []int, m, n int) []int {
	fmt.Println(coins, m, n)
	// If n is less than 0 then no solution exists
	if n < 0 {
		return nil
	}

	// If n is 0 then there is 1 solution (do not include any coin)
	if n == 0 {
		fmt.Println("n == 0", m)
		return make([]int, len(coins))
	}

	if m == 1 {
		if n%coins[0] == 0 {
			sol := make([]int, len(coins))
			sol[0] = n / coins[0]
			fmt.Println("m == 1", n, sol)
			return sol
		} else {
			fmt.Println("m == 1", n, nil)
			return nil
		}
	}

	// If there are no coins and n is greater than 0, then no solution exist
	if m <= 0 && n >= 1 {
		return nil
	}

	// count is sum of solutions (i) including coins[m-1] (ii) excluding S[m-1]
	return append(count(coins, m-1, n), mergeM(coins, m, count(coins, m, n-coins[m-1]))...)
}

func countAndHandleNotFound(coins, n) []int {
	sols := count(coins, len(coins), n)
	if len(sols) == 0 { // Hey, we must search for the minimum amount of oil which needs to be added

	}
	return sols
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
		coins, n := parseTest(line)
		//fmt.Println(n, p1, p2)
		fmt.Println(countAndHandleNotFound(coins, n))
	}
}
