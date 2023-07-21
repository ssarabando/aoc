package day09

func PartOne(filename string) int {
	movs := readMovements(filename)

	rope := newRope(2)

	coords := newListOfCoords()
	coords.markVisited(rope.getTail())

	for _, mov := range movs {
		step := 0
		switch mov.direction {
		case Up:
			for step < mov.steps {
				rope.moveUp()
				coords.markVisited(rope.getTail())
				step++
			}
		case Left:
			for step < mov.steps {
				rope.moveLeft()
				coords.markVisited(rope.getTail())
				step++
			}
		case Right:
			for step < mov.steps {
				rope.moveRight()
				coords.markVisited(rope.getTail())
				step++
			}
		case Down:
			for step < mov.steps {
				rope.moveDown()
				coords.markVisited(rope.getTail())
				step++
			}
		}
	}

	return len(coords.coords)
}

func PartTwo(filename string) int {
	movs := readMovements(filename)

	rope := newRope(10)

	coords := newListOfCoords()
	coords.markVisited(rope.getTail())

	for _, mov := range movs {
		step := 0
		switch mov.direction {
		case Up:
			for step < mov.steps {
				rope.moveUp()
				coords.markVisited(rope.getTail())
				step++
			}
		case Left:
			for step < mov.steps {
				rope.moveLeft()
				coords.markVisited(rope.getTail())
				step++
			}
		case Right:
			for step < mov.steps {
				rope.moveRight()
				coords.markVisited(rope.getTail())
				step++
			}
		case Down:
			for step < mov.steps {
				rope.moveDown()
				coords.markVisited(rope.getTail())
				step++
			}
		}
	}

	return len(coords.coords)
}
