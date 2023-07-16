package day05

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne("day05_test_input.txt")
	expected := "CMZ"
	if actual != expected {
		t.Fatalf(`Day 05, Part 01: expected %s but found %s`, expected, actual)
	}
}
