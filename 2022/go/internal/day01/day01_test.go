package day01

import "testing"

func getTestData() []string {
	return []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
}

func TestPartOne(t *testing.T) {
	actual := PartOne(getTestData())
	expected := 24000
	if actual != expected {
		t.Fatalf("Day 01, part 1: expected %d, got %d.", expected, actual)
	}
}
