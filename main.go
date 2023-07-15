package main

import (
	"fmt"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day01"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day02"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day03"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day04/part1"
)

func main() {
	fmt.Println("Day  1, part 2:", day01.PartTwo("day01_input.txt"))
	fmt.Println("Day  2, part 2:", day02.PartTwo("day02_input.txt"))
	fmt.Println("Day  3, part 1:", day03.PartOne("day03_input.txt"))
	fmt.Println("Day  3, part 2:", day03.PartTwo("day03_input.txt"))
	partOneResult := day04.PartOne("day04_input.txt")
	fmt.Println("Day  4, part 1:", partOneResult)
}
