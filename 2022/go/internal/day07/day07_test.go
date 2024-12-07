package day07

import "testing"

func getTestData() []string {
	return []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}
}

func TestPartOne(t *testing.T) {
	actual := PartOne(getTestData())
	expected := 95437
	if actual != expected {
		t.Fatalf(`Day 7, part 1: expected %d; found %d.`, expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	actual := PartTwo(getTestData())
	expected := 24933642
	if actual != expected {
		t.Fatalf(`Day 7, part 2: expected %d; found %d.`, expected, actual)
	}
}
