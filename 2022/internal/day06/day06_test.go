package day06

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
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

func TestPartTwo(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"day06_test_input_6.txt", 19},
		{"day06_test_input_7.txt", 23},
		{"day06_test_input_8.txt", 23},
		{"day06_test_input_9.txt", 29},
		{"day06_test_input_10.txt", 26},
	}
	for _, test := range tests {
		filename := test.filename
		actual := PartTwo(filename)
		expected := test.expected
		if actual != expected {
			t.Fatalf(`%s: expected %d; found %d`, filename, expected, actual)
		}
	}
}
