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
	for y := 0; y < len(*heightmap); y++ {
		for x := 0; x < len((*heightmap)[y]); x++ {
			isLowPoint := true
			if x == 0 && y == 0 {
				if (*heightmap)[y][x].height >= (*heightmap)[y][x+1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y+1][x].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else if x == len((*heightmap)[y])-1 && y == 0 {
				if (*heightmap)[y][x].height >= (*heightmap)[y][x-1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y+1][x].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else if y == 0 {
				if (*heightmap)[y][x].height >= (*heightmap)[y][x-1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y+1][x].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y][x+1].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else if x == 0 && y == len((*heightmap))-1 {
				if (*heightmap)[y][x].height >= (*heightmap)[y-1][x].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y][x+1].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else if x == len((*heightmap)[y])-1 && y == len((*heightmap))-1 {
				if (*heightmap)[y][x].height >= (*heightmap)[y-1][x].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y][x-1].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else if y == len((*heightmap))-1 {
				if (*heightmap)[y][x].height >= (*heightmap)[y][x-1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y-1][x].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y][x+1].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else if x == 0 {
				if (*heightmap)[y][x].height >= (*heightmap)[y][x+1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y+1][x].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y-1][x].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else if x == len((*heightmap)[y])-1 {
				if (*heightmap)[y][x].height >= (*heightmap)[y][x-1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y+1][x].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y-1][x].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			} else {
				if (*heightmap)[y][x].height >= (*heightmap)[y][x-1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y][x+1].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y+1][x].height ||
					(*heightmap)[y][x].height >= (*heightmap)[y-1][x].height {
					isLowPoint = false
				}

				(*heightmap)[y][x].isLowPoint = isLowPoint
			}
		}
	}
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
