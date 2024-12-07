package day05

import "testing"

func getTestData() []string {
	return []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
}

func TestPartOne(t *testing.T) {
	actual := PartOne(getTestData())
	expected := "CMZ"
	if actual != expected {
		t.Fatalf(`Day 5, Part 1: expected %s but found %s`, expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	actual := PartTwo(getTestData())
	expected := "MCD"
	if actual != expected {
		t.Fatalf(`Day 5, Part 2: expected %s but found %s`, expected, actual)
	}
}
