package battleship

import "testing"

func TestNewBoard(t *testing.T) {
	testSize := 5
	board := NewBoard(testSize)

	if len(board.Board) != testSize {
		t.Errorf("Received %v, expected %v", len(board.Board), testSize)
	}

	for i := range board.Board {
		if len(board.Board[i]) != testSize {
			t.Errorf("Received %v, expected %v", len(board.Board[i]), testSize)
		}
	}
}

func TestPlaceShip(t *testing.T) {
	tests := map[string]struct {
		x int
		y int
	}{
		"test set 1": {
			x: 2,
			y: 0,
		},
		"test set 2": {
			x: 4,
			y: 3,
		},
		"test set 3": {
			x: 1,
			y: 3,
		},
		"test set 4": {
			x: 0,
			y: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testSize := 5
			board := NewBoard(testSize)

			board.PlaceShip(test.x, test.y)

			if board.Board[test.x][test.y] != MARK_SHIP {
				t.Errorf("Test Failed: Expected a ship to be placed at the corrdinates: (%d, %d)", test.x, test.y)
			}
		})
	}
}

func TestMarkHit(t *testing.T) {
	tests := map[string]struct {
		x int
		y int
	}{
		"test set 1": {
			x: 2,
			y: 0,
		},
		"test set 2": {
			x: 4,
			y: 3,
		},
		"test set 3": {
			x: 1,
			y: 3,
		},
		"test set 4": {
			x: 0,
			y: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testSize := 5
			board := NewBoard(testSize)

			board.MarkHit(test.x, test.y)

			if board.Board[test.x][test.y] != MARK_HIT {
				t.Errorf("Test Failed: Expected a hit to be placed at the corrdinates: (%d, %d)", test.x, test.y)
			}
		})
	}
}

func TestMarkMiss(t *testing.T) {
	tests := map[string]struct {
		x int
		y int
	}{
		"test set 1": {
			x: 2,
			y: 0,
		},
		"test set 2": {
			x: 4,
			y: 3,
		},
		"test set 3": {
			x: 1,
			y: 3,
		},
		"test set 4": {
			x: 0,
			y: 0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testSize := 5
			board := NewBoard(testSize)

			board.MarkMiss(test.x, test.y)

			if board.Board[test.x][test.y] != MARK_MISS {
				t.Errorf("Test Failed: Expected a hit to be placed at the corrdinates: (%d, %d)", test.x, test.y)
			}
		})
	}
}
