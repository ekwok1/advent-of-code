package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCountOn(t *testing.T) {
	file, rebootInstructions := scanner.ScanStringsFromFile("test-input-1.txt")
	defer file.Close()

	reactor := make(Reactor)
	initializeReactorTrivial(&rebootInstructions, &reactor)
	initializationAreaCount := countOn(&reactor)
	if initializationAreaCount != 590784 {
		t.Errorf("countOn(&reactor) = %d; want 590784", initializationAreaCount)
	}
}

func TestCountAllOn(t *testing.T) {
	file, rebootInstructions := scanner.ScanStringsFromFile("test-input-2.txt")
	defer file.Close()

	prisms := make([]Prism, 0)
	pointsOfInterestX, pointsOfInterestY, pointsOfInterestZ := getPointsOfInterest(&rebootInstructions, &prisms)
	compressedPointsMapX, compressedLengthMapX := compress(&pointsOfInterestX)
	compressedPointsMapY, compressedLengthMapY := compress(&pointsOfInterestY)
	compressedPointsMapZ, compressedLengthMapZ := compress(&pointsOfInterestZ)
	reactor2 := initializeReactor(&prisms, &compressedPointsMapX, &compressedPointsMapY, &compressedPointsMapZ)
	allCount := countAllOn(reactor2, &compressedLengthMapX, &compressedLengthMapY, &compressedLengthMapZ)
	if allCount != 2758514936282235 {
		t.Errorf("countAllOn(reactor2, &compressedLengthMapX, &compressedLengthMapY, &compressedLengthMapZ) = %d; want 2758514936282235", allCount)
	}
}
