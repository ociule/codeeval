package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "math"

const lower_bound = 0

type Challenge struct {
	UpperBound uint
	Answers    []bool
}

func parseChallenge(line string) *Challenge {
	inSplit := strings.Fields(line)
	var ub uint
	fmt.Sscanf(inSplit[0], "%d", &ub)
	str_answers := inSplit[1 : len(inSplit)-1]
	answers := make([]bool, len(str_answers))
	for ix, str_answer := range str_answers {
		answers[ix] = (str_answer != "Lower")
	}
	return &Challenge{ub, answers}

}

func getMedianOfRange(lower, upper uint) uint {
	return uint(math.Floor((float64(lower+upper))/2 + 0.5))
}

func (c *Challenge) solve() (sol uint) {
	lb := uint(0)
	ub := c.UpperBound
	for _, answer := range c.Answers {
		guess := getMedianOfRange(lb, ub)
		if answer {
			lb, ub = guess+1, ub
		} else {
			lb, ub = lb, guess-1
		}
	}
	finalGuess := getMedianOfRange(lb, ub)
	return finalGuess
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
		c := parseChallenge(line)
		fmt.Println(c.solve())
	}
}
