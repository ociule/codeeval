// This solution sometimes has 90% scores, so it fails on some rare inputs.

package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "math"

type signal []int

func parse(line string) (out [2000]int) {
	inSplit := strings.Fields(line)
	for ix, sval := range inSplit {
		var n int
		fmt.Sscanf(sval, "%d", &n)
		out[ix] = n
	}
	return
}

func (s *signal) average() float64 {
	sum := 0
	for _, sample := range *s {
		sum += sample
	}
	return float64(sum) / float64(len(*s))
}

func (s *signal) removeLinearBias() {
	var start signal = (*s)[0:1000]
	var end signal = (*s)[len(*s)-1000 : len(*s)]
	bias_start := start.average()
	bias_end := end.average()

	p1x, p2x := 500.0, 1500.0
	p1y, p2y := bias_start, bias_end

	// y = m * x + b
	m := (p2y - p1y) / (p2x - p1x)
	b := p1y - m*p1x
	for ix, sample := range *s {
		local_bias := float64(ix)*m + b
		(*s)[ix] = sample - int(local_bias)
		//if sample == 0 {
		//    fmt.Println("zero at ", ix)
		//}
	}
}

func (s *signal) dump(fileName string) {
	// Dumps the signal in a format useful for viewing with gnuplot, just: open "out.txt"
	fo, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for _, sample := range *s {
		fo.WriteString(strconv.FormatInt(int64(sample), 10))
		fo.WriteString("\n")
	}
}

func sameSign(x, y int) bool {
	return x*y >= 0
}

func (s *signal) findZeroCrossings() []int {
	// Signal must be already debiased!
	zc := make([]int, 0)
	ixLastNotZero := 0
	for ix, sample := range *s {
		sampleLastNotZero := (*s)[ixLastNotZero]
		if !sameSign(sample, sampleLastNotZero) {
			//fmt.Println("Found zc at", ix)
			ixLastNotZero = ix
			zc = append(zc, ix)
		}
	}
	return zc
}

func zcToPeriods(zcs []int) []int {
	ps := make([]int, 0)

	for ix := 2; ix < len(zcs); ix += 3 {
		p := zcs[ix] - zcs[ix-2]
		ps = append(ps, p)
	}
	return ps
}

func periodsToF(ps []int) int {
	sum := 0.0
	for _, p := range ps {
		sum += float64(p)
	}
	avgp := sum / float64(len(ps))

	p_us := avgp * 50 // 50 us is the intersample period - 20khz sampling rate

	f := (1000000.0 / p_us)

	intf := int(math.Floor(f/10+0.5)) * 10
	return int(intf)
}

func (s *signal) solve() (sol int) {
	s.removeLinearBias()
	zcs := s.findZeroCrossings()
	ps := zcToPeriods(zcs)
	//fmt.Println(ps)
	return periodsToF(ps)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//i := 0
	for scanner.Scan() {
		//i += 1
		line := strings.TrimSpace(scanner.Text())
		c := parse(line)
		var s signal = c[0:len(c)]
		//name := fmt.Sprintf("out%d.txt", i)
		//s.removeLinearBias()
		fmt.Println(s.solve())
		//s.dump(name)
	}
}
