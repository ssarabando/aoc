package day07

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne("day07_test_input.txt")
	expected := 95437
	if actual != expected {
		t.Fatalf(`Day 7, part 1: expected %d; found %d.`, expected, actual)
	}
}
