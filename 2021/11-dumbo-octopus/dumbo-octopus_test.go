package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestSimulate10Days(t *testing.T) {
	file, initialEnergyGrid := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	octopusGrid := initialOctopusEnergyGrid(&initialEnergyGrid)
	flashes := simulateDays(10, &octopusGrid)
	if flashes != 204 {
		t.Errorf("simulateDays(10, &octopusGrid) = %d; want 204", flashes)
	}
}

func TestSimulate100Days(t *testing.T) {
	file, initialEnergyGrid := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	octopusGrid := initialOctopusEnergyGrid(&initialEnergyGrid)
	flashes := simulateDays(100, &octopusGrid)
	if flashes != 1656 {
		t.Errorf("simulateDays(100, &octopusGrid) = %d; want 1656", flashes)
	}
}
