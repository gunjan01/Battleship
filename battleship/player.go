package battleship

// Player represents a player playing a battleship round.
type Player struct {
	Board           BattleShipBoard
	TotalMissilies  int
	TotalShips      int
	ShipCoordinates []Coordinates
	HitCoordinates  []Coordinates
	TotalPoints     int
}

// PlaceShips places ships on a players board at the given coordinates.
func (p *Player) PlaceShips(shipCoordinates []Coordinates) {
	for _, coordinate := range shipCoordinates {
		p.Board.PlaceShip(coordinate.X, coordinate.Y)
	}
}

// FireShots fires missiles at the opponents boards at the given coordinates.
// If a ship is present at the opponent players at the given coordinates, we mark it a hit.
// Else we mark it a MISS.
// On HIT, we advance the players points for each successful shot.
func (p *Player) FireShots(opponentBoard BattleShipBoard) {
	for _, coordinate := range p.HitCoordinates {
		if opponentBoard.Board[coordinate.X][coordinate.Y] == "B" {
			opponentBoard.MarkHit(coordinate.X, coordinate.Y)

			// Increase present players points.
			p.TotalPoints = p.TotalPoints + 1
		} else {
			opponentBoard.MarkMiss(coordinate.X, coordinate.Y)
		}
	}
}
