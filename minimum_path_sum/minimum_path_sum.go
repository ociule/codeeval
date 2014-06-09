package main

import "fmt"
import "log"
import "bufio"
import "strconv"
import "os"
import "strings"
import "math"

func ScanMultipleVariableLineTests(scanner *bufio.Scanner) map[int][]string {
	tests := make(map[int][]string)
	current_test_left := 0
	test_count := 0
	for scanner.Scan() {
		if current_test_left > 0 {
			current_test_left -= 1
			line := scanner.Text()
			tests[test_count-1] = append(tests[test_count-1], line)
		} else {
			current_test_left, _ = strconv.Atoi(scanner.Text())
			tests[test_count] = make([]string, 0, current_test_left)
			test_count += 1
		}
	}
	return tests
}

func parseTests(testsRaw map[int][]string) [][][]int {
	tests := make([][][]int, len(testsRaw))
	for ixTest, testRaw := range testsRaw {
		test := make([][]int, len(testRaw))
		tests[ixTest] = test
		for ixRow, lineRaw := range testRaw {
			line := make([]int, len(testRaw))
			tests[ixTest][ixRow] = line
			lineSplit := strings.Split(lineRaw, ",")
			for ixColumn, val := range lineSplit {
				tests[ixTest][ixRow][ixColumn], _ = strconv.Atoi(val)
			}
		}
	}
	return tests
}

func minimumSum2(matrix [][]int) int {
	size := len(matrix)
	for i := size - 2; i >= 0; i-- {
		matrix[size-1][i] += matrix[size-1][i+1]
		matrix[i][size-1] += matrix[i+1][size-1]
	}

	for i := size - 2; i >= 0; i-- {
		for j := size - 2; j >= 0; j-- {
			matrix[i][j] += int(math.Min(float64(matrix[i+1][j]), float64(matrix[i][j+1])))
		}
	}
	return matrix[0][0]
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	testsRaw := ScanMultipleVariableLineTests(scanner)
	tests := parseTests(testsRaw)
	for _, test := range tests {
		fmt.Println(minimumSum2(test))
	}
}
