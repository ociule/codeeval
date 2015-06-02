package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "sort"

func simpleSorting(in string) string {
	inSplit := strings.Fields(in)
	outSlice := make([]float64, 0, len(inSplit))
	for _, n := range inSplit {
		fn, _ := strconv.ParseFloat(n, 64)
		outSlice = append(outSlice, fn)
	}

	fs := sort.Float64Slice(outSlice[0:])
	sort.Sort(fs)

	strOutSlice := make([]string, 0, len(fs))
	for _, n := range fs {
		sn := strconv.FormatFloat(n, 'f', 3, 64)
		strOutSlice = append(strOutSlice, sn)
	}

	return strings.Join(strOutSlice, " ")
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
		fmt.Println(simpleSorting(line))
	}
}
