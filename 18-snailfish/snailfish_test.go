package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestCalculateMagnitude(t *testing.T) {
	file, snailfishNumbers := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	magnitude := calculateMagnitude(snailfishNumbers)
	if magnitude != 4140 {
		t.Errorf("calculateMagnitude(snailfishNumbers) = %d; want 4140", magnitude)
	}
}

func TestFindGreatestMagnitude(t *testing.T) {
	file, snailfishNumbers := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	greatestMagnitude := findGreatestMagnitude(&snailfishNumbers)
	if greatestMagnitude != 3993 {
		t.Errorf("findGreatestMagnitude(&snailfishNumbers) = %d; want 3993", greatestMagnitude)
	}
}
