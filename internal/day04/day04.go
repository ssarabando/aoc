package day04

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type section struct {
	Start  int
	Finish int
}

// contains returns true if the lhs section fully contains the rhs section.
func (lhs *section) Contains(rhs *section) bool {
	return lhs.Start <= rhs.Start && lhs.Finish >= rhs.Finish
}

// contains returns true if the lhs section overlaps the rhs section.
func (lhs *section) Overlaps(rhs *section) bool {
	return lhs.Start <= rhs.Finish && lhs.Finish >= rhs.Start
}

func NewSection(data string) *section {
	items := strings.Split(data, "-")

	start, err := strconv.Atoi(items[0])
	if err != nil {
		log.Fatal(err)
	}

	finish, err := strconv.Atoi(items[1])
	if err != nil {
		log.Fatal(err)
	}

	return &section{start, finish}
}

type pair struct {
	Elf1 *section
	Elf2 *section
}

func NewPair(data string) *pair {
	ranges := strings.Split(data, ",")

	elf1 := NewSection(ranges[0])
	elf2 := NewSection(ranges[1])

	return &pair{elf1, elf2}
}

func appendPair(pairs []pair, newPair *pair) []pair {
	m := len(pairs)

	newSliceOfPairs := make([]pair, m+1)
	copy(newSliceOfPairs, pairs)
	pairs = newSliceOfPairs

	pairs[m] = *newPair

	return pairs
}

func readInputFile(filename string) []pair {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pairs := make([]pair, 0)

	scanner := bufio.NewScanner(file)
	for {
		if success := scanner.Scan(); !success {
			break
		}
		pair := NewPair(scanner.Text())
		pairs = appendPair(pairs, pair)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pairs
}

func PartOne(filename string) int {
	overlapping_pairs := 0
	pairs := readInputFile(filename)
	for _, pair := range pairs {
		if pair.Elf1.Contains(pair.Elf2) || pair.Elf2.Contains(pair.Elf1) {
			overlapping_pairs += 1
		}
	}
	return overlapping_pairs
}

func PartTwo(filename string) int {
	overlapping_pairs := 0
	pairs := readInputFile(filename)
	for _, pair := range pairs {
		if pair.Elf1.Overlaps(pair.Elf2) || pair.Elf2.Overlaps(pair.Elf1) {
			overlapping_pairs += 1
		}
	}
	return overlapping_pairs
}
