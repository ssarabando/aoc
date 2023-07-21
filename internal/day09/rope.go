package day09

type rope struct {
	head coord
	tail coord
}

func (r *rope) isStraight() bool {
	return r.head.x == r.tail.x || r.head.y == r.tail.y
}

func (r *rope) makeHorizontal() {
	r.tail.alignHorizontallyWith(&r.head)
}

func (r *rope) makeVertical() {
	r.tail.alignVerticallyWith(&r.head)
}

func (r *rope) moveUp() {
	r.head.moveUp()
	if !r.head.isAdjacent(&r.tail) {
		r.tail.moveUp()
		if !r.isStraight() {
			r.makeVertical()
		}
	}
}

func (r *rope) moveLeft() {
	r.head.moveLeft()
	if !r.head.isAdjacent(&r.tail) {
		r.tail.moveLeft()
		if !r.isStraight() {
			r.makeHorizontal()
		}
	}
}

func (r *rope) moveRight() {
	r.head.moveRight()
	if !r.head.isAdjacent(&r.tail) {
		r.tail.moveRight()
		if !r.isStraight() {
			r.makeHorizontal()
		}
	}
}

func (r *rope) moveDown() {
	r.head.moveDown()
	if !r.head.isAdjacent(&r.tail) {
		r.tail.moveDown()
		if !r.isStraight() {
			r.makeVertical()
		}
	}
}
