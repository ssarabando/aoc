package day11

import (
	"testing"
)

func getTestData() []string {
	return []string{
		"Monkey 0:",
		"  Starting items: 79, 98",
		"  Operation: new = old * 19",
		"  Test: divisible by 23",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 3",
		"",
		"Monkey 1:",
		"  Starting items: 54, 65, 75, 74",
		"  Operation: new = old + 6",
		"  Test: divisible by 19",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 0",
		"",
		"Monkey 2:",
		"  Starting items: 79, 60, 97",
		"  Operation: new = old * old",
		"  Test: divisible by 13",
		"    If true: throw to monkey 1",
		"    If false: throw to monkey 3",
		"",
		"Monkey 3:",
		"  Starting items: 74",
		"  Operation: new = old + 3",
		"  Test: divisible by 17",
		"    If true: throw to monkey 0",
		"    If false: throw to monkey 1",
		"",
	}
}

func TestPartOne(t *testing.T) {
	var actual uint64
	actual, _ = PartOne(getTestData())
	const expected uint64 = 10605
	if actual != expected {
		t.Fatalf("Day 11, part 1: expected %d, got %d.", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	var actual uint64
	actual, _ = PartTwo(getTestData())
	const expected uint64 = 2713310158
	if actual != expected {
		t.Fatalf("Day 11, part 2: expected %d, got %d.", expected, actual)
	}
}
