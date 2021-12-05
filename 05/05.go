package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, coordinates := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	var verticalLines []string
	var horizontalLines []string
	var diagonalLines []string
	maxX := 0
	maxY := 0
	for _, coordinates := range coordinates {
		xyxy := strings.Fields(coordinates)
		start := xyxy[0]
		end := xyxy[2]

		startFields := strings.Split(start, ",")
		endFields := strings.Split(end, ",")

		startX, _ := strconv.ParseInt(startFields[0], 10, 64)
		startY, _ := strconv.ParseInt(startFields[1], 10, 64)
		endX, _ := strconv.ParseInt(endFields[0], 10, 64)
		endY, _ := strconv.ParseInt(endFields[1], 10, 64)

		if int(startY) > maxY {
			maxY = int(startY)
		}

		if int(startX) > maxX {
			maxX = int(startX)
		}

		if startY == endY {
			horizontalLines = append(horizontalLines, coordinates)
		} else if startX == endX {
			verticalLines = append(verticalLines, coordinates)
		} else {
			diagonalLines = append(diagonalLines, coordinates)
		}
	}

	var gridRow []int
	var max int

	if maxX > maxY {
		max = maxX
	} else {
		max = maxY
	}

	var grid [][]int
	for i := 0; i <= max; i++ {
		gridRow = make([]int, max+1)
		grid = append(grid, gridRow)
	}

	for _, coordinates := range horizontalLines {
		xyxy := strings.Fields(coordinates)
		start := xyxy[0]
		end := xyxy[2]

		startFields := strings.Split(start, ",")
		endFields := strings.Split(end, ",")

		yIndex, _ := strconv.ParseInt(startFields[1], 10, 64)
		xStart, _ := strconv.ParseInt(startFields[0], 10, 64)
		xEnd, _ := strconv.ParseInt(endFields[0], 10, 64)

		if xStart < xEnd {
			for i := xStart; i <= xEnd; i++ {
				grid[yIndex][i]++
			}
		} else {
			for i := xEnd; i <= xStart; i++ {
				grid[yIndex][i]++
			}
		}
	}

	for _, coordinates := range verticalLines {
		xyxy := strings.Fields(coordinates)
		start := xyxy[0]
		end := xyxy[2]

		startFields := strings.Split(start, ",")
		endFields := strings.Split(end, ",")

		xIndex, _ := strconv.ParseInt(startFields[0], 10, 64)
		yStart, _ := strconv.ParseInt(startFields[1], 10, 64)
		yEnd, _ := strconv.ParseInt(endFields[1], 10, 64)

		if yStart < yEnd {
			for i := yStart; i <= yEnd; i++ {
				grid[i][xIndex]++
			}
		} else {
			for i := yEnd; i <= yStart; i++ {
				grid[i][xIndex]++
			}
		}
	}

	for _, coordinates := range diagonalLines {
		xyxy := strings.Fields(coordinates)
		start := xyxy[0]
		end := xyxy[2]

		startFields := strings.Split(start, ",")
		endFields := strings.Split(end, ",")

		xStart, _ := strconv.ParseInt(startFields[0], 10, 64)
		xEnd, _ := strconv.ParseInt(endFields[0], 10, 64)
		yStart, _ := strconv.ParseInt(startFields[1], 10, 64)
		yEnd, _ := strconv.ParseInt(endFields[1], 10, 64)

		points := int(math.Abs(float64(xStart - xEnd)))
		for i := 0; i <= points; i++ {
			if yStart < yEnd {
				if xStart < xEnd {
					grid[int(yStart)+i][int(xStart)+i]++
				} else {
					grid[int(yStart)+i][int(xStart)-i]++
				}
			} else {
				if xStart < xEnd {
					grid[int(yStart)-i][int(xStart)+i]++
				} else {
					grid[int(yStart)-i][int(xStart)-i]++
				}
			}
		}
	}

	intersections := 0
	for _, row := range grid {
		for _, dot := range row {
			if dot >= 2 {
				intersections++
			}
		}
	}

	fmt.Println("At Least 2 Intersections:", intersections)
}
