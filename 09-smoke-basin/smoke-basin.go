package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, heightData := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	heightmap := initializeHeightmap(&heightData)
	analyzeHeightmap(&heightmap)

	lowPointRisk := calculateLowPointRisk(&heightmap)
	fmt.Println("Risk:", lowPointRisk)
}

func calculateLowPointRisk(heightmap *[][]Location) (risk int) {
	for y := 0; y < len(*heightmap); y++ {
		for x := 0; x < len((*heightmap)[y]); x++ {
			location := (*heightmap)[y][x]
			if location.isLowPoint {
				risk += location.height + 1
			}
		}
	}

	return
}

func analyzeHeightmap(heightmap *[][]Location) {
	directions := 4
	rowNumbers := [4]int{-1, 0, 0, 1}
	columnNumbers := [4]int{0, 1, -1, 0}

	for row := 0; row < len(*heightmap); row++ {
		for col := 0; col < len((*heightmap)[row]); col++ {
			isLowPoint := true

			for i := 0; i < directions; i++ {
				testRow := row + rowNumbers[i]
				testCol := col + columnNumbers[i]

				if isSafe(heightmap, testRow, testCol) {
					location := (*heightmap)[row][col]
					adjacentLocation := (*heightmap)[testRow][testCol]
					isLowPoint = location.height < adjacentLocation.height

					if !isLowPoint {
						break
					}
				}
			}

			(*heightmap)[row][col].isLowPoint = isLowPoint
		}
	}
}

func isSafe(heightmap *[][]Location, row int, column int) bool {
	totalRows := len(*heightmap)
	totalColumns := len((*heightmap)[0])
	return row >= 0 && column >= 0 && row < totalRows && column < totalColumns
}

func initializeHeightmap(heightData *[]string) (heightmap [][]Location) {
	for _, heightDataRow := range *heightData {
		var locations []Location

		heights := strings.Split(heightDataRow, "")
		for _, height := range heights {
			height, err := strconv.Atoi(height)
			if err != nil {
				fmt.Println("Cannot parse int from string:", height)
			}
			locations = append(locations, Location{height: height})
		}

		heightmap = append(heightmap, locations)
	}

	return
}

type Location struct {
	height     int
	isLowPoint bool
	visited    bool
}
