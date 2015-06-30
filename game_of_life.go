package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

var DEBUG = false

type Grid struct {
	grid [][]bool
	N    int
	B    int
	S1   int
	S2   int
}

const ALIVE = '*'
const DEAD = '.'

func (g *Grid) Init(n, b, s1, s2 int) {
	g.N = n
	g.B = b
	g.S1 = s1
	g.S2 = s2
	g.grid = initGrid(n)
}

func initGrid(n int) [][]bool {
	grid := make([][]bool, n)
	for i := range grid {
		grid[i] = make([]bool, n)
	}
	return grid
}

func (g *Grid) String() string {
	out := make([]string, 0, g.N*(g.N+1))
	for _, line := range g.grid {
		for _, cell := range line {
			out = append(out, g.CellToString(cell))
		}
		out = append(out, "\n")
	}
	return strings.Join(out, "")
}

func (g *Grid) CellToString(cell bool) string {
	if cell {
		return string(ALIVE)
	} else {
		return string(DEAD)
	}
}

func (g *Grid) AddLine(line string, currentLine int) {
	for ix, cell := range line {
		g.grid[currentLine][ix] = (cell == ALIVE)
	}
}

func (g *Grid) CountAliveNeighbours(line, col int) (n int) {
	if line > 0 {
		if col > 0 {
			if g.grid[line-1][col-1] {
				n += 1
			}
		}
		if g.grid[line-1][col] {
			n += 1
		}
		if col < g.N-1 {
			if g.grid[line-1][col+1] {
				n += 1
			}
		}
	}
	if col > 0 && g.grid[line][col-1] {
		n += 1
	}
	if col < g.N-1 && g.grid[line][col+1] {
		n += 1
	}
	if line < g.N-1 {
		if col > 0 && g.grid[line+1][col-1] {
			n += 1
		}
		if g.grid[line+1][col] {
			n += 1
		}
		if col < g.N-1 && g.grid[line+1][col+1] {
			n += 1
		}
	}
	return
}

func (g *Grid) Step() {
	nextGen := initGrid(g.N)
	for pLine, line := range g.grid {
		for pCol, cell := range line {
			nAlive := g.CountAliveNeighbours(pLine, pCol)
			if cell { // Alive cell
				if nAlive == g.S1 || nAlive == g.S2 {
					nextGen[pLine][pCol] = true
				}
			} else {
				if nAlive == g.B { // Should be born ?
					nextGen[pLine][pCol] = true
				}
			}
		}
	}

	g.grid = nextGen
}

func (g *Grid) StepN(n int) {
	count := 0
	for count < n {
		count += 1
		g.Step()
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grid = new(Grid)

	currentLine := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if grid.N == 0 {
			// Implement B3S23 game-of-life rules
			// Birth if 3 alive neighbours
			// Survive if 2 or 3 alive neigbours
			grid.Init(len(line), 3, 2, 3)
		}
		grid.AddLine(line, currentLine)
		currentLine += 1
	}
	grid.StepN(10)
	fmt.Println(grid)
}
