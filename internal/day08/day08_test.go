package day08

import "testing"

func TestPartOne(t *testing.T) {
    actual := PartOne("day08_test_input.txt")
    expected := 21
    if actual != expected {
        t.Fatalf(`Day 08, part 1: expected %d; actual %d.`, expected, actual)
    }
}
