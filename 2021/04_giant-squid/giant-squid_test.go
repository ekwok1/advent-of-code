package main

import (
	"strings"
	"testing"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func TestCalculateScoreOfWinningBoardAndNumber(t *testing.T) {
	file, bingoInfo := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	allBingoBoardInfo := bingoInfo[1:]
	bingoBoards := createBingoBoards(&allBingoBoardInfo, 5)
	numbers := strings.Split(bingoInfo[0], ",")

	winningBoard, winningNumber := getWinningBoardAndNumber(&numbers, &bingoBoards)

	score := calculateScore(&winningBoard, winningNumber)

	if score != 4512 {
		t.Errorf("calculateScore(&winningBoard, winningNumber) = %d; want 4512", score)
	}
}

func TestCalculateScoreOfLastBoardAndNumber(t *testing.T) {
	file, bingoInfo := scanner.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	allBingoBoardInfo := bingoInfo[1:]
	bingoBoards := createBingoBoards(&allBingoBoardInfo, 5)
	numbers := strings.Split(bingoInfo[0], ",")

	lastBoard, lastNumber := getLastBoardAndNumber(&numbers, &bingoBoards)

	score := calculateScore(&lastBoard, lastNumber)

	if score != 1924 {
		t.Errorf("calculateScore(&lastBoard, lastNumber) = %d; want 1924", score)
	}
}
