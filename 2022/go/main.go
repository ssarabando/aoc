package main

import (
	"fmt"

	"github.com/ssarabando/aoc/2022/go/internal/day01"
	"github.com/ssarabando/aoc/2022/go/internal/day02"
	"github.com/ssarabando/aoc/2022/go/internal/day03"
	"github.com/ssarabando/aoc/2022/go/internal/day04"
	"github.com/ssarabando/aoc/2022/go/internal/day05"
	"github.com/ssarabando/aoc/2022/go/internal/day06"
	"github.com/ssarabando/aoc/2022/go/internal/day07"
	"github.com/ssarabando/aoc/2022/go/internal/day08"
	"github.com/ssarabando/aoc/2022/go/internal/day09"
	"github.com/ssarabando/aoc/2022/go/internal/day10"
	"github.com/ssarabando/aoc/2022/go/internal/day11"
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
	fmt.Println("Day  8, part 2:", day08.PartTwo("day08_input.txt"))
	fmt.Println("Day  9, part 1:", day09.PartOne("day09_input.txt"))
	fmt.Println("Day  9, part 2:", day09.PartTwo("day09_input.txt"))
	fmt.Println("Day 10, part 1:", day10.PartOne("day10_input.txt"))
	fmt.Printf("Day 10, part 2:\n%s\n", day10.PartTwo("day10_input.txt"))
	fmt.Println("Day 11, part 1:", day11.PartOne("day11_input.txt"))
}
