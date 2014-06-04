package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

type Building struct {
	L int
	H int
	R int
}

func parseTest(test string) []Building {
	sbuildings := strings.Split(test, ");")
	buildings := make([]Building, len(sbuildings))
	for ix, v := range sbuildings {
		v = strings.TrimSpace(v)
		v = strings.TrimPrefix(v, "(")
		v = strings.TrimSuffix(v, ")")
		lhr := strings.Split(v, ",")
		l, _ := strconv.Atoi(lhr[0])
		h, _ := strconv.Atoi(lhr[1])
		r, _ := strconv.Atoi(lhr[2])
		sq := Building{l, h, r}
		buildings[ix] = sq
	}
	return buildings
}

func skylineMark(skyline []int, i int, h int) {
        if skyline[i] >= h {
            return
        }
        skyline[i] = h
}

func getSkyline(buildings []Building) string {
    // What is the start and length of this skyline ?
    skylineLength := 0
    skylineL := 100000
	for _, bld := range buildings {
        if bld.L < skylineL {
            skylineL = bld.L
        }
        if bld.R > skylineLength {
            skylineLength = bld.R
        }
    }
    // Our skyline is represented as a slice of integer new heights
    // New heights mean the height 2 at x=1 means the skyline has height 2 starting at x=1
    // Whether this is an closed or open interval is not clear but not necessary to give a solution
    skyline := make([]int, skylineLength - skylineL + 1)
	for _, bld := range buildings {
        for i := bld.L; i <= bld.R - 1; i++ {
            skylineMark(skyline, i - skylineL, bld.H)
        }
    }

    oldValue := 0
    output := make([]string, 0)
    for ix, v := range skyline {
        if v != oldValue {
            output = append(output, fmt.Sprintf("%d %d", ix + skylineL, v))
            oldValue = v
        }
    }
    //fmt.Println(output)
    // Now we output the skyline in the demanded format
    // Just remove duplicates by ignoring non-changing heights
    return strings.Join(output, " ")
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(getSkyline(parseTest(strings.TrimSpace(scanner.Text()))))
	}
}
