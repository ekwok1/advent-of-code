package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestDive(t *testing.T) {
	file, commands := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	product := dive(commands)
	if product != 150 {
		t.Errorf("dive(commands) = %d; want 150", product)
	}
}

func TestAimedDive(t *testing.T) {
	file, commands := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	product := aimedDive(commands)
	if product != 900 {
		t.Errorf("aimedDive(commands) = %d; want 900", product)
	}
}
