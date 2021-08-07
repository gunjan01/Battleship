package battleship

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) New(x, y int) Coordinates {
	return Coordinates{
		X: x,
		Y: y,
	}
}
