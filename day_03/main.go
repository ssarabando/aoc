package main

import (
	"bufio"
	"fmt"
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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part_1_priorities := 0
	part_2_priorities := 0

	scanner := bufio.NewScanner(file)
	// We can read 3 lines at a time since Part 2 guarantees that there always
	// are 3 elves in each group (that is, 3 lines).
	for {
		success := scanner.Scan()
		if !success {
			break
		}
		line1 := scanner.Text()

		_ = scanner.Scan()
		line2 := scanner.Text()

		_ = scanner.Scan()
		line3 := scanner.Text()

		part_1_priorities += line_priority(line1)
		part_1_priorities += line_priority(line2)
		part_1_priorities += line_priority(line3)

		part_2_priorities += group_priorities(line1, line2, line3)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of elf priorites:", part_1_priorities)
	fmt.Println("Sum of group priorites:", part_2_priorities)
}
