package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "math"
import "sort"

type Vertex struct {
	X int
	Y int
}

func parse_test(test string) []Vertex {
	svertices := strings.Split(test, "),")
	vertices := make([]Vertex, 4)
	for ix, v := range svertices {
		v = strings.TrimSpace(v)
		v = strings.TrimPrefix(v, "(")
		v = strings.TrimSuffix(v, ")")
		xy := strings.Split(v, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		vertex := Vertex{x, y}
		vertices[ix] = vertex
	}
	return vertices
}

func get_len(a, b Vertex) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

const e = 1.0 / 1000000000

func test_float_eq(a, b float64) bool {
	return (math.Abs(a - b)) < e
}
func test_sq_tri(points []Vertex) (Vertex, bool) {
	a, b, c := points[0], points[1], points[2]
	ab := get_len(a, b)
	bc := get_len(b, c)
	ac := get_len(a, c)

	sides := []float64{ab, bc, ac}
	sort.Float64s(sides)

	longest := sides[2]
	sqrt2 := math.Sqrt(float64(2.0))
	if longest == 0 || sides[0] != sides[1] || !test_float_eq(sides[0]*sqrt2, longest) {
		return Vertex{}, false
	} else {
		switch longest {
		case ab:
			return c, true
		case ac:
			return b, true
		case bc:
			return a, true
		}
		return Vertex{}, false
	}
}

func test_square(points []Vertex) bool {
	sq_vertex, sq_tri := test_sq_tri(points[0:3])
	if !sq_tri {
		return false
	} else {
		var others []Vertex
		a, b, c := points[0], points[1], points[2]
		switch sq_vertex {
		case a:
			others = []Vertex{b, c}
		case b:
			others = []Vertex{a, c}
		case c:
			others = []Vertex{a, b}
		}
		others = append(others, points[3])
		sq_vertex, sq_tri = test_sq_tri(others)
		if sq_tri && sq_vertex == points[3] {
			return true
		} else {
			return false
		}
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		fmt.Println(test_square(parse_test(scanner.Text())))
	}
}
