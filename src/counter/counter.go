// This implements a counter similar to python's collections.Counter

package counter

import "sort"

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
