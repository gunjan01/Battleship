package battleship

import (
	"testing"
)

func TestNewCoordinates(t *testing.T) {
	tests := map[string]struct {
		x int
		y int
	}{
		"test set 1": {
			x: 5,
			y: 10,
		},
		"test set 2": {
			x: 15,
			y: -10,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			coordinates := NewCoordinates(test.x, test.y)

			if test.x != coordinates.X {
				t.Errorf("Received %v, expected %v", coordinates.X, test.x)
			}

			if test.y != coordinates.Y {
				t.Errorf("Received %v, expected %v", coordinates.Y, test.y)
			}
		})
	}
}
