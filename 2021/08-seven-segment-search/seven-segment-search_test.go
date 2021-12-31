package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCountTrivialPatterns(t *testing.T) {
	file, patternData := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	count := countTrivialPatterns(&patternData)
	if count != 26 {
		t.Errorf("countTrivialPatterns(&patternData) = %d; want 26", count)
	}
}

func TestCalculateTotalOutput(t *testing.T) {
	file, patternData := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	total := calculateTotalOutput(&patternData)
	if total != 61229 {
		t.Errorf("countTrivialPatterns(&patternData) = %d; want 61229", total)
	}
}
