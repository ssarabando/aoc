package day09

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type movementDirection int

const (
	Unknown movementDirection = iota
	Up
	Left
	Right
	Down
)

type movement struct {
	direction movementDirection
	steps     int
}

func sToMovementDirection(d rune) (movementDirection, error) {
	switch d {
	case 'U':
		return Up, nil
	case 'L':
		return Left, nil
	case 'R':
		return Right, nil
	case 'D':
		return Down, nil
	default:
		return Unknown, fmt.Errorf(`invalid movement direction %c`, d)
	}
}

func newMovementFromString(instruction string) movement {
	direction, derr := sToMovementDirection(rune(instruction[0]))
	if derr != nil {
		log.Fatal(derr)
	}

	steps, err := strconv.Atoi(instruction[2:])
	if err != nil {
		log.Fatal(err)
	}

	return movement{direction, steps}
}

type movements []*movement

func (movs *movements) add(mov movement) {
	*movs = append(*movs, &mov)
}

func readMovements(filename string) movements {
	movs := movements{}

	if file, openerr := os.Open(filename); openerr != nil {
		log.Fatal(openerr)
	} else {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			movs.add(newMovementFromString(line))
		}
		if scanerr := scanner.Err(); scanerr != nil {
			log.Fatal(scanerr)
		}
	}

	return movs
}
