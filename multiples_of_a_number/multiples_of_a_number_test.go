package main

import "fmt"
import "testing"

var tests = []struct {
    x, n, expected int
}{
    {8, 2, 4},
    {13, 8, 1},
    {17, 16, 1},
    {15, 16, 0},
    {63, 8, 7},
}


func Test_linearDivision(t *testing.T) {
	for _, test := range tests {
		actual := linearDivision(test.x, test.n)
		if fmt.Sprintf("%d", actual) != fmt.Sprintf("%d", test.expected) {
			t.Errorf("linearDivision(%d, %d): expected %d, got %d", test.x, test.n, test.expected, actual)
		}
	}
}
