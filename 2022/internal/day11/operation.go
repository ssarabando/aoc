package day11

import (
	"strconv"
	"strings"
)

type operationCalculator interface {
	calculate(old int) int
}

type addOperation struct {
	increment int
}

func (o *addOperation) calculate(old int) int {
	return old + o.increment
}

type multOperation struct {
	factor int
}

func (o *multOperation) calculate(old int) int {
	return old * o.factor
}

type powOperation struct{}

func (o *powOperation) calculate(old int) int {
	return old * old
}

func isOperationLine(line string) bool {
	return len(line) >= len(operationLineKey) && line[:len(operationLineKey)] == operationLineKey
}

func readOperation(line string) operationCalculator {
	terms := strings.Split(line[len(operationLineKey):], " ")
	_, op, rhs := terms[0], terms[1], terms[2]
	switch op {
	case "+":
		value, _ := strconv.Atoi(rhs)
		return &addOperation{value}
	case "*":
		if rhs == "old" {
			return &powOperation{}
		}
		value, _ := strconv.Atoi(rhs)
		return &multOperation{value}
	}

	return nil
}
