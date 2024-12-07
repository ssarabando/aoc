package day08

import (
	"log"
	"strconv"
)

type treeHeights struct {
	treeHeights []int
	width       int
	height      int
}

func newTreeHeights(size int) *treeHeights {
	// The forest is always a square
	return &treeHeights{
		make([]int, size*size),
		size,
		size,
	}
}

func (f *treeHeights) getHeight(row, col int) int {
	return f.treeHeights[row*f.width+col]
}

func (f *treeHeights) setHeight(row, col, height int) {
	f.treeHeights[row*f.width+col] = height
}

func readTreeHeights(lines []string) *treeHeights {
	var forest *treeHeights

	row := 0
	for _, line := range lines {
		if row == 0 {
			// A forest is always a square in this case
			forest = newTreeHeights(len(line))
		}
		for treeNumber, char := range line {
			if height, err := strconv.Atoi(string(char)); err != nil {
				log.Fatal(err)
			} else {
				forest.setHeight(row, treeNumber, height)
			}
		}
		row++
	}

	return forest
}
