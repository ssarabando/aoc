package day10

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func computeSignalStrength(cycle, x int) int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		signalStrength := cycle * x
		return signalStrength
	}
	return 0
}

func PartOne(filename string) int {
	file, openErr := os.Open(filename)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer file.Close()

	cycle := 0
	x := 1
	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "noop" {
			cycle++
			result += computeSignalStrength(cycle, x)
		} else if line[:4] == "addx" {
			if value, err := strconv.Atoi(line[5:]); err != nil {
				log.Fatal(err)
			} else {
				cycle++
				result += computeSignalStrength(cycle, x)

				cycle++
				result += computeSignalStrength(cycle, x)

				x += value
			}
		}
	}
	if scanError := scanner.Err(); scanError != nil {
		log.Fatal(scanError)
	}

	return result
}

func PartTwo(filename string) string {
	program := *readProgram(filename)

	screen := makeScreen(40, 6)

	cycle := 0
	x := 1

	for {
		pc := program.getInstruction(cycle)

		col := cycle % screen.width
		row := (cycle - col) / screen.width

		if x-1 <= col && col <= x+1 {
			screen.set(row, col)
		} else {
			screen.unset(row, col)
		}

		switch op := pc.(type) {
		case addx:
			x += addx(op).value
		}

		cycle++

		if cycle >= len(program.instructions) {
			break
		}
	}

	return screen.String()
}
