package main

import "testing"
import "errors"
import "reflect"

func errCompare(a, b error) bool {
	return a == b || (reflect.TypeOf(a) == reflect.TypeOf(b))
}

func Test_sum16Bits(t *testing.T) {
	var tests = []struct {
		in         []byte
		output     int
		output_err error
	}{
		{[]byte{0, 1, 0, 1}, 2, nil},
		{[]byte{0, 2, 0, 2}, 4, nil},
		{[]byte{1, 0, 1, 0}, 512, nil},
		{[]byte{0, 1, 1, 0, 0, 1}, 258, nil},
		{[]byte{0, 0, 1}, 0, errors.New("len(in) must be even and longer than 4 - must have at least 2 16 bits number to sum!")},
	}

	for _, tt := range tests {
		actual, err := sum16Bits(tt.in)
		if actual != tt.output || !errCompare(err, tt.output_err) {
			t.Errorf("sum16Bits(%s): expected %d %v, got %d %v", tt.in, tt.output, tt.output_err, actual, err)
		}
	}
}
