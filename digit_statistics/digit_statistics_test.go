package main

import "testing"

func Test_get_last_digit_of_pow(t *testing.T) {
	var tests = []struct {
		a      int
		n      int
		output int
	}{
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 8},
		{2, 4, 6},
		{2, 5, 2},

		{3, 1, 3},
		{3, 2, 9},
		{3, 3, 7},
		{3, 4, 1},
		{3, 5, 3},

		{9, 1, 9},
		{9, 2, 1},
		{9, 3, 9},
		{9, 4, 1},
		{9, 5, 9},
	}
	for _, tt := range tests {
		actual := get_last_digit_of_pow(tt.a, tt.n)
		if actual != tt.output {
			t.Errorf("get_last_digit_of_pow(%d, %d): expected %d, got %d", tt.a, tt.n, tt.output, actual)
		}

	}
}
