package day03

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const priorites string = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func line_priority(line string) int {
	number_of_items := len(line)

	mid_point := number_of_items / 2
	for pos := 0; pos < mid_point; pos++ {
		item := string(line[pos])
		if strings.Contains(line[mid_point:number_of_items], item) {
			return strings.Index(priorites, item)
		}
	}

	return 0
}

func group_priorities(elf_1_items string, elf_2_items string, elf_3_items string) int {
	common_items := map[rune]struct{}{}
	for _, item := range elf_1_items {
		if strings.ContainsRune(elf_2_items, item) {
			common_items[item] = struct{}{}
		}
	}
	for k := range common_items {
		if strings.ContainsRune(elf_3_items, k) {
			return strings.Index(priorites, string(k))
		}
	}
	return 0
}

func readData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	results := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}
		// Grow results by one
		l := len(results)
		newResults := make([]string, l+1)
		copy(newResults, results)
		results = newResults
		// Append line to results
		results[l] = scanner.Text()
	}

	return results
}

func PartOne(filename string) int {
	result := 0

	lines := readData(filename)
	for _, line := range lines {
		result += line_priority(line)
	}

	return result
}

func PartTwo(filename string) int {
	result := 0

	lines := readData(filename)
	number_of_lines := len(lines)
	for line_number := 2; line_number < number_of_lines; line_number += 3 {
		result += group_priorities(lines[line_number-2], lines[line_number-1], lines[line_number])
	}

	return result
}
