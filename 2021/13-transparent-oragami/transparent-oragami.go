package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, transparentPaper := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	foldInstructionsIndex := utilities.IndexOf("fold", &transparentPaper)
	strCoordinates := transparentPaper[:foldInstructionsIndex]
	instructions := transparentPaper[foldInstructionsIndex:]

	uniqueCoordinates := foldAndGetUniqueCoordinates(&strCoordinates, &instructions, 1)
	fmt.Println("Number of unique coordinates:", len(uniqueCoordinates))

	uniqueCoordinatesPart2 := foldAndGetUniqueCoordinates(&strCoordinates, &instructions, len(instructions))
	plot(uniqueCoordinatesPart2)
}

func plot(uniqueCoordinates map[Coordinates]bool) {
	gridSize := getGridSize(&uniqueCoordinates)
	grid := utilities.CreateIntGrid(gridSize + 1)

	for coordinates := range uniqueCoordinates {
		grid[coordinates.y][coordinates.x]++
	}

	for _, row := range grid {
		fmt.Println(row)
	}
}

func foldAndGetUniqueCoordinates(strCoordinates *[]string, instructions *[]string, folds int) map[Coordinates]bool {
	var coordinates []Coordinates
	for _, coordinate := range *strCoordinates {
		x, y := getCoordinates(coordinate)
		coordinates = append(coordinates, Coordinates{x: x, y: y})
	}

	for i, instruction := range *instructions {
		if i >= folds {
			continue
		}

		axis, number := getFoldingInstruction(instruction)
		switch axis {
		case "y":
			for i, coordinate := range coordinates {
				if coordinate.y > number {
					delta := (coordinate.y - number) * 2
					coordinates[i].y -= delta
				}
			}
		case "x":
			for i, coordinate := range coordinates {
				if coordinate.x > number {
					delta := (coordinate.x - number) * 2
					coordinates[i].x -= delta
				}
			}
		}
	}

	uniqueCoordinates := make(map[Coordinates]bool)
	for _, coordinate := range coordinates {
		uniqueCoordinates[Coordinates{x: coordinate.x, y: coordinate.y}] = true
	}

	return uniqueCoordinates
}

func getCoordinates(coordinates string) (x, y int) {
	fmt.Sscanf(coordinates, "%d,%d", &x, &y)
	return
}

func getFoldingInstruction(instruction string) (axis string, coordinate int) {
	foldAlong := strings.Fields(instruction)
	foldAlongInstructions := strings.Split(foldAlong[2], "=")
	axis = foldAlongInstructions[0]
	coordinate, _ = strconv.Atoi(foldAlongInstructions[1])

	return
}

func getGridSize(coordinatesMap *map[Coordinates]bool) (gridSize int) {
	for coordinates := range *coordinatesMap {
		x := coordinates.x
		y := coordinates.y

		if x > gridSize {
			gridSize = x
		}

		if y > gridSize {
			gridSize = y
		}
	}

	return
}

type Coordinates struct {
	x int
	y int
}
