package battleship

type Coordinates struct {
	x int
	y int
}

func (c *Coordinates) New(x, y int) Coordinates {
	return Coordinates{
		x: x,
		y: y,
	}
}
