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
