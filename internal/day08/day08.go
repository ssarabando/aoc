package day08

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type scenicScore struct {
	top    int
	left   int
	right  int
	bottom int
}

func (score *scenicScore) compute() int {
	return score.top * score.left * score.right * score.bottom
}

type forest [][]int

func (f *forest) countOfRows() int {
	return len(*f)
}

func (f *forest) countOfCols() int {
	if f.countOfRows() > 0 {
		return len((*f)[0])
	}
	return 0
}

func (f *forest) scenicScore(row, col int) int {
	treeHeight := (*f)[row][col]
	rows := len(*f) - 1
	cols := len((*f)[row]) - 1

	score := scenicScore{}

	pos := 0

	var height int

	// Compute top scenic score
	if row > 0 {
		pos = row - 1
		for pos >= 0 {
			score.top++
			height = (*f)[pos][col]
			if height >= treeHeight {
				// We count a tree of equal height but mustn't count further
				break
			}
			pos--
		}
	}
	// Compute left scenic score
	if col > 0 {
		pos = col - 1
		for pos >= 0 {
			score.left++
			height = (*f)[row][pos]
			if height >= treeHeight {
				// We count a tree of equal height but mustn't count further
				break
			}
			pos--
		}
	}
	// Compute right scenic score
	if col < cols {
		pos = col + 1
		for pos <= cols {
			score.right++
			height = (*f)[row][pos]
			if height >= treeHeight {
				// We count a tree of equal height but mustn't count further
				break
			}
			pos++
		}
	}
	// Compute bottom scenic score
	if row < rows {
		pos = row + 1
		for pos <= rows {
			score.bottom++
			height = (*f)[pos][col]
			if height >= treeHeight {
				// We count a tree of equal height but mustn't count further
				break
			}
			pos++
		}
	}

	return score.compute()
}

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
	treeHeights := readTreeHeights(filename)
	f := (forest)(*treeHeights)

	maxScenicScore := 0

	// Only check the inner trees; the ones on the edges will
	// always have a scenic score of 0.
	row := 1
	rows, cols := f.countOfRows()-1, f.countOfCols()-1
	for row < rows {
		col := 1
		for col < cols {
			treeScenicScore := f.scenicScore(row, col)
			if treeScenicScore > maxScenicScore {
				maxScenicScore = treeScenicScore
			}
			col++
		}
		row++
	}

	return maxScenicScore
}
