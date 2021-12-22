package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestCalculateLowPointRisk(t *testing.T) {
	file, heightData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	heightmap := initializeHeightmap(&heightData)
	analyzeHeightmap(&heightmap)

	lowPointRisk := calculateLowPointRisk(&heightmap)
	if lowPointRisk != 15 {
		t.Errorf("calculateLowPointRisk(&heightmap) = %d; want 15", lowPointRisk)
	}
}

func TestLargestBasinProduct(t *testing.T) {
	file, heightData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	heightmap := initializeHeightmap(&heightData)
	analyzeHeightmap(&heightmap)

	product := largestBasinProduct(&heightmap, 3)
	if product != 1134 {
		t.Errorf("largestBasinProduct(&heightmap, 3) = %d; want 1134", product)
	}
}
