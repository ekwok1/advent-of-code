package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestCalcPowerConsumption(t *testing.T) {
	file, binaries := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	powerConsumption := calcPowerConsumption(binaries)
	if powerConsumption != 198 {
		t.Errorf("countIncreases(measurements) = %d; want 198", powerConsumption)
	}
}

func TestGetOxygenGeneratorRating(t *testing.T) {
	file, binaries := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	oxygenGeneratorRating := getOxygenGeneratorRating(binaries, 0)
	if oxygenGeneratorRating != 23 {
		t.Errorf("countIncreases(measurements) = %d; want 23", oxygenGeneratorRating)
	}
}

func TestGetCarbonDioxideScrubberRating(t *testing.T) {
	file, binaries := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	carbonDioxideScrubberRating := getCarbonDioxideScrubberRating(binaries, 0)
	if carbonDioxideScrubberRating != 10 {
		t.Errorf("countIncreases(measurements) = %d; want 10", carbonDioxideScrubberRating)
	}
}
