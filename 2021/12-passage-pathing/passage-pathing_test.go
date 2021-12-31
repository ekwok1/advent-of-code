package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCountCavePathsNoReturnTestInput1(t *testing.T) {
	file, connections := scanner.ScanStringsFromFile("test-input-1.txt")
	defer file.Close()

	caveSystem := initializeCaveSystem(&connections)

	start := caveSystem.caves["start"]
	end := caveSystem.caves["end"]

	pathsNoReturn := countCavePathsNoReturn(caveSystem, start, end, start)
	if pathsNoReturn != 10 {
		t.Errorf("countCavePathsNoReturn(caveSystem, start, end, start) = %d; want 10", pathsNoReturn)
	}
}

func TestCountCavePathsNoReturnTestInput2(t *testing.T) {
	file, connections := scanner.ScanStringsFromFile("test-input-2.txt")
	defer file.Close()

	caveSystem := initializeCaveSystem(&connections)

	start := caveSystem.caves["start"]
	end := caveSystem.caves["end"]

	pathsNoReturn := countCavePathsNoReturn(caveSystem, start, end, start)
	if pathsNoReturn != 19 {
		t.Errorf("countCavePathsNoReturn(caveSystem, start, end, start) = %d; want 19", pathsNoReturn)
	}
}

func TestCountCavePathsNoReturnTestInput3(t *testing.T) {
	file, connections := scanner.ScanStringsFromFile("test-input-3.txt")
	defer file.Close()

	caveSystem := initializeCaveSystem(&connections)

	start := caveSystem.caves["start"]
	end := caveSystem.caves["end"]

	pathsNoReturn := countCavePathsNoReturn(caveSystem, start, end, start)
	if pathsNoReturn != 226 {
		t.Errorf("countCavePathsNoReturn(caveSystem, start, end, start) = %d; want 226", pathsNoReturn)
	}
}

func TestCountCavePathsOneReturnTestInput1(t *testing.T) {
	file, connections := scanner.ScanStringsFromFile("test-input-1.txt")
	defer file.Close()

	caveSystem := initializeCaveSystem(&connections)

	start := caveSystem.caves["start"]
	end := caveSystem.caves["end"]

	pathsOneReturn := countCavePathsOneReturn(caveSystem, start, end, start)
	if pathsOneReturn != 36 {
		t.Errorf("countCavePathsOneReturn(caveSystem, start, end, start) = %d; want 36", pathsOneReturn)
	}
}

func TestCountCavePathsOneReturnTestInput2(t *testing.T) {
	file, connections := scanner.ScanStringsFromFile("test-input-2.txt")
	defer file.Close()

	caveSystem := initializeCaveSystem(&connections)

	start := caveSystem.caves["start"]
	end := caveSystem.caves["end"]

	pathsOneReturn := countCavePathsOneReturn(caveSystem, start, end, start)
	if pathsOneReturn != 103 {
		t.Errorf("countCavePathsOneReturn(caveSystem, start, end, start) = %d; want 103", pathsOneReturn)
	}
}

func TestCountCavePathsOneReturnTestInput3(t *testing.T) {
	file, connections := scanner.ScanStringsFromFile("test-input-3.txt")
	defer file.Close()

	caveSystem := initializeCaveSystem(&connections)

	start := caveSystem.caves["start"]
	end := caveSystem.caves["end"]

	pathsOneReturn := countCavePathsOneReturn(caveSystem, start, end, start)
	if pathsOneReturn != 3509 {
		t.Errorf("countCavePathsOneReturn(caveSystem, start, end, start) = %d; want 3509", pathsOneReturn)
	}
}
