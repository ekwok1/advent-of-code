package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestGetMinimumTotalRisk(t *testing.T) {
	file, riskData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	riskGrid := initialRiskGrid(&riskData)

	minimumTotalCost := getMinimumTotalRisk(&riskGrid)
	if minimumTotalCost != 40 {
		t.Errorf("calculateGreatestDiff(&elementCounter) = %d; want 40", minimumTotalCost)
	}
}

func TestGetMinimumTotalRiskFullCave(t *testing.T) {
	file, riskData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	riskGrid := initialRiskGrid(&riskData)

	minimumTotalCostFullCave := getMinimumTotalRiskFullCave(&riskGrid)
	if minimumTotalCostFullCave != 315 {
		t.Errorf("calculateGreatestDiff(&elementCounter) = %d; want 315", minimumTotalCostFullCave)
	}
}
