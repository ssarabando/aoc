package day10

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

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

type screen struct {
	data   []bool
	height int
	width  int
}

func makeScreen(width, height int) *screen {
	return &screen{
		make([]bool, width*height),
		height,
		width,
	}
}

func (s *screen) set(row, col int) {
	if row >= 0 && row < s.height && col >= 0 && col < s.width {
		s.data[s.width*row+col] = true
	}
}

func (s *screen) unset(row, col int) {
	if row >= 0 && row < s.height && col >= 0 && col < s.width {
		s.data[s.width*row+col] = false
	}
}

func (s *screen) String() string {
	m := map[bool]rune{
		true:  '#',
		false: '.',
	}

	result := strings.Builder{}

	for i := 0; i < len(s.data); i++ {
		if i > 0 && i%s.width == 0 {
			result.WriteRune('\n')
		}
		result.WriteRune(m[s.data[i]])
	}
	return result.String()
}

func PartTwo(filename string) string {
	program := *readProgram(filename)

	screen := makeScreen(40, 6)

	x = 1
	cycle = 0
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
