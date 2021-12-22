package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestCountIncreases(t *testing.T) {
	file, measurements := utilities.ScanIntsFromFile("test-input.txt")
	defer file.Close()

	increases := countIncreases(measurements)
	if increases != 7 {
		t.Errorf("countIncreases(measurements) = %d; want 7", increases)
	}
}

func TestCountSlidingIncreases(t *testing.T) {
	file, measurements := utilities.ScanIntsFromFile("test-input.txt")
	defer file.Close()

	increases := countSlidingIncreases(measurements)
	if increases != 5 {
		t.Errorf("countSlidingIncreases(measurements) = %d; want 5", increases)
	}
}
