// My intuition, for k > 2, is to use the first k - 2 for a binary search.
// This will reduce the search domain
package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "math"

func parseTest(in string) (int, int) {
	k, n := 0, 0
	_, _ = fmt.Sscanf(in, "%d %d", &k, &n)
	return k, n
}

// How many floors can we test with d drops and k eggs ?
// See http://www.maa.org/sites/default/files/pdf/cmj_ftp/CMJ/November%202010/3%20Articles/1%20Denman/TowerAndMarbles.pdf
func nFloors(d, k int, cache *map[complex128]int) int {
	if d == 0 || k == 0 {
		return 0
	} else {
		cmp := complex128(complex(float64(d), float64(k)))
		vcache := *cache
		if val, present := vcache[cmp]; present {
			return val
		}
		return nFloors(d-1, k-1, cache) + nFloors(d-1, k, cache) + 1
	}
}

// If k == 2
// Solves d + (d - 1) + (d - 2) + (d - 3) + (d - 4) + ... + 1 = n
// This comes to d(d + 1) / 2 = n and then to
// d ** 2 + d - 2 * n = 0, so we solve this quadratic equoation to get d
// If k > 2, we check if it's large enough to do a binary search
//   If yes, we return the number of drops a binary search requires
//   If not, we use nFloors. d is the smalles number such that nFloors(d, k) >= n
func calcDrops(k, n int, cache *map[complex128]int) (d int) {
	need2binarySearch := int(math.Ceil(math.Log2(float64(n))))
	if k > need2binarySearch {
		return need2binarySearch
	}

	if k == 2 {
		//fmt.Println(">>>", k - 2, n, k)
		delta := 1 + 4*2*n
		floatD := (-1 + math.Sqrt(float64(delta))) / 2
		d = int(math.Ceil(floatD)) + (k - 2)
		return d
	}

	// If none the above, search for d such that nFloors(d, k) >= n
	//fmt.Println("Recurrence")
	d = 1
	for {
		nf := nFloors(d, k, cache)
		if nf >= n {
			return d
		}
		d += 1
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cache := make(map[complex128]int, 10000)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		//fmt.Println(line)
		k, n := parseTest(line)
		fmt.Println(calcDrops(k, n, &cache))
	}
}
