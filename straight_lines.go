package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

const DEBUG = true

type Point struct {
	X, Y int
}

type Line struct {
	A, B Point
}

func parse_test(test string) []Point {
	stest := strings.Split(test, "|")

	points := make([]Point, 0, len(stest))
	for _, spoint := range stest {
		x, y := 0, 0
		fmt.Sscanf(strings.TrimSpace(spoint), "%d %d", &x, &y)
		p := Point{X: x, Y: y}
		points = append(points, p)
	}
	return points
}

const EPS = 1

func (l *Line) HasPoint(c Point) bool {
	// (x2 - x1) * (y3 - y1) = (x3 - x1) * (y2 - y1)
	if c == l.A || c == l.B {
		return false
	}
	d1 := (l.B.X - l.A.X) * (c.Y - l.A.Y)
	d2 := (c.X - l.A.X) * (l.B.Y - l.A.Y)
	onLine := d1 == d2
	return onLine
}

func test(points []Point) {
	knownLines := make([]Line, 0, len(points)*4)
	knownLinesWith3Points := make(map[Line]bool, len(knownLines))
	for ixP1, p1 := range points[:len(points)-1] {
		if DEBUG {
			fmt.Println("Considering p1 ", p1)
		}
		for _, p2 := range points[ixP1+1:] {
			if DEBUG {
				fmt.Println("Considering p2", p2)
			}
			onAtLeastOneLine := false
			for _, line := range knownLines {
				onLine := line.HasPoint(p2)
				onAtLeastOneLine = onAtLeastOneLine || onLine
				if onLine {
					if DEBUG {
						fmt.Println("Found line with 3p ", p2, line, p1)
					}
					knownLinesWith3Points[line] = true
				}
			}
			if !onAtLeastOneLine {
				knownLines = append(knownLines, Line{A: p1, B: p2})
				if DEBUG {
					fmt.Println("Found new line", Line{A: p1, B: p2})
				}
			}
		}
	}
	fmt.Println(len(knownLinesWith3Points))
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
		test(parse_test(line))
	}
}
