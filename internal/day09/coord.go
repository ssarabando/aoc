package day09

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// y
// ^
// | (-1,-1) (-1,0) (-1,1)
// | ( 0,-1) ( 0,0) ( 0,1)
// | ( 1,-1) ( 1,0) ( 1,1)
// +----------------------> x
//
// Up decrements y.
// Left decrements x.
// Right increments x.
// Down increments y.
type coord struct {
	x, y int
}

func (c *coord) moveUp() {
	c.y -= 1
}

func (c *coord) moveLeft() {
	c.x -= 1
}

func (c *coord) moveRight() {
	c.x += 1
}

func (c *coord) moveDown() {
	c.y += 1
}

func (source *coord) isDirectlyUp(target *coord) bool {
	return source.y < target.y && source.x == target.x
}

func (source *coord) isDirectlyDown(target *coord) bool {
	return source.y > target.y && source.x == target.x
}

func (source *coord) isDirectlyLeft(target *coord) bool {
	return source.x < target.x && source.y == target.y
}

func (source *coord) isDirectlyRight(target *coord) bool {
	return source.x > target.x && source.y == target.y
}

func (source *coord) isUp(target *coord) bool {
	return source.y < target.y
}

func (source *coord) isLeft(target *coord) bool {
	return source.x < target.x
}

func (source *coord) isRight(target *coord) bool {
	return source.x > target.x
}

func (source *coord) isDown(target *coord) bool {
	return source.y > target.y
}

func (source *coord) isAdjacent(dest *coord) bool {
	vertDistance := abs(source.y - dest.y)
	horzDistance := abs(source.x - dest.x)
	return vertDistance <= 1 && horzDistance <= 1
}

func (tail *coord) follow(head *coord) {
	if head.isDirectlyUp(tail) {
		tail.moveUp()
	} else if head.isDirectlyLeft(tail) {
		tail.moveLeft()
	} else if head.isDirectlyRight(tail) {
		tail.moveRight()
	} else if head.isDirectlyDown(tail) {
		tail.moveDown()
	} else if head.isUp(tail) && head.isLeft(tail) {
		tail.moveUp()
		tail.moveLeft()
	} else if head.isLeft(tail) && head.isDown(tail) {
		tail.moveLeft()
		tail.moveDown()
	} else if head.isDown(tail) && head.isRight(tail) {
		tail.moveDown()
		tail.moveRight()
	} else if head.isRight(tail) && head.isUp(tail) {
		tail.moveRight()
		tail.moveUp()
	}
}

type listOfCoords struct {
	coords map[coord]bool
}

func newListOfCoords() *listOfCoords {
	return &listOfCoords{
		make(map[coord]bool),
	}
}

func (l *listOfCoords) markVisited(coord *coord) {
	if _, found := l.coords[*coord]; !found {
		l.coords[*coord] = true
	}
}
