package day08

import (
	"bufio"
	"log"
	"os"
    "strconv"
)

type treeHeights struct {
	width       int
	height      int
	treeHeights []int
}

func newTreeHeights(size int) *treeHeights {
	// The forest is always a square
	return &treeHeights{
		size,
		size,
		make([]int, size*size),
	}
}

func (f *treeHeights) getHeight(row, col int) int {
	return f.treeHeights[row*f.width+col]
}

func (f *treeHeights) setHeight(row, col, height int) {
	f.treeHeights[row*f.width+col] = height
}

func readTreeHeights(filename string) *treeHeights {
	var forest *treeHeights

	file, openerr := os.Open(filename)
	if openerr != nil {
		log.Fatal(openerr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return forest
}
