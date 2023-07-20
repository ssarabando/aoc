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

type scenicScores struct {
	width           int
	height          int
	scenicScores    []scenicScore
	bestScenicScore int
}

func (scores *scenicScores) setScenicScore(row, col int, score scenicScore) {
	scores.scenicScores[row*scores.width+col] = score
	currScenicScore := score.compute()
	if scores.bestScenicScore < currScenicScore {
		scores.bestScenicScore = currScenicScore
	}
}

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

type visibleTrees struct {
	width                int
	height               int
	data                 []int
	numberOfVisibleTrees int
}

func NewVisibleTrees(width, height int) *visibleTrees {
	// The edges are always visible (width times two because top row and bottom
	// row plus height times two because left column and right column minus
	// the four corners -- we don't want to count them twice: once in the rows
	// and another in the cols).
	numberOfAlwaysVisibleTrees := width*2 + height*2 - 4

	return &visibleTrees{
		width,
		height,
		make([]int, width*height),
		numberOfAlwaysVisibleTrees,
	}
}

func (source *treeHeights) newVisibleTreesFromTreeHeights() *visibleTrees {
	dest := NewVisibleTrees(source.width, source.height)

	var row, col, maxHeight, currHeight int

	// Start at (1, 1), because the edges are always visible
	row = 1
	for row < dest.height-1 {
		// Compute the heights from the left edge to the right:
		maxHeight = source.getHeight(row, 0)
		col = 1
		for col < dest.width-1 {
			currHeight = source.getHeight(row, col)
			if currHeight > maxHeight {
				maxHeight = currHeight
				dest.markVisited(row, col)
			}
			col++
		}
		// Compute the heights from the right edge to the left
		col = dest.width - 2
		maxHeight = source.getHeight(row, dest.width-1)
		for col > 0 {
			currHeight = source.getHeight(row, col)
			if currHeight > maxHeight {
				maxHeight = currHeight
				dest.markVisited(row, col)
			}
			col--
		}
		row++
	}
	col = 1
	for col < dest.width-1 {
		// Compute the heights from the top edge to the bottom
		row = 1
		maxHeight = source.getHeight(0, col)
		for row < dest.height-1 {
			currHeight = source.getHeight(row, col)
			if currHeight > maxHeight {
				maxHeight = currHeight
				dest.markVisited(row, col)
			}
			row++
		}
		// Compute the heights from the bottom edge to the top
		row = dest.height - 2
		maxHeight = source.getHeight(dest.height-1, col)
		for row > 0 {
			currHeight = source.getHeight(row, col)
			if currHeight > maxHeight {
				maxHeight = currHeight
				dest.markVisited(row, col)
			}
			row--
		}
		col++
	}

	return dest
}

func (v *visibleTrees) markVisited(row, col int) {
	pos := row*v.width + col
	if v.data[pos] != 1 {
		v.data[pos] = 1
		v.numberOfVisibleTrees++
	}
}

func (source *treeHeights) newScenicScoresFromTreeHeights() *scenicScores {
	dest := scenicScores{
		source.width,
		source.height,
		make([]scenicScore, source.width*source.height),
		0,
	}

	dest.setScenicScore(0, 0, scenicScore{})
	dest.setScenicScore(0, source.width-1, scenicScore{})
	dest.setScenicScore(source.height-1, 0, scenicScore{})
	dest.setScenicScore(source.height-1, source.width-1, scenicScore{})

	// Compute the scenic score of each tree.
	// Notes that all the edges have a scenic score of 0.
	for row := 0; row < dest.height-1; row++ {
		for col := 0; col < dest.width-1; col++ {
			isAtEdge := row == 0 || col == 0 || row == dest.height-1 || col == dest.width-1
			if isAtEdge {
				continue
			}

			currTreeHeight := source.getHeight(row, col)
			score := scenicScore{}
			// Compute scenic score to the top
			for pos := row - 1; pos >= 0; pos-- {
				score.top++
				destTreeHeight := source.getHeight(pos, col)
				if destTreeHeight >= currTreeHeight {
					// We count a tree of equal height but not further
					break
				}
			}
			// Compute scenic score to the left
			for pos := col - 1; pos >= 0; pos-- {
				score.left++
				destTreeHeight := source.getHeight(row, pos)
				if destTreeHeight >= currTreeHeight {
					// We count a tree of equal height but not further
					break
				}
			}
			// Compute scenic score to the right
			for pos := col + 1; pos <= dest.width-1; pos++ {
				score.right++
				destTreeHeight := source.getHeight(row, pos)
				if destTreeHeight >= currTreeHeight {
					// We count a tree of equal height but not further
					break
				}
			}
			// Compute scenic score to the bottom
			for pos := row + 1; pos <= dest.height-1; pos++ {
				score.bottom++
				destTreeHeight := source.getHeight(pos, col)
				if destTreeHeight >= currTreeHeight {
					// We count a tree of equal height but not further
					break
				}
			}
			dest.setScenicScore(row, col, score)
		}
	}

	return &dest
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

func PartOne(filename string) int {
	heightMap := readTreeHeights(filename)

	visibilityMap := heightMap.newVisibleTreesFromTreeHeights()

	return visibilityMap.numberOfVisibleTrees
}

func PartTwo(filename string) int {
	heightMap := readTreeHeights(filename)

	scenicScoreMap := heightMap.newScenicScoresFromTreeHeights()

	return scenicScoreMap.bestScenicScore
}
