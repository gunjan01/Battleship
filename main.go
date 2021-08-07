package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gunjan01/battleship/battleship"
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
			}

		case 1:
			{
				totalShips, err = strconv.Atoi(input)
				if err != nil {
					log.Fatal("Error while converting input to int: ", err)
				}
			}

		case 2:
			{
				p1ShipPositions = sanitizeInputForCoordinates(input)
			}

		case 3:
			{
				p2ShipPositions = sanitizeInputForCoordinates(input)
			}

		case 4:
			{
				totalMissiles, err = strconv.Atoi(input)
				if err != nil {
					log.Fatal("Error while converting input to int: ", err)
				}
			}

		case 5:
			{
				p1MissileMoves = sanitizeInputForCoordinates(input)
			}

		case 6:
			{
				p2MissileMoves = sanitizeInputForCoordinates(input)
			}

		}

		currentLine = currentLine + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading file: ", err)
	}

	fmt.Println(gridSize)
	fmt.Println(totalMissiles)
	fmt.Println(totalShips)
	fmt.Println(p1ShipPositions)
	fmt.Println(p2ShipPositions)
	fmt.Println(p1MissileMoves)
	fmt.Println(p2MissileMoves)
}
