package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

func getContents(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\r\n")
}

func main() {
	{
		lines := getContents("../day01_input.txt")
		fmt.Println("Day  1, part 1:", day01.PartOne(lines))
		fmt.Println("Day  1, part 2:", day01.PartTwo(lines))
	}
	{
		lines := getContents("../day02_input.txt")
		fmt.Println("Day  2, part 1:", day02.PartOne(lines))
		fmt.Println("Day  2, part 2:", day02.PartTwo(lines))
	}
	{
		lines := getContents("../day03_input.txt")
		fmt.Println("Day  3, part 1:", day03.PartOne(lines))
		fmt.Println("Day  3, part 2:", day03.PartTwo(lines))
	}
	{
		lines := getContents("../day04_input.txt")
		fmt.Println("Day  4, part 1:", day04.PartOne(lines))
		fmt.Println("Day  4, part 2:", day04.PartTwo(lines))
	}
	{
		lines := getContents("../day05_input.txt")
		fmt.Println("Day  5, part 1:", day05.PartOne(lines))
		fmt.Println("Day  5, part 2:", day05.PartTwo(lines))
	}
	{
		lines := getContents("../day06_input.txt")
		fmt.Println("Day  6, part 1:", day06.PartOne(lines))
		fmt.Println("Day  6, part 2:", day06.PartTwo(lines))
	}
	{
		lines := getContents("../day07_input.txt")
		fmt.Println("Day  7, part 1:", day07.PartOne(lines))
		fmt.Println("Day  7, part 2:", day07.PartTwo(lines))
	}
	{
		lines := getContents("../day08_input.txt")
		fmt.Println("Day  8, part 1:", day08.PartOne(lines))
		fmt.Println("Day  8, part 2:", day08.PartTwo(lines))
	}
	{
		lines := getContents("../day09_input.txt")
		fmt.Println("Day  9, part 1:", day09.PartOne(lines, "../day09_input.txt"))
		fmt.Println("Day  9, part 2:", day09.PartTwo(lines))
	}
	{
		lines := getContents("../day10_input.txt")
		fmt.Println("Day 10, part 1:", day10.PartOne(lines))
		fmt.Printf("Day 10, part 2:\n%s\n", day10.PartTwo(lines))
	}
	{
		lines := getContents("../day11_input.txt")
		if result, err := day11.PartOne(lines); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Day 11, part 1:", result)
		}
		if result, err := day11.PartTwo(lines); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Day 11, part 2:", result)
		}
	}
}
