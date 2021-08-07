package battleship

// Game represnts a battleship game
type Game struct {
	Players []Player
}

// NewGame returns a game configured according to the number of players.
func NewGame(numberOfPlayers int) Game {
	return Game{
		Players: make([]Player, numberOfPlayers),
	}
}

// SetUpPlayer sets the players boards
func (g Game) SetUpPlayer(playerNumber, gridSize, totalShips, totalMissiles int, shipPositions, missileMoves []Coordinates) {
	if playerNumber < len(g.Players) {
		player := Player{
			TotalMissilies:  totalMissiles,
			TotalShips:      totalShips,
			ShipCoordinates: shipPositions,
			HitCoordinates:  missileMoves,
		}

		// Set up players board
		player.Board = NewBoard(gridSize)
		// Place players
		player.PlaceShips(shipPositions)

		g.Players[playerNumber] = player
	}
}

// Setup phase
// Fire phase
// Player 1 fires at player 2
// Player 2 fires at player 1
