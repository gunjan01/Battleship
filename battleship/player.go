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
		p.Board.PlaceShip(coordinate.x, coordinate.y)
	}
}

// FireShots fires missiles at the opponents boards at the given coordinates.
// If a ship is present at the opponent players at the given coordinates, we mark it a hit.
// Else we mark it a miss.
// On HIT, we advance the players points for each successful shot.
func (p *Player) FireShots(opponentBoard BattleShipBoard) {
	for _, coordinate := range p.HitCoordinates {
		if opponentBoard.Board[coordinate.x][coordinate.y] == "B" {
			opponentBoard.MarkHit(coordinate.x, coordinate.y)

			// Increase present players points.
			p.TotalPoints = p.TotalPoints + 1
		} else {
			opponentBoard.MarkMiss(coordinate.x, coordinate.y)
		}
	}
}

// Player
//Board
//Coordinates

// Game
// -> Setup phase
// -> Missle phase
// -> File writing phase

// Player
// Board
// Total Hits
// Total Misses
// Points
// Ttoal Sthots
// Ship Coordinates
// Hit Coordinates
