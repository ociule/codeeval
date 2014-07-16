package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
//import "math"

type Node struct {
    Value int
    Left *Node
    Right *Node
}

func parse_test(test string) (int, int) {
    x, y := 0, 0
    fmt.Sscanf(test, "%d %d", &x, &y)
    return x, y
}

func (root *Node) findByValue(value int) (*Node) {
    for root != nil {
        if root.Value == value {
            return root
        } else {
            if root.Value > value {
                root = root.Left
            } else {
                root = root.Right
            }
        }
    }
    return root
}

func (root *Node) find_lowest_common_ancestor(x, y int) *Node {
    if (y < root.Value && root.Value < x) || y == root.Value || x == root.Value {
        return root
    } else {
        if y < root.Value {
            return root.Left.find_lowest_common_ancestor(x, y)
        } else {
            return root.Right.find_lowest_common_ancestor(x, y)
        }
    }
}

func max_min(x, y int) (int, int) {
    if x > y {
        return x, y
    } else {
        return y, x
    }
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

    tree := Node{30,
        &Node{8,
            &Node{3, nil, nil},
            &Node{20, &Node{10, nil, nil}, &Node{29, nil, nil}}},
        &Node{52, nil, nil}}

	for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        //fmt.Println(line)
		x, y := parse_test(line)
        x, y = max_min(x, y)
		fmt.Println(tree.find_lowest_common_ancestor(x, y).Value)
	}
}
