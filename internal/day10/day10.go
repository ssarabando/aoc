package day10

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

var targetCycles []int = []int{
	20,
	60,
	100,
	140,
	180,
	220,
}

var cycle int = 0
var x int = 1

func computeSignalStrength() int {
	if slices.Contains(targetCycles, cycle) {
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

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "noop" {
			cycle++
			result += computeSignalStrength()
		} else if line[:4] == "addx" {
			if value, err := strconv.Atoi(line[5:]); err != nil {
				log.Fatal(err)
			} else {
				cycle++
				result += computeSignalStrength()

				cycle++
				result += computeSignalStrength()

				x += value
			}
		}
	}
	if scanError := scanner.Err(); scanError != nil {
		log.Fatal(scanError)
	}

	return result
}
