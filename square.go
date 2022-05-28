package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square Square) End() Point {
	return Point{x: square.start.x + int(square.a), y: square.start.y + int(square.a)}
}

func (sqaure Square) Area() uint {
	return sqaure.a * sqaure.a
}

func (square Square) Perimeter() uint {
	return square.a * 4
}
