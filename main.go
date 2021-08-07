package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gunjan01/battleship/battleship"
)

const (
	// Player represents a battleship player.
	player int = 0
	// Opponent represents a battleship Opponent.
	opponent int = 1
)

var (
	gridSize        int
	totalShips      int
	p1ShipPositions []battleship.Coordinates
	p2ShipPositions []battleship.Coordinates
	totalMissiles   int
	p1MissileMoves  []battleship.Coordinates
	p2MissileMoves  []battleship.Coordinates
)

func sanitizeInputForCoordinates(input string) []battleship.Coordinates {
	positionsArray := strings.Split(input, ":")

	coordinatesArray := []battleship.Coordinates{}

	for _, positions := range positionsArray {
		values := strings.Split(positions, ",")

		coordinate := battleship.Coordinates{}

		for index, value := range values {
			if pos, err := strconv.Atoi(value); err == nil {
				switch index {
				case 0:
					coordinate.X = pos
				case 1:
					coordinate.Y = pos
				}
			}
		}

		coordinatesArray = append(coordinatesArray, coordinate)
	}

	return coordinatesArray
}

func generateOutputFile(game battleship.Game) {
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
	defer outputFile.Close()

	// Write the player boards to the output file
	scoreBoard := make(map[int]string)
	for index := range game.Players {
		player := game.Players[index]
		_, err := outputFile.WriteString("Player" + strconv.Itoa(index+1) + "\n")

		scoreBoard[index] = "P" + strconv.Itoa(index+1) + ":" + strconv.Itoa(player.TotalPoints)

		for i := range player.Board.Board {
			var str string
			for j := range player.Board.Board[i] {
				if player.Board.Board[i][j] != "" {
					str = str + player.Board.Board[i][j] + " "
				} else {
					str = str + "_" + " "
				}
			}

			// Write the string to file
			_, err = outputFile.WriteString(str + "\n")
			if err != nil {
				log.Fatal("Error while writing to a file: ", err)
			}
		}
		_, err = outputFile.WriteString("\n")
	}

	// Write point to the output file.
	for _, value := range scoreBoard {
		_, err = outputFile.WriteString(value + "\n")
		if err != nil {
			log.Fatal("Error while writing to a file: ", err)
		}
	}

	// Write the final result to the output file.
	var finalResult string = "It is a draw"
	if game.Players[player].TotalPoints > game.Players[opponent].TotalPoints {
		finalResult = "Player 1 wins"
	}

	if game.Players[opponent].TotalPoints > game.Players[player].TotalPoints {
		finalResult = "Player 2 wins"
	}

	_, err = outputFile.WriteString("\n" + finalResult + "\n")
	if err != nil {
		log.Fatal("Error while writing to a file: ", err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	defer file.Close()

	var currentLine int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		switch currentLine {
		case 0:
			{
				gridSize, err = strconv.Atoi(input)
				if err != nil {
					// log the fatal error
					log.Fatal("Error while converting input to int: ", err)
				}

				// Validate the number of missiles.
				if !(gridSize > 0 && gridSize < 10) {
					log.Fatal("Invalid input: The grid size must be greater than 0 and lesser than 10.")
				}
			}

		case 1:
			{
				totalShips, err = strconv.Atoi(input)
				if err != nil {
					log.Fatal("Error while converting input to int: ", err)
				}

				if !(totalShips > 0 && totalShips < ((gridSize*gridSize)/2)) {
					log.Fatal("Invalid input: Total number of ships are invalid.")
				}
			}

		case 2:
			{
				p1ShipPositions = sanitizeInputForCoordinates(input)
				if len(p1ShipPositions) != totalShips {
					log.Fatal("Invalid input: The total ship coordinates for player 1 must be equal to the number of ships.")
				}
			}

		case 3:
			{
				p2ShipPositions = sanitizeInputForCoordinates(input)
				if len(p2ShipPositions) != totalShips {
					log.Fatal("Invalid input: The total ship coordinates for player 2 must be equal to the number of ships.")
				}
			}

		case 4:
			{
				totalMissiles, err = strconv.Atoi(input)
				if err != nil {
					log.Fatal("Error while converting input to int: ", err)
				}
				// Validate the number of missiles.
				if !(totalMissiles > 0 && totalMissiles < 100) {
					log.Fatal("Invalid input: The total missiles must be greater than 0 and lesser than 100.")
				}
			}

		case 5:
			{
				p1MissileMoves = sanitizeInputForCoordinates(input)
				if len(p1MissileMoves) != totalMissiles {
					log.Fatal("Invalid input: The total hit coordinates for player 1 must be equal to the total number of missiles.")
				}
			}

		case 6:
			{
				p2MissileMoves = sanitizeInputForCoordinates(input)
				if len(p2MissileMoves) != totalMissiles {
					log.Fatal("Invalid input: The total hit coordinates for player 2 must be equal to the total number of missiles.")
				}
			}

		}

		currentLine = currentLine + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading file: ", err)
	}

	// Start the game
	var numberOfPlayers int = 2
	game := battleship.NewGame(numberOfPlayers)

	// Setup Player and opponent
	game.SetUpPlayer(player, gridSize, totalShips, totalMissiles, p1ShipPositions, p1MissileMoves)
	game.SetUpPlayer(opponent, gridSize, totalShips, totalMissiles, p2ShipPositions, p2MissileMoves)

	var wg sync.WaitGroup
	wg.Add(len(game.Players))

	go func() {
		defer wg.Done()
		game.FireMissiles(player, opponent)
	}()

	go func() {
		defer wg.Done()
		game.FireMissiles(opponent, player)
	}()

	// Wait for the routines to finish.
	wg.Wait()

	generateOutputFile(game)
}
