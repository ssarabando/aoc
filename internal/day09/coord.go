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
//
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

func (target *coord) alignHorizontallyWith(source *coord) {
	target.y = source.y
}

func (target *coord) alignVerticallyWith(source *coord) {
	target.x = source.x
}

func (source *coord) isAdjacent(dest *coord) bool {
	vertDistance := abs(source.y - dest.y)
	horzDistance := abs(source.x - dest.x)
	return vertDistance <= 1 && horzDistance <= 1
}

type listOfCoords struct {
	coords map[coord]bool
}

func newListOfCoords() *listOfCoords {
	return &listOfCoords{
		make(map[coord]bool),
	}
}

func (l *listOfCoords) markVisited(coord coord) {
	if _, found := l.coords[coord]; !found {
		l.coords[coord] = true
	}
}
