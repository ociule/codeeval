package main

import "fmt"
import "log"
import "bufio"
import "os"
import "unicode"
import "sort"

// Start of Counter module - this should be in a different import but no easy way to import my own module in codeeval
// This Counter is inspired by python's collections.Counter
type Counter struct {
	Data map[interface{}]int
}

func (c *Counter) Add(o interface{}) {
	if c.Data == nil {
		c.Data = make(map[interface{}]int)
	}
	c.Data[o] += 1
}

type ItemWithCount struct {
	Item  interface{}
	Count int
}

// Used by MostCommon for sorting the list of ItemWithCount
type ByCount []ItemWithCount

func (this ByCount) Len() int {
	return len(this)
}

func (this ByCount) Less(i, j int) bool {
	return this[i].Count < this[j].Count
}

func (this ByCount) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// Returns a list of the n most common elements, from most common to least. If howMany is 0, MostCommon returns all the elements.
func (c *Counter) MostCommon(howMany int) []ItemWithCount {
	if howMany == 0 {
		howMany = len(c.Data)
	}
	out := make([]ItemWithCount, 0, len(c.Data))
	for k, v := range c.Data {
		out = append(out, ItemWithCount{k, v})
	}
	sort.Sort(sort.Reverse(ByCount(out)))
	return out[0:howMany]
}

// End of counter module

func beautyScore(line string) int {
	freqs := make(map[rune]int)
	c := Counter{}
	for _, char := range line {
		if unicode.IsLetter(char) {
			char = unicode.ToLower(char)
			freqs[char] += 1
			c.Add(char)
		}
	}
	mc := c.MostCommon(0)
	score := 26 // We know we must start at 26 and then go down
	beauty := 0
	for ix, item := range mc {
		count := item.Count
		beauty += count * (score - ix)
	}
	return beauty
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(beautyScore(line))
	}
}
