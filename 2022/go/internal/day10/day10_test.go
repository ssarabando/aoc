package day10

import "testing"

func TestPartOne(t *testing.T) {
	actual := PartOne("day10_test_input.txt")
	expected := 13140
	if actual != expected {
		t.Fatalf(`Day 10, part 1: expected %d, got %d.`, expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	actual := PartTwo("day10_test_input.txt")
	expected := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	if actual != expected {
		t.Fatalf("Day 10, part 2: expected \n%s\n, got \n%s\n.", expected, actual)
	}
}
