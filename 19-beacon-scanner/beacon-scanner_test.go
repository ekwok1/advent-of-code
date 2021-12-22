package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestFindUniqueBeacons(t *testing.T) {
	file, allData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	beacons := GetBeacons(&allData)

	uniqueBeacons, _ := analyzeBeaconsAndScanners(beacons)
	if len(uniqueBeacons) != 79 {
		t.Errorf("analyzeBeaconsAndScanners(beacons) = %d; want 79", len(uniqueBeacons))
	}
}

func TestGetMaxManhattanDistance(t *testing.T) {
	file, allData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	beacons := GetBeacons(&allData)
	_, scannerVectors := analyzeBeaconsAndScanners(beacons)

	maxManhattanDistance := getMaxManhattanDistance(&scannerVectors)
	if maxManhattanDistance != 3621 {
		t.Errorf("analyzeBeaconsAndScanners(beacons) = %d; want 3621", maxManhattanDistance)
	}
}
