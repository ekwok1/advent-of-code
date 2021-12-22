package main

import (
	"fmt"
	"os"
	"sort"
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

	product := largestBasinProduct(&heightmap, 3)
	fmt.Println("Largest Basin Product:", product)
}

func largestBasinProduct(heightmap *[][]Location, howMany int) int {
	basinSizes := getBasinSizes(heightmap)
	length := len(basinSizes)
	largestSizes := (basinSizes)[length-howMany:]

	product := 1
	for _, size := range largestSizes {
		product *= size
	}

	return product
}

func getBasinSizes(heightmap *[][]Location) (sizes []int) {
	rows := len((*heightmap))
	cols := len((*heightmap)[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			location := (*heightmap)[row][col]
			if location.height != 9 && !location.visited {
				size := 1
				basinDFS(heightmap, row, col, &size)

				sizes = append(sizes, size)
			}
		}
	}

	sort.Ints(sizes)
	return
}

func basinDFS(heightmap *[][]Location, row, col int, size *int) {
	rows := [4]int{-1, 0, 0, 1}
	cols := [4]int{0, -1, 1, 0}

	(*heightmap)[row][col].visited = true

	for i := 0; i < 4; i++ {
		if isDFSSafe(heightmap, row+rows[i], col+cols[i]) {
			(*size)++
			basinDFS(heightmap, row+rows[i], col+cols[i], size)
		}
	}
}

func isDFSSafe(heightmap *[][]Location, row, col int) bool {
	return isSafe(heightmap, row, col) &&
		(*heightmap)[row][col].height != 9 &&
		!(*heightmap)[row][col].visited
}

func calculateLowPointRisk(heightmap *[][]Location) (risk int) {
	totalRows := len(*heightmap)
	totalColumns := len((*heightmap)[0])

	for row := 0; row < totalRows; row++ {
		for col := 0; col < totalColumns; col++ {
			location := (*heightmap)[row][col]
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

func isSafe(heightmap *[][]Location, row, column int) bool {
	totalRows := len(*heightmap)
	totalColumns := len((*heightmap)[0])
	return row >= 0 && column >= 0 && row < totalRows && column < totalColumns
}

func initializeHeightmap(heightData *[]string) (heightmap [][]Location) {
	for _, heightDataRow := range *heightData {
		var locations []Location

		heights := strings.Split(heightDataRow, "")
		for _, height := range heights {
			height, _ := strconv.Atoi(height)
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
