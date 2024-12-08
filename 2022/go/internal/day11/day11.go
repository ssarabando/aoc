package day11

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type monkey struct {
	Term1               string
	Operation           string
	Term2               string
	StartingItems       []uint64
	Number              int
	DivisibleBy         int
	MonkeyNumberIfTrue  int
	MonkeyNumberIfFalse int
	InspectionNumber    uint64
}

func newMonkey(notes []string) (*monkey, error) {
	m := monkey{}

	if number, err := extractMonkeyNumber(notes[0]); err != nil {
		return nil, err
	} else {
		m.Number = number
	}

	items := strings.Split(strings.TrimPrefix(notes[1], "  Starting items: "), ", ")
	m.StartingItems = make([]uint64, len(items))
	for i, item := range items {
		if itemNumber, err := strconv.ParseUint(item, 10, 0); err != nil {
			return nil, err
		} else {
			m.StartingItems[i] = itemNumber
		}
	}

	// Extracts the right hand side of the operation as a slice.
	// Example: '  Operation: new = old * 19' yields []string{"old", "*", "19"}.
	operationTerms := strings.Fields(strings.TrimPrefix(notes[2], "  Operation: new = "))
	m.Term1 = operationTerms[0]
	m.Operation = operationTerms[1]
	m.Term2 = operationTerms[2]

	// Extracts the rightmost number, which is the divisor.
	// Example: '  Test: divisible by 23' yields []string{"Test:", "divisible", "by", "23"}. 23 is the divisor.
	testFields := strings.Fields(notes[3])
	if divisor, err := strconv.Atoi(testFields[len(testFields)-1]); err != nil {
		return nil, err
	} else {
		m.DivisibleBy = divisor
	}

	// As above, but this time extract the number of the monkey if the test is true.
	ifTrueFields := strings.Fields(notes[4])
	if monkeyNumber, err := strconv.Atoi(ifTrueFields[len(ifTrueFields)-1]); err != nil {
		return nil, err
	} else {
		m.MonkeyNumberIfTrue = monkeyNumber
	}

	// As above, but this time extract the number of the monkey if the test is false.
	ifFalseFields := strings.Fields(notes[5])
	if monkeyNumber, err := strconv.Atoi(ifFalseFields[len(ifFalseFields)-1]); err != nil {
		return nil, err
	} else {
		m.MonkeyNumberIfFalse = monkeyNumber
	}

	m.InspectionNumber = 0

	return &m, nil
}

func computeMonkeyBusiness(lines []string, rounds int, partOne bool) (uint64, error) {
	monkeys := []monkey{}
	notes := [6]string{}
	for lineNumber, line := range lines {
		if line == "" {
			if m, err := newMonkey(notes[:]); err != nil {
				return 0, err
			} else {
				monkeys = append(monkeys, *m)
			}
			notes = [6]string{}
			continue
		}

		notes[lineNumber%7] = line
	}

	var common_multiple uint64 = 1
	for _, monkey := range monkeys {
		common_multiple *= uint64(monkey.DivisibleBy)
	}

	for round := 1; round <= rounds; round++ {
		for monkeyNumber := 0; monkeyNumber < len(monkeys); monkeyNumber++ {
			m := &monkeys[monkeyNumber]
			m.InspectionNumber += uint64(len(m.StartingItems))
			for _, worryLevel := range m.StartingItems {
				var term1 uint64 = 0
				if m.Term1 == "old" {
					term1 = worryLevel
				} else {
					if term, err := strconv.ParseUint(m.Term1, 10, 0); err != nil {
						return 0, err
					} else {
						term1 = term
					}
				}
				var term2 uint64 = 0
				if m.Term2 == "old" {
					term2 = worryLevel
				} else {
					if term, err := strconv.ParseUint(m.Term2, 10, 0); err != nil {
						return 0, err
					} else {
						term2 = term
					}
				}
				switch m.Operation {
				case "+":
					worryLevel = term1 + term2
				case "*":
					worryLevel = term1 * term2
				default:
					return 0, fmt.Errorf("Invalid operation %s in monkey %d.", m.Operation, m.Number)
				}

				if partOne {
					worryLevel /= 3
				} else {
					// Without this, it would overflow, event with an uint64.
					// Modulo over a common multiple between all monkeys. Since each divisor is a prime,
					// multiplying all of them yields a number that we can use do decrease the worry levels.
					// Do not fully understand it, but for more information, see
					// (this)[https://www.reddit.com/r/adventofcode/comments/zifqmh/comment/izupmlj/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button]
					worryLevel %= common_multiple
				}

				var targetMonkey *monkey
				if worryLevel%uint64(m.DivisibleBy) == 0 {
					targetMonkey = &monkeys[m.MonkeyNumberIfTrue]
				} else {
					targetMonkey = &monkeys[m.MonkeyNumberIfFalse]
				}

				targetMonkey.StartingItems = append(targetMonkey.StartingItems, worryLevel)
			}
			m.StartingItems = []uint64{}
		}
	}

	slices.SortFunc(monkeys, func(monkeyA, monkeyB monkey) int {
		// Sort descending
		result := int64(monkeyB.InspectionNumber - monkeyA.InspectionNumber)
		if result < 0 {
			return -1
		} else if result > 0 {
			return 1
		} else {
			return 0
		}
	})

	return monkeys[0].InspectionNumber * monkeys[1].InspectionNumber, nil
}

func extractMonkeyNumber(line string) (int, error) {
	trimmedLine := strings.TrimSuffix(line, ":")
	numberRepr := strings.Fields(trimmedLine)[1]
	return strconv.Atoi(numberRepr)
}

func PartOne(lines []string) (uint64, error) {
	return computeMonkeyBusiness(lines, 20, true)
}

func PartTwo(lines []string) (uint64, error) {
	return computeMonkeyBusiness(lines, 10000, false)
}
