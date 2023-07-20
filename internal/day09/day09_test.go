package day09

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne("day09_test_input.txt")
	expected := 13
	if actual != expected {
		t.Fatalf(`Day 09, part 1: expected %d, got %d.`, expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	actual := PartTwo("day09_test_input.txt")
	// TODO: set expected when known (after solving part 1)
	expected := -1
	if actual != expected {
		t.Fatalf(`Day 09, part 2: expected %d, got %d.`, expected, actual)
	}
}
