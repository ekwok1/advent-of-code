package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCalculateGreatestDiff10Steps(t *testing.T) {
	file, polymerData := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	polymerInsertionRuleList := polymerData[1:]
	polymerInsertionRules := initializePolymerInsertionRuleMap(&polymerInsertionRuleList)

	polymerTemplate := polymerData[0]
	elementCounter, polymerTracker := initializeCounterAndTracker(&polymerTemplate)

	steps := 10
	step(steps, &elementCounter, &polymerTracker, &polymerInsertionRules)

	diff := calculateGreatestDiff(&elementCounter)
	if diff != 1588 {
		t.Errorf("calculateGreatestDiff(&elementCounter) = %d; want 1588", diff)
	}
}

func TestCalculateGreatestDiff40Steps(t *testing.T) {
	file, polymerData := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	polymerInsertionRuleList := polymerData[1:]
	polymerInsertionRules := initializePolymerInsertionRuleMap(&polymerInsertionRuleList)

	polymerTemplate := polymerData[0]
	elementCounter, polymerTracker := initializeCounterAndTracker(&polymerTemplate)

	steps := 40
	step(steps, &elementCounter, &polymerTracker, &polymerInsertionRules)

	diff := calculateGreatestDiff(&elementCounter)
	if diff != 2188189693529 {
		t.Errorf("calculateGreatestDiff(&elementCounter) = %d; want 2188189693529", diff)
	}
}
