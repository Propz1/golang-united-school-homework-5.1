package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	p := Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
	return p
}

func (s Square) Area() uint {
	return s.a * s.a

}

func (s Square) Perimeter() uint {
	return 2 * (s.a + s.a)
}
