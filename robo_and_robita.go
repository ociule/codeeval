package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

const DEBUG = false

type Board struct {
	Width, Height, RX, RY int
	Visited               []bool
}

func InitBoard(in string) Board {
	b := Board{}
	_, _ = fmt.Sscanf(in, "%dx%d | %d %d", &b.Width, &b.Height, &b.RX, &b.RY)
	b.Visited = make([]bool, b.Width*b.Height)
	return b
}

func (b *Board) StartPos() (int, int) {
	return 1, b.Height
}

func (b *Board) MarkVisited(x, y int) {
	flatIndex := (x - 1) + (y-1)*b.Width
	b.Visited[flatIndex] = true
}

func (b *Board) OutsideOrVisited(x, y int) bool {
	if x < 1 || y < 1 || x > b.Width || y > b.Height {
		return true
	}
	flatIndex := (x - 1) + (y-1)*b.Width
	visited := b.Visited[flatIndex]
	return visited
}

func (b *Board) Rotate(dx, dy int) (int, int) {
	switch {
	case dx == 1:
		return 0, -1
	case dy == -1:
		return -1, 0
	case dx == -1:
		return 0, 1
	case dy == 1:
		return 1, 0
	}
	panic("Should never be here")
	return 0, 0
}

func (b *Board) Run() int {
	X, Y := b.StartPos()
	dX, dY := 1, 0

	nuts := 0

	if DEBUG {
		fmt.Println("Starting at", X, Y)
	}
	for { //i := 0; i < 100; i++ {
		nuts += 1
		b.MarkVisited(X, Y)
		if X == b.RX && Y == b.RY {
			return nuts
		}

		fX := X + dX
		fY := Y + dY

		if b.OutsideOrVisited(fX, fY) {
			dX, dY = b.Rotate(dX, dY)
			if DEBUG {
				fmt.Println("Rotating to", dX, dY)
			}
		}
		X += dX
		Y += dY
		if DEBUG {
			fmt.Println(X, Y)
		}
	}

	return -1
}

func (b *Board) PPrint() {
	for y := b.Height; y > 0; y-- {
		for x := 1; x <= b.Width; x += 1 {
		}
		fmt.Println("")
	}
}

func solve(in string) int {
	p := InitBoard(in)
	return p.Run()
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
