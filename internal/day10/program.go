package day10

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type noop struct{}

type addx struct {
	value int
}

func makeAddx(value string) addx {
	result := addx{}
	if v, err := strconv.Atoi(value); err != nil {
		log.Fatal(err)
	} else {
		result.value = v
	}
	return result
}

type program struct {
	instructions []interface{}
}

func makeProgram() *program {
	return &program{
		[]interface{}{},
	}
}

func (p *program) addNoop() {
	p.instructions = append(p.instructions, noop{})
}

func (p *program) addAddX(value string) {
	// Add a noop to advance the cycle counter append the the real addx so that
	// it is easier to get the timing right between cycles and ops.
	p.instructions = append(p.instructions, noop{}, makeAddx(value))
}

func (p *program) getInstruction(cycle int) interface{} {
	if cycle < len(p.instructions) {
		return p.instructions[cycle]
	}
	return noop{}
}

func readProgram(filename string) *program {
	file, openErr := os.Open(filename)
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer file.Close()

	program := makeProgram()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		op := scanner.Text()

		if len(op) >= 4 && op[:4] == "noop" {
			program.addNoop()
			continue
		}

		if len(op) > 5 && op[:4] == "addx" {
			program.addAddX(op[5:])
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return program
}
