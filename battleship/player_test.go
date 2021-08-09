package battleship

import "testing"

func TestPlaceShips(t *testing.T) {
	tests := map[string]struct {
		player      Player
		coordinates []Coordinates
	}{
		"test set 1": {
			player: Player{
				Board: NewBoard(5),
			},
			coordinates: []Coordinates{
				Coordinates{
					X: 0,
					Y: 0,
				},
			},
		},
		"test set 2": {
			player: Player{
				Board: NewBoard(2),
			},
			coordinates: []Coordinates{
				Coordinates{
					X: 1,
					Y: 1,
				},
				Coordinates{
					X: 0,
					Y: 1,
				},
				Coordinates{
					X: 1,
					Y: 0,
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.player.PlaceShips(test.coordinates)

			for _, coordinates := range test.coordinates {
				if test.player.Board.Board[coordinates.X][coordinates.Y] != MARK_SHIP {
					t.Errorf("Test Failed: Expected to have a ship placed here")
				}
			}
		})
	}
}

func TestFireShots(t *testing.T) {
	tests := map[string]struct {
		playerA Player
		playerB Player
	}{
		"test set 1": {
			playerA: Player{
				Board: NewBoard(5),
				HitCoordinates: []Coordinates{
					Coordinates{
						X: 0,
						Y: 0,
					},
					Coordinates{
						X: 1,
						Y: 1,
					},
				},
				TotalPoints: 0,
			},
			playerB: Player{
				Board: NewBoard(5),
			},
		},
		"test set 2": {
			playerA: Player{
				Board: NewBoard(5),
				HitCoordinates: []Coordinates{
					Coordinates{
						X: 0,
						Y: 1,
					},
					Coordinates{
						X: 3,
						Y: 3,
					},
				},
				TotalPoints: 0,
			},
			playerB: Player{
				Board: NewBoard(5),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.playerA.FireShots(test.playerB.Board)

			for _, coordinates := range test.playerA.HitCoordinates {
				value := test.playerB.Board.Board[coordinates.X][coordinates.Y]
				if value != MARK_HIT && value != MARK_MISS {
					t.Errorf("Test Failed: Expected to have a hit or miss merked here.")
				}
			}
		})
	}
}
