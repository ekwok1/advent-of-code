package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, initialEnergyGrid := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	octopusGrid := initialOctopusEnergyGrid(&initialEnergyGrid)
	flashes := simulateDays(100, &octopusGrid)
	fmt.Println("Flashes after 100 steps:", flashes)

	octopusGrid = initialOctopusEnergyGrid(&initialEnergyGrid)
	simulateDays(1000, &octopusGrid)
}

func simulateDays(days int, octopusGrid *[][]Octopus) (flashes int) {
	rows := len(*octopusGrid)
	cols := len((*octopusGrid)[0])

	var queue []Location

Loop:
	for day := 0; day < days; day++ {
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				newEnergyLevel := (*octopusGrid)[row][col].energy + 1

				if newEnergyLevel > 9 {
					queue = append(queue, Location{row: row, col: col})
				}

				(*octopusGrid)[row][col].energy = newEnergyLevel
			}
		}

		rowDirections := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
		colDirections := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}

		for i := 0; i < len(queue); i++ {
			for dir := 0; dir < 8; dir++ {
				newRow := queue[i].row + rowDirections[dir]
				newCol := queue[i].col + colDirections[dir]

				if isSafe(octopusGrid, newRow, newCol) {
					newEnergyLevel := (*octopusGrid)[newRow][newCol].energy + 1

					if newEnergyLevel > 9 {
						queue = append(queue, Location{row: newRow, col: newCol})
					}

					(*octopusGrid)[newRow][newCol].energy = newEnergyLevel
				}
			}
		}

		if len(queue) == 100 {
			fmt.Println("On day", day+1, "all octopus flashed!")
			break Loop
		}
		flashes += len(queue)

		for _, flashed := range queue {
			(*octopusGrid)[flashed.row][flashed.col].energy = 0
		}

		queue = nil
	}

	return
}

func isSafe(octopusGrid *[][]Octopus, row int, column int) bool {
	totalRows := len(*octopusGrid)
	totalColumns := len((*octopusGrid)[0])

	return row >= 0 &&
		column >= 0 &&
		row < totalRows &&
		column < totalColumns &&
		(*octopusGrid)[row][column].energy < 10
}

func initialOctopusEnergyGrid(initialEnergyGrid *[]string) (grid [][]Octopus) {
	for _, row := range *initialEnergyGrid {
		initialEnergyRow := strings.Split(row, "")

		var ocotopusRow []Octopus
		for _, initialEnergy := range initialEnergyRow {
			energy, _ := strconv.Atoi(initialEnergy)
			ocotopusRow = append(ocotopusRow, Octopus{energy: energy, flashed: make(map[int]bool)})
		}

		grid = append(grid, ocotopusRow)
	}

	return
}

type Location struct {
	row int
	col int
}

type Octopus struct {
	energy  int
	flashed map[int]bool
}
