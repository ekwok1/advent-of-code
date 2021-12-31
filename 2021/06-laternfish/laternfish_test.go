package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
	"github.com/ekwok1/aoc-2021/utilities/slice"
)

func TestSimulate18Days(t *testing.T) {
	file, initialTimes := scanner.ScanIntsFromDelimitedString("test-input.txt", ",")
	defer file.Close()

	school := setupInitialState(initialTimes)
	simulateDays(&school, 18)
	total := slice.SumInts(&school)

	if total != 26 {
		t.Errorf("utilities.CountTotal(&school) = %d; want 26", total)
	}
}

func TestSimulate80Days(t *testing.T) {
	file, initialTimes := scanner.ScanIntsFromDelimitedString("test-input.txt", ",")
	defer file.Close()

	school := setupInitialState(initialTimes)
	simulateDays(&school, 80)
	total := slice.SumInts(&school)

	if total != 5934 {
		t.Errorf("utilities.CountTotal(&school) = %d; want 5934", total)
	}
}

func TestSimulate256Days(t *testing.T) {
	file, initialTimes := scanner.ScanIntsFromDelimitedString("test-input.txt", ",")
	defer file.Close()

	school := setupInitialState(initialTimes)
	simulateDays(&school, 256)
	total := slice.SumInts(&school)

	if total != 26984457539 {
		t.Errorf("utilities.CountTotal(&school) = %d; want 26984457539", total)
	}
}
