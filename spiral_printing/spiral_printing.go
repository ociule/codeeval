package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

var DEBUG = false

func parseMatrix(n, m int, raw string) [][]string {
	matrix := make([][]string, n)
	rawSplit := strings.Fields(raw)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = rawSplit[i*m+j]
		}
	}
	return matrix
}

func spiral(matrix [][]string, oy, ox, ey, ex int) []string {
	if ex <= 0 || ey <= 0 {
		fmt.Println("EMPTY", oy, ox, ey, ex)
		return []string{}
	}
	//fmt.Println(oy, ox, ey, ex, matrix)
	out := make([]string, 0)
	// Top side
	for i := ox; i < ex; i++ {
		if DEBUG {
			fmt.Println("top", matrix[oy][i])
		}
		out = append(out, matrix[oy][i])
	}
	if ey > 1 {
		// Right side
		for i := oy + 1; i < ey - 1; i++ {
			if DEBUG {
				fmt.Println("right", matrix[i][ex-1])
			}
			out = append(out, matrix[i][ex-1])
		}
		// Bottom side
		for i := ex - 1; i >= 0; i-- {
			if DEBUG {
				fmt.Println("bottom", matrix[ey-1][i])
			}
			out = append(out, matrix[ey-1][i])
		}
		if ey > 2 {
			if ex > 1 {
				// Left side
				for i := ey - 2; i > 0; i-- {
					if DEBUG {
						fmt.Println("left", matrix[i][ox])
					}
					out = append(out, matrix[i][ox])
				}
			}
		}
	}
	if ex >= 3 && ey >= 3 {
		moreMatrix := subMatrix(matrix, oy+1, ox+1, ey-2, ex-2)
		more := spiral(moreMatrix, 0, 0, len(moreMatrix), len(moreMatrix[0]))
		if len(more) > 0 {
			for _, elem := range more {
				out = append(out, elem)
			}
		}
	}
	return out
}

func subMatrix(matrix [][]string, oy, ox, ey, ex int) [][]string {
	if DEBUG {
		fmt.Println("MM", oy, ox, ey, ex, matrix)
	}
	out := make([][]string, ey)
	for i := 0; i < ey; i++ {
		out[i] = make([]string, ex)
		for j := 0; j < ex; j++ {
			out[i][j] = matrix[oy+i][ox+j]
		}
	}
	if DEBUG {
		fmt.Println(out)
	}
	return out
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), ";")
		n, err := strconv.Atoi(lineSplit[0])
        if err != nil {
            log.Fatal(err)
        }
		m, err := strconv.Atoi(lineSplit[1])
        if err != nil {
            log.Fatal(err)
        }
		matrix := parseMatrix(n, m, lineSplit[2])
		out := spiral(matrix, 0, 0, n, m)
		if len(out) != m*n {
			fmt.Println(matrix, out)
			panic(fmt.Sprint("len(out)=", len(out), "!= m * n", m*n))
		}
		fmt.Println(strings.Join(out, " "))
	}
}
