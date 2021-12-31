package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCalcMinimumBasicFuel(t *testing.T) {
	file, positions := scanner.ScanIntsFromDelimitedString("test-input.txt", ",")
	defer file.Close()

	min, max := getPositionRange(&positions)

	basicFuel := calcMinimumFuel(min, max, &positions, calcBasicFuel)
	if basicFuel != 37 {
		t.Errorf("calcMinimumFuel(min, max, &positions, calcBasicFuel) = %d; want 37", basicFuel)
	}
}

func TestCalcMininmumExpensiveFuel(t *testing.T) {
	file, positions := scanner.ScanIntsFromDelimitedString("test-input.txt", ",")
	defer file.Close()

	min, max := getPositionRange(&positions)

	expensiveFuel := calcMinimumFuel(min, max, &positions, calcExpensiveFuel)
	if expensiveFuel != 168 {
		t.Errorf("calcMinimumFuel(min, max, &positions, calcBasicFuel) = %d; want 168", expensiveFuel)
	}
}
