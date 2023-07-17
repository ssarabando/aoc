package day06

import (
	"testing"
)

type partOneTest struct {
	filename string
	expected int
}

func TestPartOne(t *testing.T) {
	tests := []partOneTest{
		{"day06_test_input_1.txt", 7},
		{"day06_test_input_2.txt", 5},
		{"day06_test_input_3.txt", 6},
		{"day06_test_input_4.txt", 10},
		{"day06_test_input_5.txt", 11},
	}
	for _, test := range tests {
		filename := test.filename
		actual := PartOne(filename)
		expected := test.expected
		if actual != expected {
			t.Fatalf(`%s: expected %d; found %d`, filename, expected, actual)
		}
	}
}
