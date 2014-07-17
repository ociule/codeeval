package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "math"
import "strconv"
import "sort"

type Hole struct {
	X1, Y1, X2, Y2 int
}

type Brick struct {
	Index, X1, Y1, Z1, X2, Y2, Z2 int
}

func parse_test(test string) (hole Hole, bricks []Brick) {
	stest := strings.Split(test, "|")
	shole, sbricks := stest[0], strings.Split(stest[1], ";")
	x1, y1, x2, y2 := 0, 0, 0, 0
	fmt.Sscanf(shole, "[%d,%d] [%d,%d]", &x1, &y1, &x2, &y2)
	hole = Hole{x1, y1, x2, y2}
	bricks = make([]Brick, 0, len(sbricks))
	for _, sbrick := range sbricks {
		index, x1, y1, z1, x2, y2, z2 := 0, 0, 0, 0, 0, 0, 0
		fmt.Sscanf(sbrick, "(%d [%d,%d,%d] [%d,%d,%d])", &index, &x1, &y1, &z1, &x2, &y2, &z2)
		brick := Brick{index, x1, y1, z1, x2, y2, z2}
		bricks = append(bricks, brick)
	}
	return hole, bricks
}

func (h *Hole) SideLengths() (int, int) {
	xd := int(math.Abs(float64(h.X2 - h.X1)))
	yd := int(math.Abs(float64(h.Y2 - h.Y1)))
	if xd > yd {
		return xd, yd
	} else {
		return yd, xd
	}
}

func (h *Hole) Fits(hole Hole) bool {

	h1L, h1l := h.SideLengths()
	h2L, h2l := hole.SideLengths()
	//fmt.Println("sides", h1L, h1l, h2L, h2l)
	if (h1L <= h2L) && (h1l <= h2l) {
		//fmt.Println(h, "fits?", hole, " ", h1L, h1l, h2L, h2l)
		return true
	}
	return false
}

func (b *Brick) Fits(hole Hole) bool {
	// 3 faces:
	// x1, z1 x2, z2
	// x1, y1 x2, y2
	// y1, z1 y2, z2

	f1 := Hole{b.X1, b.Z1, b.X2, b.Z2}
	f2 := Hole{b.X1, b.Y1, b.X2, b.Y2}
	f3 := Hole{b.Y1, b.Z1, b.Y2, b.Z2}

	faces := []Hole{f1, f2, f3}

	for _, face := range faces {
		if face.Fits(hole) {
			return true
		}
	}
	return false
}

func test_brick(hole Hole, bricks []Brick) {
	s := make([]int, 0, len(bricks))
	for _, brick := range bricks {
		if brick.Fits(hole) {
			s = append(s, brick.Index)
		}
	}
	sort.Sort(sort.IntSlice(s))
	if len(s) > 0 {
		ss := make([]string, 0, len(s))
		for _, ix := range s {
			ss = append(ss, strconv.FormatInt(int64(ix), 10))
		}
		fmt.Println(strings.Join(ss, ","))
	} else {
		fmt.Println("-")
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
		line := strings.TrimSpace(scanner.Text())
		//fmt.Println(line)
		test_brick(parse_test(line))
	}
}
