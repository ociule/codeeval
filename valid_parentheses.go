package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

var DEBUG = false

type ParenthesesStack struct {
	top  *Element
	size int
}

type Element struct {
	value rune
	next  *Element
}

func isClosingParen(val rune) bool {
	return val == ')' || val == ']' || val == '}'
}

func (s *ParenthesesStack) CanClose(val rune) bool {
	// Can the top paren on the stack be closed with val ?
	// e.g. ( with ), [ with ]

	// If the stack is empty, no
	if s.size == 0 {
		return false
	}
	if (s.top.value == '(' && val == ')') ||
		(s.top.value == '[' && val == ']') ||
		(s.top.value == '{' && val == '}') {
		return true
	}
	return false
}

// Push a new element onto the stack
func (s *ParenthesesStack) Push(val rune) bool {
	if DEBUG {
		fmt.Println(string(val))
	}
	if isClosingParen(val) {
		if DEBUG {
			fmt.Println("close", s.CanClose(val))
		}
		if s.CanClose(val) {
			_ = s.Pop()
		} else { // Mismatching paren
			return false
		}
	} else {
		s.top = &Element{value: val, next: s.top}
		s.size++
	}
	return true
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *ParenthesesStack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *ParenthesesStack) Size() int {
	return s.size
}

func solve(line string) bool {
	stack := new(ParenthesesStack)

	for _, p := range line {
		ok := stack.Push(p)
		if !ok {
			return false
		}
	}

	// Are these still parentheses on the stack ? Malformed
	return stack.Size() == 0
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
		if solve(line) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
