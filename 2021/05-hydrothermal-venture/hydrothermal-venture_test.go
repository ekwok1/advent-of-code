package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCountIntersectionsStraightLinesOnly(t *testing.T) {
	file, coordinates := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	gridSize, horizontal, vertical, _ := getGridSizeAndFilterLineTypes(&coordinates)
	grid := utilities.CreateIntGrid(gridSize + 1)

	drawHorizontal(&horizontal, &grid)
	drawVertical(&vertical, &grid)

	intersections := countIntersections(&grid)

	if intersections != 5 {
		t.Errorf("countIntersections(&grid) = %d; want 5", intersections)
	}
}

func TestCountIntersections(t *testing.T) {
	file, coordinates := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	gridSize, horizontal, vertical, diagonal := getGridSizeAndFilterLineTypes(&coordinates)
	grid := utilities.CreateIntGrid(gridSize + 1)

	drawHorizontal(&horizontal, &grid)
	drawVertical(&vertical, &grid)
	drawDiagonal(&diagonal, &grid)

	intersections := countIntersections(&grid)

	if intersections != 12 {
		t.Errorf("countIntersections(&grid) = %d; want 12", intersections)
	}
}
