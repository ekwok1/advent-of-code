package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCalculateSyntaxErrorScore(t *testing.T) {
	file, chunks := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	syntaxMap := initializeSyntaxMap()
	illegalCharacters, _ := analyzeChunks(&chunks, &syntaxMap)

	syntaxErrorScore := calculateSyntaxErrorScore(&illegalCharacters)
	if syntaxErrorScore != 26397 {
		t.Errorf("calculateSyntaxErrorScore(&illegalCharacters) = %d; want 26397", syntaxErrorScore)
	}
}

func TestCalculateCompletionScore(t *testing.T) {
	file, chunks := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	syntaxMap := initializeSyntaxMap()
	_, incompleteChunks := analyzeChunks(&chunks, &syntaxMap)

	completionScore := calculateCompletionScore(&incompleteChunks, &syntaxMap)
	if completionScore != 288957 {
		t.Errorf("calculateCompletionScore(&incompleteChunks, &syntaxMap) = %d; want 288957", completionScore)
	}
}
