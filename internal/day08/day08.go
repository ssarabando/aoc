package day08

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func readTreeHeights(filename string) *[][]int {
	trees := [][]int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var height, row int

	scanner := bufio.NewScanner(file)
	row = 0
	for scanner.Scan() {
		line := scanner.Text()
		numberOfTrees := len(line)
		heights := make([]int, numberOfTrees)
		for treeNumber, char := range line {
			height, err = strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			heights[treeNumber] = height
		}
		trees = append(trees, heights)
		row++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &trees
}

type coord struct {
	row int
	col int
}

type coordinates []coord

func (trees *coordinates) contains(row, col int) bool {
	for _, coord := range *trees {
		if coord.row != row {
			continue
		}
		if coord.col == col {
			return true
		}
	}
	return false
}

func (trees *coordinates) tryAdd(row, col int) bool {
	if trees.contains(row, col) {
		return false
	}
	coords := (*[]coord)(trees)
	*coords = append(*coords, coord{row, col})
	return true
}

func (trees *coordinates) len() int {
	return len(*trees)
}

func PartOne(filename string) int {
	numberOfVisibleTrees := 0

	treeHeights := readTreeHeights(filename)

	numberOfRows := len(*treeHeights)
	numberOfCols := len((*treeHeights)[0])

	// The outside trees are always visible.
	// That includes the top and bottom rows...
	numberOfVisibleTrees += numberOfRows*2 - 2
	// ... and the 1st and last columns.
	numberOfVisibleTrees += numberOfCols*2 - 2

	// Coordinates of the trees that are visible from one of the edges
	coords := coordinates{}

	var maxHeight, treeHeight int

	// Excluding edges, find trees visible from left and right
	for row := 1; row < numberOfRows-1; row++ {
		// Start with the left edge height
		maxHeight = (*treeHeights)[row][0]
		// Go right the current row and stop before the right edge
		for col := 1; col < numberOfCols-1; col++ {
			treeHeight = (*treeHeights)[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				coords.tryAdd(row, col)
			}
		}
		// Now start with the right edge height
		maxHeight = (*treeHeights)[row][numberOfCols-1]
		// Go left the current row and stop before the left edge
		for col := numberOfCols - 2; col > 0; col-- {
			treeHeight = (*treeHeights)[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				coords.tryAdd(row, col)
			}
		}
	}
	// Excluding edges, find trees visible from top and bottom
	for col := 1; col < numberOfCols-1; col++ {
		// Start with the top edge height
		maxHeight = (*treeHeights)[0][col]
		// Go down the current column and stop before the bottom edge
		for row := 1; row < numberOfRows-1; row++ {
			treeHeight = (*treeHeights)[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				coords.tryAdd(row, col)
			}
		}
		// Now with the bottom edge height
		maxHeight = (*treeHeights)[numberOfRows-1][col]
		// Go up the current column and stop before the top edge
		for row := numberOfRows - 2; row > 0; row-- {
			treeHeight = (*treeHeights)[row][col]
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				coords.tryAdd(row, col)
			}
		}
	}

	numberOfVisibleTrees += coords.len()

	return numberOfVisibleTrees
}

func PartTwo(filename string) int {
	return 0
}
