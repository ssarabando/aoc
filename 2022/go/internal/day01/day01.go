package day01

import (
	"strconv"
)

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
