package day04

import (
	"testing"
)

func TestSectionCreation(t *testing.T) {
	s := NewSection("2-4")
	if s.Start != 2 {
		t.Fatalf(`NewSection("2-4") should start with 2 but starts with %d.`, s.Start)
	}
	if s.Finish != 4 {
		t.Fatalf(`NewSection("2-4") should finish with 4 but finishes with %d.`, s.Finish)
	}
}

func TestContains(t *testing.T) {
	p1 := NewPair("1-2,3-4")
	if p1.Elf1.Contains(p1.Elf2) {
		t.Fatalf(`NewPair("1-2,3-4") Elf1 should not fully contain Elf2.`)
	}
	if p1.Elf2.Contains(p1.Elf1) {
		t.Fatalf(`%v-%v: 2nd pair should fully contain the 1st pair.`, p1.Elf1, p1.Elf2)
	}

	p2 := NewPair("1-3,1-2")
	if !p2.Elf1.Contains(p2.Elf2) {
		t.Fatalf(`%v-%v: 1st pair should fully contain the 2nd pair.`, p2.Elf1, p2.Elf2)
	}

	p3 := NewPair("1-3,2-3")
	if !p3.Elf1.Contains(p3.Elf2) {
		t.Fatalf(`%v-%v: 1st pair should fully contain the 2nd pair.`, p3.Elf1, p3.Elf2)
	}

	p4 := NewPair("1-3,2-2")
	if !p4.Elf1.Contains(p4.Elf2) {
		t.Fatalf(`%v-%v: 1st pair should fully contain the 2nd pair.`, p4.Elf1, p4.Elf2)
	}

	p5 := NewPair("1-3,1-3")
	if !p5.Elf1.Contains(p5.Elf2) {
		t.Fatalf(`%v-%v: 1st pair should fully contain the 2nd pair.`, p5.Elf1, p5.Elf2)
	}
}

func TestOverlaps(t *testing.T) {
	overlaps := PartTwo([]string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"})
	if overlaps != 4 {
		t.Fatalf(`Part 2: number of overlaps should be 4 but was %d.`, overlaps)
	}
}

func TestPairCreation(t *testing.T) {
	p := NewPair("1-2,3-4")
	if p.Elf1.Start != 1 {
		t.Fatalf(`NewPair("1-2,3-4").Elf1.Start should be 1 but is %d.`, p.Elf1.Start)
	}
	if p.Elf1.Finish != 2 {
		t.Fatalf(`NewPair("1-2,3-4").Elf1.Finish should be 2 but is %d.`, p.Elf1.Finish)
	}
	if p.Elf2.Start != 3 {
		t.Fatalf(`NewPair("1-2,3-4").Elf2.Start should be 3 but is %d.`, p.Elf2.Start)
	}
	if p.Elf2.Finish != 4 {
		t.Fatalf(`NewPair("1-2,3-4").Elf2.Finish should be 4 but is %d.`, p.Elf2.Finish)
	}
}

func TestPartOneWithSmallDataset(t *testing.T) {
	overlapping_pairs := PartOne([]string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8"})
	if overlapping_pairs == 0 {
		t.Fatalf(`Number of test overlapping pairs should be 2. Was %d.`, overlapping_pairs)
	}
}
