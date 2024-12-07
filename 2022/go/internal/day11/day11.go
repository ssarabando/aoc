package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type round struct{}

type note struct {
	monkey      int
	bag         bag
	operation   operationCalculator
	divisibleBy int
	yayMonkey   int
	nayMonkey   int
}

const (
	monkeyNumberLineKey  string = "Monkey "
	startingItemsLineKey string = "  Starting items: "
	operationLineKey     string = "  Operation: new = "
	testLineKey          string = "  Test: divisible by "
	yayLineKey           string = "    If true: throw to monkey "
	nayLineKey           string = "    If false: throw to monkey "
)

func isLine(line string, key string) bool {
	return len(line) >= len(key) && line[:len(key)] == key
}

func readMonkeyNumber(line string) int {
	var monkeyNumber int
	var err error
	if monkeyNumber, err = strconv.Atoi(line[len(monkeyNumberLineKey) : len(line)-1]); err != nil {
		log.Fatal(err)
	}

	return monkeyNumber
}

func readTest(line string) int {
	result, _ := strconv.Atoi(line[len(testLineKey):])
	return result
}

func readYayMonkey(line string) int {
	result, _ := strconv.Atoi(line[len(yayLineKey):])
	return result
}

func readNayMonkey(line string) int {
	result, _ := strconv.Atoi(line[len(yayLineKey):])
	return result
}

func readNotes(filename string) []note {
	notes := []note{}

	if file, err := os.Open(filename); err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(file)

		currentNote := note{}

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				// end of note
				notes = append(notes, currentNote)
				currentNote = note{}
				continue
			}

			switch {
			case isLine(line, monkeyNumberLineKey):
				currentNote.monkey = readMonkeyNumber(line)
			case isLine(line, startingItemsLineKey):
				currentNote.bag = readStartingItems(line)
			case isLine(line, operationLineKey):
				currentNote.operation = readOperation(line)
			case isLine(line, testLineKey):
				currentNote.divisibleBy = readTest(line)
			case isLine(line, yayLineKey):
				currentNote.yayMonkey = readYayMonkey(line)
			case isLine(line, nayLineKey):
				currentNote.nayMonkey = readNayMonkey(line)
			}
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return notes
}

func PartOne(filename string) int {
	notes := readNotes(filename)
	worryLevel := -1
	fmt.Printf("%v %d", notes, worryLevel)
	return 0
}
