package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestGetMaxHeight(t *testing.T) {
	file, targetArea := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	_, _, yMin, _ := getCoordinateRange(targetArea[0])
	maxHeight := getMaxHeight(yMin)
	if maxHeight != 45 {
		t.Errorf("getCoordinateRange(targetArea[0]) = %d; want 45", maxHeight)
	}
}

func TestCountHits(t *testing.T) {
	file, targetArea := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	xMin, xMax, yMin, yMax := getCoordinateRange(targetArea[0])
	hits := countHits(xMin, xMax, yMin, yMax)
	if hits != 112 {
		t.Errorf("countHits(xMin, xMax, yMin, yMax) = %d; want 112", hits)
	}
}
