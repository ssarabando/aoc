package day08

import "testing"

func getTestData() []string {
	return []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
}

func TestPartOne(t *testing.T) {
	actual := PartOne(getTestData())
	expected := 21
	if actual != expected {
		t.Fatalf(`Day 08, part 1: expected %d; actual %d.`, expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	actual := PartTwo(getTestData())
	expected := 8
	if actual != expected {
		t.Fatalf(`Day 08, part 2: expected %d; actual %d.`, expected, actual)
	}
}
