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
	StartingItems       []int
	Number              int
	DivisibleBy         int
	MonkeyNumberIfTrue  int
	MonkeyNumberIfFalse int
	InspectionNumber    int
}

func newMonkey(notes []string) (*monkey, error) {
	m := monkey{}

	if number, err := extractMonkeyNumber(notes[0]); err != nil {
		return nil, err
	} else {
		m.Number = number
	}

	items := strings.Split(strings.TrimPrefix(notes[1], "  Starting items: "), ", ")
	m.StartingItems = make([]int, len(items))
	for i, item := range items {
		if itemNumber, err := strconv.Atoi(item); err != nil {
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

func PartOne(lines []string, filename string) (int, error) {
	monkeys := []monkey{}
	notes := [6]string{}
	for lineNumber, line := range lines {
		if line == "" {
			if m, err := newMonkey(notes[:]); err != nil {
				return -1, err
			} else {
				monkeys = append(monkeys, *m)
			}
			notes = [6]string{}
			continue
		}

		notes[lineNumber%7] = line
	}

	for round := 1; round < 21; round++ {
		for monkeyNumber := 0; monkeyNumber < len(monkeys); monkeyNumber++ {
			m := &monkeys[monkeyNumber]
			// for _, m := range monkeys {
			m.InspectionNumber += len(m.StartingItems)
			for _, worryLevel := range m.StartingItems {
				term1 := 0
				if m.Term1 == "old" {
					term1 = worryLevel
				} else {
					if term, err := strconv.Atoi(m.Term1); err != nil {
						return -1, err
					} else {
						term1 = term
					}
				}
				term2 := 0
				if m.Term2 == "old" {
					term2 = worryLevel
				} else {
					if term, err := strconv.Atoi(m.Term2); err != nil {
						return -1, err
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
					return -1, fmt.Errorf("Invalid operation %s in monkey %d.", m.Operation, m.Number)
				}

				worryLevel /= 3

				var targetMonkey *monkey
				if worryLevel%m.DivisibleBy == 0 {
					targetMonkey = &monkeys[m.MonkeyNumberIfTrue]
				} else {
					targetMonkey = &monkeys[m.MonkeyNumberIfFalse]
				}

				targetMonkey.StartingItems = append(targetMonkey.StartingItems, worryLevel)
			}
			m.StartingItems = []int{}
		}
	}

	slices.SortFunc(monkeys, func(monkeyA, monkeyB monkey) int {
		// Sort descending
		return monkeyB.InspectionNumber - monkeyA.InspectionNumber
	})

	return monkeys[0].InspectionNumber * monkeys[1].InspectionNumber, nil
}

func extractMonkeyNumber(line string) (int, error) {
	trimmedLine := strings.TrimSuffix(line, ":")
	numberRepr := strings.Fields(trimmedLine)[1]
	return strconv.Atoi(numberRepr)
}
