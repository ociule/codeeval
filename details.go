package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

var DEBUG = false

type Detail struct {
	lines []string
	size  int
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (d *Detail) Init(lines string) {
	splitLines := strings.Split(lines, ",")
	d.size = len(splitLines[0])
	d.lines = splitLines

	if DEBUG {
		fmt.Println(lines)
	}
}

func (d *Detail) CheckEndsAreCovered() {
	for _, line := range d.lines {
		if line[0] != 'X' || line[len(line)-1] != 'Y' {
			fmt.Println("Uncovered ends", line)
		}
	}
}

func (d *Detail) CountSpaceBetweenDetails() int {
	minSpace := 1000000
	for _, line := range d.lines {
		rightMostX := strings.LastIndex(line, "X")
		leftMostY := strings.Index(line, "Y")

		if DEBUG {
			fmt.Println("processing", line, line[rightMostX+1:leftMostY])
		}
		space := leftMostY - rightMostX - 1
		if space < minSpace {
			if DEBUG {
				fmt.Println("found new min", space)
			}
			minSpace = space
		}
	}
	return minSpace
}

func solve(line string) int {
	detail := new(Detail)
	detail.Init(line)

	return detail.CountSpaceBetweenDetails()
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
		fmt.Println(solve(line))
	}
}
