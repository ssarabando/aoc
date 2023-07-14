package main

import (
	"fmt"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day01"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day04/part1"
)

func main() {
	fmt.Println("Day  1, part 2:", day01.PartTwo("day01_input.txt"))
	partOneResult := day04.PartOne("day04_input.txt")
	fmt.Println("Day  4, part 1:", partOneResult)
}