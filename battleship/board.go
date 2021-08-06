package battleship

const (
	MARK_SHIP = "B"
	MARK_HIT  = "X"
	MARK_MISS = "O"
)

// BattleShipBoard represents the player's board.
type BattleShipBoard struct {
	Board [][]string
}

// NewBoard Returns a new board.
func (b *BattleShipBoard) NewBoard(size int) BattleShipBoard {
	return BattleShipBoard{
		Board: make([][]string, size),
	}
}

// PlaceShip places a ship on the board at the given coordinates.
func (b *BattleShipBoard) PlaceShip(x, y int) {
	b.Board[x][y] = MARK_SHIP
}

// MarkHit marks a hit on the board at the given coordinates.
func (b *BattleShipBoard) MarkHit(x, y int) {
	b.Board[x][y] = MARK_HIT
}

// MarkMiss marks a miss on the board at the given coordinates.
func (b *BattleShipBoard) MarkMiss(x, y int) {
	b.Board[x][y] = MARK_MISS
}
