package day08

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
