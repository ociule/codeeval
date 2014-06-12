package main

import "fmt"
import "testing"

func Test_subMatrix(t *testing.T) {
	tests := []struct {
		in             [][]string
		oy, ox, ey, ex int
		expected       [][]string
	}{
		{[][]string{{"1", "2", "3", "4", "5", "6", "7"}, {"1", "2", "3", "4", "5", "6", "7"}, {"1", "2", "3", "4", "5", "6", "7"}},
			1, 0, 1, 7,
			[][]string{{"1", "2", "3", "4", "5", "6", "7"}}},
		{[][]string{{"1", "2", "3", "4", "5", "6", "7"}, {"1", "2", "3", "4", "5", "6", "7"}, {"1", "2", "3", "4", "5", "6", "7"}},
			1, 1, 1, 5,
			[][]string{{"2", "3", "4", "5", "6"}}},
	}
	for _, test := range tests {
		actual := subMatrix(test.in, test.oy, test.ox, test.ey, test.ex)
		if fmt.Sprintf("%d", actual) != fmt.Sprintf("%d", test.expected) {
			t.Errorf("subMatrix(%s, %d, %d, %d, %d): expected %s, got %s", test.in, test.oy, test.ox, test.ey, test.ex, test.expected, actual)
		}
	}
}

func Test_parseMatrix(t *testing.T) {
	tests := []struct {
		n, m     int
		raw      string
		expected [][]string
	}{
		{2, 1, "1 2", [][]string{{"1"}, {"2"}}},
		{2, 2, "1 2 3 4", [][]string{{"1", "2"}, {"3", "4"}}},
	}

	for _, test := range tests {
		actual := parseMatrix(test.n, test.m, test.raw)
		if fmt.Sprintf("%s", actual) != fmt.Sprintf("%s", test.expected) {
			t.Errorf("parseMatrix(%d, %d, '%s'): expected %s, got %s", test.n, test.m, test.raw, test.expected, actual)
		}
	}
}
