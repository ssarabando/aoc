package day11

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne("day11_test_input.txt")
	expected := 10605
	if actual != expected {
		t.Fatalf("Day 11, part 1: expected %d, got %d.", expected, actual)
	}
}
