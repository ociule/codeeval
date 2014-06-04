package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

// The Levenshtein distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
// http://en.wikipedia.org/wiki/Levenshtein_distance
//
// This implemention is optimized to use O(min(m,n)) space.
// It is based on the optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
func Distance(s1, s2 string) int {
	var cost, lastdiag, olddiag int
	len_s1 := len(s1)
	len_s2 := len(s2)

	column := make([]int, len_s1+1)

	for y := 1; y <= len_s1; y++ {
		column[y] = y
	}

	for x := 1; x <= len_s2; x++ {
		column[0] = x
		lastdiag = x - 1
		for y := 1; y <= len_s1; y++ {
			olddiag = column[y]
			cost = 0
			if s1[y-1] != s2[x-1] {
				cost = 1
			}
			column[y] = min(
				column[y]+1,
				column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return column[len_s1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func find_father(sn map[string][]string, test string) string {
	for father, sons := range sn {
		if present(sons, test) {
			return father
		}
	}
	return ""
}

func present(list []string, test string) bool {
	for _, possible := range list {
		if possible == test {
			return true
		}
	}
	return false
}

func get_graph_key(a string, b string) string {
	if a > b {
		return a + "::" + b
	} else {
		return b + "::" + a
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tests_sorted := make([]string, 0)
	tests := make(map[string]bool)
	explored := make(map[string]bool)
	to_explore := make(map[string]bool)
	still_tests := true
	words := make([]string, 24000)
	sn := make(map[string][]string)

	graph := make(map[string]int)
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		line := scanner.Text()
		if still_tests {
			if strings.HasPrefix(line, "END OF INPUT") {
				still_tests = false
				continue
			}
			tests[line] = true
			tests_sorted = append(tests_sorted, line)
		} else {
			word := strings.TrimSpace(line)
			words = append(words, word)
		}
	}
	lenw := len(words)
	counter := 0
	for _, word := range words {
		for _, word2 := range words {
			gkey := get_graph_key(word, word2)
			if _, present := graph[gkey]; !present {
				graph[gkey] = Distance(word, word2)
			}
		}
		counter += 1
		fmt.Println(counter, lenw)
	}

	//fmt.Println(to_explore)

	file.Close()
	for len(to_explore) > 0 {

		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
		still_tests = true
		to_explore_again := make(map[string]bool)
		for scanner.Scan() {
			//'scanner.Text()' represents the test case, do something with it
			line := scanner.Text()
			if still_tests {
				if strings.HasPrefix(line, "END OF INPUT") {
					still_tests = false
					continue
				}
			} else {
				word := line
				for test, _ := range to_explore {
					if word != test && Distance(word, test) == 1 && !tests[word] && !explored[word] {
						//fmt.Println(word, test)
						father := find_father(sn, test)
						if !present(sn[father], word) {
							to_explore_again[word] = true
							sn[father] = append(sn[father], word)
							//explored
						}
					}
				}
			}
		}
		for word, _ := range to_explore {
			explored[word] = true
		}
		to_explore = to_explore_again

		//fmt.Println(to_explore_again, len(to_explore))
	}
	//fmt.Println(sn)
	for _, word := range tests_sorted {
		fmt.Println(len(sn[word]) + 1)
	}
}
