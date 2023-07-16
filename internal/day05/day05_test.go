package day05

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne("day05_test_input.txt")
	expected := "CMZ"
	if actual != expected {
		t.Fatalf(`Day 5, Part 1: expected %s but found %s`, expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	actual := PartTwo("day05_test_input.txt")
	expected := "MCD"
	if actual != expected {
		t.Fatalf(`Day 5, Part 2: expected %s but found %s`, expected, actual)
	}
}
