package main

import (
	"testing"

	"github.com/ekwok1/aoc-2021/utilities"
)

func TestProductLowestScoreTotalRolls(t *testing.T) {
	file, positionData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	player1Position, player2Position := getStartingPositions(&positionData)

	diceRolls := initializeDicerolls()
	score1, score2, totalRolls := play(player1Position, player2Position, 0, 0, &diceRolls, 0)
	product := productLowestScoreTotalRolls(score1, score2, totalRolls)
	if product != 739785 {
		t.Errorf("productLowestScoreTotalRolls(score1, score2, totalRolls) = %d; want 739785", product)
	}
}

func TestGetMostWins(t *testing.T) {
	file, positionData := utilities.ScanStringsFromFile("test-input.txt")
	defer file.Close()

	player1Position, player2Position := getStartingPositions(&positionData)

	states := make(map[GameState][]int)
	winCounter := countWins(player1Position, player2Position, 0, 0, &states)
	mostWins := getMostWins(winCounter)
	if mostWins != 444356092776315 {
		t.Errorf("getMostWins(winCounter) = %d; want 444356092776315", mostWins)
	}
}
