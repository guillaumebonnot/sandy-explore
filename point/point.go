package point

type Point struct {
	x, y int16
}

func (p *Point) Path() uint32 {
	// need to convert to uint16, otherwise go add some 1
	x := uint16(p.x)
	y := uint16(p.y)
	return (uint32(x) << 16) + uint32(y)
}

func New(x, y int16) *Point {
	return &Point{x: x, y: y}
}

// return true if there is a mine at (x,y)
func (p *Point) IsMine() bool {
	return sum_digits(p.x)+sum_digits(p.y) > 21
}

func abs(x int16) int16 {
	if x < 0 {
		return -x
	}
	return x
}

func sum_digits(x int16) int16 {
	sum := int16(0)
	x = abs(x)
	for x > int16(0) {
		sum += x % 10
		x = x / 10
	}
	return sum
}

func (p *Point) Neighbours() []*Point {
	one := int16(1)
	return []*Point{
		&Point{p.x + one, p.y},
		&Point{p.x - one, p.y},
		&Point{p.x, p.y + one},
		&Point{p.x, p.y - one},
	}
}
