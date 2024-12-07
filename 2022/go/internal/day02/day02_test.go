package day02

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne([]string{"A Y", "B X", "C Z"})
	expected := 15
	if actual != expected {
		t.Fatalf("Day 02 part 02: expected %d, got %d.", expected, actual)
	}
}
