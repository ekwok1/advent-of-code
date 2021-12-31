package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestFoldAndGetUniqueCoordinatesOneFold(t *testing.T) {
	file, transparentPaper := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	foldInstructionsIndex := utilities.IndexOf("fold", &transparentPaper)
	strCoordinates := transparentPaper[:foldInstructionsIndex]
	instructions := transparentPaper[foldInstructionsIndex:]

	uniqueCoordinates := foldAndGetUniqueCoordinates(&strCoordinates, &instructions, 1)
	if len(uniqueCoordinates) != 17 {
		t.Errorf("foldAndGetUniqueCoordinates(&strCoordinates, &instructions, 1) = %d; want 17", len(uniqueCoordinates))
	}
}

func TestFoldAndGetUniqueCoordinatesTwoFold(t *testing.T) {
	file, transparentPaper := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	foldInstructionsIndex := utilities.IndexOf("fold", &transparentPaper)
	strCoordinates := transparentPaper[:foldInstructionsIndex]
	instructions := transparentPaper[foldInstructionsIndex:]

	uniqueCoordinates := foldAndGetUniqueCoordinates(&strCoordinates, &instructions, 2)
	if len(uniqueCoordinates) != 16 {
		t.Errorf("foldAndGetUniqueCoordinates(&strCoordinates, &instructions, 2) = %d; want 16", len(uniqueCoordinates))
	}
}
