package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestDive(t *testing.T) {
	file, commands := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	product := dive(commands)
	if product != 150 {
		t.Errorf("countIncreases(measurements) = %d; want 150", product)
	}
}

func TestAimedDive(t *testing.T) {
	file, commands := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	product := aimedDive(commands)
	if product != 900 {
		t.Errorf("countIncreases(measurements) = %d; want 900", product)
	}
}
