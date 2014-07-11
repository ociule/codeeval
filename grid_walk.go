package main

import "fmt"
import "container/list"

type Vertex struct {
	X int
	Y int
}

const LIMIT = 19

func digit_sum(n int) int {
	if n < 0 {
		n = n * -1
	}
	sum := 0
	d, r := 0, 0
	for n > 0 {
		d = n / 10
		r = n % 10
		sum += r
		n -= r
		if d > 0 {
			n /= 10
		}
	}
	return sum
}

func accessible(v Vertex) bool {
	return digit_sum(v.X)+digit_sum(v.Y) <= LIMIT
}

func explore(visited map[Vertex]bool, queue *list.List, v Vertex) {
	_, ok := visited[v]
	if accessible(v) && !ok {
		visited[v] = true
		queue.PushBack(v)
	}
}

func main() {
	start := Vertex{0, 0}
	visited := make(map[Vertex]bool)
	queue := list.New()

	visited[start] = true
	queue.PushBack(start)

	for queue.Len() > 0 {
		next := queue.Front()
		queue.Remove(next)
		nv := next.Value.(Vertex)
		x, y := nv.X, nv.Y

		explore(visited, queue, Vertex{x - 1, y})
		explore(visited, queue, Vertex{x + 1, y})
		explore(visited, queue, Vertex{x, y - 1})
		explore(visited, queue, Vertex{x, y + 1})
	}

	fmt.Println(len(visited))
}
