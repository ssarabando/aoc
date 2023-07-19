package main

import (
	"fmt"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day01"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day02"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day03"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day04"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day05"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day06"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day07"
	"github.com/ssarabando/go_advent_of_code_2022/internal/day08"
)

func main() {
	fmt.Println("Day  1, part 2:", day01.PartTwo("day01_input.txt"))
	fmt.Println("Day  2, part 2:", day02.PartTwo("day02_input.txt"))
	fmt.Println("Day  3, part 1:", day03.PartOne("day03_input.txt"))
	fmt.Println("Day  3, part 2:", day03.PartTwo("day03_input.txt"))
	fmt.Println("Day  4, part 1:", day04.PartOne("day04_input.txt"))
	fmt.Println("Day  4, part 2:", day04.PartTwo("day04_input.txt"))
	fmt.Println("Day  5, part 1:", day05.PartOne("day05_input.txt"))
	fmt.Println("Day  5, part 2:", day05.PartTwo("day05_input.txt"))
	fmt.Println("Day  6, part 1:", day06.PartOne("day06_input.txt"))
	fmt.Println("Day  6, part 2:", day06.PartTwo("day06_input.txt"))
	fmt.Println("Day  7, part 1:", day07.PartOne("day07_input.txt"))
	fmt.Println("Day  7, part 2:", day07.PartTwo("day07_input.txt"))
	fmt.Println("Day  8, part 1:", day08.PartOne("day08_input.txt"))
}
