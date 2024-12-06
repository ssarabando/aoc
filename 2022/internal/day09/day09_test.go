package day09

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne("day09_test_input.txt")
	expected := 13
	if actual != expected {
		t.Fatalf(`Day 09, part 1: expected %d, got %d.`, expected, actual)
	}
}

func TestPartTwoSimple(t *testing.T) {
	actual := PartTwo("day09_test_input.txt")
	expected := 1
	if actual != expected {
		t.Fatalf(`Day 09, part 2 (simple dataset): expected %d, got %d.`, expected, actual)
	}
}

func TestPartTwoComplex(t *testing.T) {
	actual := PartTwo("day09p2_test_input.txt")
	expected := 36
	if actual != expected {
		t.Fatalf(`Day 09, part 2 (complex dataset): expected %d, got %d.`, expected, actual)
	}
}
