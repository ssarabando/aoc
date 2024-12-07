package day09

import "testing"

func getSimpleTestData() []string {
	return []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
}

func getComplexTestData() []string {
	return []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
}

func TestPartOne(t *testing.T) {
	actual := PartOne(getSimpleTestData(), "day09_test_input.txt")
	expected := 13
	if actual != expected {
		t.Fatalf(`Day 09, part 1: expected %d, got %d.`, expected, actual)
	}
}

func TestPartTwoSimple(t *testing.T) {
	actual := PartTwo(getSimpleTestData())
	expected := 1
	if actual != expected {
		t.Fatalf(`Day 09, part 2 (simple dataset): expected %d, got %d.`, expected, actual)
	}
}

func TestPartTwoComplex(t *testing.T) {
	actual := PartTwo(getComplexTestData())
	expected := 36
	if actual != expected {
		t.Fatalf(`Day 09, part 2 (complex dataset): expected %d, got %d.`, expected, actual)
	}
}
