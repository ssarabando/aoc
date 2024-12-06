package day08

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
