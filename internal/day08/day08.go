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

type visibleTrees struct {
	width                uint
	height               uint
	data                 []uint
	numberOfVisibleTrees uint
}

func NewVisibleTrees(width, height uint) visibleTrees {
	// The edges are always visible (width times two because top row and bottom
	// row plus height times two because left column and right column minus
	// the four corners -- we don't want to count them twice: once in the rows
	// and another in the cols).
	numberOfAlwaysVisibleTrees := width*2 + height*2 - 4

	return visibleTrees{
		width,
		height,
		make([]uint, width*height),
		numberOfAlwaysVisibleTrees,
	}
}

func (v *visibleTrees) markVisited(row, col uint) {
	pos := row*v.width + col
	if v.data[pos] != 1 {
		v.data[pos] = 1
		v.numberOfVisibleTrees++
	}
}

func PartOne(filename string) int {
	treeHeights := readTreeHeights(filename)

	numberOfRows := uint(len(*treeHeights))
	numberOfCols := uint(len((*treeHeights)[0]))

	// Coordinates of the trees that are visible from one of the edges
	visibleTrees := NewVisibleTrees(numberOfRows, numberOfCols)

	var maxHeight, treeHeight uint

	// Excluding edges, find trees visible from left and right
	var row, col uint
	for row = 1; row < numberOfRows-1; row++ {
		// Start with the left edge height
		maxHeight = uint((*treeHeights)[row][0])
		// Go right the current row and stop before the right edge
		for col = 1; col < numberOfCols-1; col++ {
			treeHeight = uint((*treeHeights)[row][col])
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibleTrees.markVisited(row, col)
			}
		}
		// Now start with the right edge height
		maxHeight = uint((*treeHeights)[row][numberOfCols-1])
		// Go left the current row and stop before the left edge
		for col = numberOfCols - 2; col > 0; col-- {
			treeHeight = uint((*treeHeights)[row][col])
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibleTrees.markVisited(row, col)
			}
		}
	}
	// Excluding edges, find trees visible from top and bottom
	for col = 1; col < numberOfCols-1; col++ {
		// Start with the top edge height
		maxHeight = uint((*treeHeights)[0][col])
		// Go down the current column and stop before the bottom edge
		for row = 1; row < numberOfRows-1; row++ {
			treeHeight = uint((*treeHeights)[row][col])
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibleTrees.markVisited(row, col)
			}
		}
		// Now with the bottom edge height
		maxHeight = uint((*treeHeights)[numberOfRows-1][col])
		// Go up the current column and stop before the top edge
		for row = numberOfRows - 2; row > 0; row-- {
			treeHeight = uint((*treeHeights)[row][col])
			if treeHeight > maxHeight {
				maxHeight = treeHeight
				visibleTrees.markVisited(row, col)
			}
		}
	}

	return int(visibleTrees.numberOfVisibleTrees)
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
