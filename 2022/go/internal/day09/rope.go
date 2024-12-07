package day09

type rope struct {
	knots []coord
}

func newRope(size int) *rope {
	if size < 2 {
		size = 2
	}

	return &rope{
		make([]coord, size),
	}
}

func (r *rope) size() int {
	return len(r.knots)
}

func (r *rope) getHead() *coord {
	return &r.knots[0]
}

func (r *rope) getTail() *coord {
	return &r.knots[r.size()-1]
}

func (r *rope) moveUp() {
	prev := r.getHead()
	prev.moveUp()
	for pos := 1; pos < r.size(); pos++ {
		curr := &r.knots[pos]
		if !curr.isAdjacent(prev) {
			curr.follow(prev)
		}
		prev = curr
	}
}

func (r *rope) moveLeft() {
	prev := r.getHead()
	prev.moveLeft()
	for pos := 1; pos < r.size(); pos++ {
		curr := &r.knots[pos]
		if !curr.isAdjacent(prev) {
			curr.follow(prev)
		}
		prev = curr
	}
}

func (r *rope) moveRight() {
	prev := r.getHead()
	prev.moveRight()
	for pos := 1; pos < r.size(); pos++ {
		curr := &r.knots[pos]
		if !curr.isAdjacent(prev) {
			curr.follow(prev)
		}
		prev = curr
	}
}

func (r *rope) moveDown() {
	prev := r.getHead()
	prev.moveDown()
	for pos := 1; pos < r.size(); pos++ {
		curr := &r.knots[pos]
		if !curr.isAdjacent(prev) {
			curr.follow(prev)
		}
		prev = curr
	}
}
