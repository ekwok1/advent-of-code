package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCountLitPixels_2Enhacements(t *testing.T) {
	file, allData := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	algorithm := allData[0]
	inputImage := allData[1:]
	image := initializeImage(&inputImage)

	enhancements := 2
	enhancer(enhancements, &algorithm, &image)

	litPixels := countLitPixels(image)
	if litPixels != 35 {
		t.Errorf("countLitPixels(image) = %d; want 35", litPixels)
	}
}

func TestCountLitPixels_50Enhacements(t *testing.T) {
	file, allData := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	algorithm := allData[0]
	inputImage := allData[1:]
	image := initializeImage(&inputImage)

	enhancements := 50
	enhancer(enhancements, &algorithm, &image)

	litPixels := countLitPixels(image)
	if litPixels != 3351 {
		t.Errorf("countLitPixels(image) = %d; want 3351", litPixels)
	}
}
