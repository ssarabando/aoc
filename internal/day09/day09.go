package day09

func PartOne(filename string) int {
	movs := readMovements(filename)

	rope := rope{}

	coords := newListOfCoords()
	coords.markVisited(rope.tail)

	for _, mov := range movs {
		step := 0
		switch mov.direction {
		case Up:
			for step < mov.steps {
				rope.moveUp()
				coords.markVisited(rope.tail)
				step++
			}
		case Left:
			for step < mov.steps {
				rope.moveLeft()
				coords.markVisited(rope.tail)
				step++
			}
		case Right:
			for step < mov.steps {
				rope.moveRight()
				coords.markVisited(rope.tail)
				step++
			}
		case Down:
			for step < mov.steps {
				rope.moveDown()
				coords.markVisited(rope.tail)
				step++
			}
		}
	}

	return len(coords.coords)
}

func PartTwo(filename string) int {
	return 0
}
