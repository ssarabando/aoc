package day01

import (
	"strconv"
)

func PartOne(lines []string) int {
	maxCalories := 0
	currentElfCalories := 0
	for _, line := range lines {
		if line == "" {
			// End of the list of calories carried by the current Elf.
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}
			currentElfCalories = 0
			continue
		}
		calories, _ := strconv.Atoi(line)
		currentElfCalories += calories
	}
	return maxCalories
}

func PartTwo(lines []string) int {
	var top1, top2, top3, current int

	for _, line := range lines {
		if line == "" {
			if current > top1 {
				top3 = top2
				top2 = top1
				top1 = current
			} else if current > top2 {
				top3 = top2
				top2 = current
			} else if current > top3 {
				top3 = current
			}
			current = 0
		} else {
			calories, _ := strconv.Atoi(line)
			current += calories
		}
	}

	return top1 + top2 + top3
}
