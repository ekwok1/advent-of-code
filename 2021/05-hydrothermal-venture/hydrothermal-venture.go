package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ekwok1/aoc-2021/utilities"
	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, coordinates := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	gridSize, horizontal, vertical, diagonal := getGridSizeAndFilterLineTypes(&coordinates)
	grid := utilities.CreateIntGrid(gridSize + 1)

	drawHorizontal(&horizontal, &grid)
	drawVertical(&vertical, &grid)
	drawDiagonal(&diagonal, &grid)

	intersections := countIntersections(&grid)
	fmt.Println("At Least 2 Intersections:", intersections)
}

func countIntersections(grid *[][]int) (intersections int) {
	for _, row := range *grid {
		for _, space := range row {
			if space >= 2 {
				intersections++
			}
		}
	}

	return
}

func drawHorizontal(horizontal *[]string, grid *[][]int) {
	for _, coordinates := range *horizontal {
		xStart, yIndex, xEnd, _ := getCoordinates(coordinates)

		if xStart < xEnd {
			for x := xStart; x <= xEnd; x++ {
				(*grid)[yIndex][x]++
			}
		} else {
			for x := xStart; x >= xEnd; x-- {
				(*grid)[yIndex][x]++
			}
		}
	}
}

func drawVertical(vertical *[]string, grid *[][]int) {
	for _, coordinates := range *vertical {
		_, yStart, xIndex, yEnd := getCoordinates(coordinates)

		if yStart < yEnd {
			for y := yStart; y <= yEnd; y++ {
				(*grid)[y][xIndex]++
			}
		} else {
			for y := yStart; y >= yEnd; y-- {
				(*grid)[y][xIndex]++
			}
		}
	}
}

func drawDiagonal(diagonal *[]string, grid *[][]int) {
	for _, coordinates := range *diagonal {
		xStart, yStart, xEnd, yEnd := getCoordinates(coordinates)

		points := int(math.Abs(float64(xStart - xEnd)))
		for i := 0; i <= points; i++ {
			if yStart < yEnd {
				if xStart < xEnd {
					(*grid)[yStart+i][xStart+i]++
				} else {
					(*grid)[yStart+i][xStart-i]++
				}
			} else {
				if xStart < xEnd {
					(*grid)[yStart-i][xStart+i]++
				} else {
					(*grid)[yStart-i][xStart-i]++
				}
			}
		}
	}
}

func getGridSizeAndFilterLineTypes(coordinates *[]string) (gridSize int, horizontal, vertical, diagonal []string) {
	for _, coordinates := range *coordinates {
		xStart, yStart, xEnd, yEnd := getCoordinates(coordinates)

		coords := [4]int{xStart, yStart, xEnd, yEnd}
		for _, coord := range coords {
			if coord > gridSize {
				gridSize = coord
			}
		}

		if yStart == yEnd {
			horizontal = append(horizontal, coordinates)
		} else if xStart == xEnd {
			vertical = append(vertical, coordinates)
		} else {
			diagonal = append(diagonal, coordinates)
		}
	}

	return
}

func getCoordinates(coordinates string) (xStart, yStart, xEnd, yEnd int) {
	fmt.Sscanf(coordinates, "%d,%d -> %d,%d", &xStart, &yStart, &xEnd, &yEnd)
	return
}
