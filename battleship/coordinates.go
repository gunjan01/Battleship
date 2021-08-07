package battleship

// Coordinates represent a pair of coordinates on a grid.
type Coordinates struct {
	X int
	Y int
}

// NewCoordinates returns a newly configured pair of coordinates
func NewCoordinates(x, y int) Coordinates {
	return Coordinates{
		X: x,
		Y: y,
	}
}
