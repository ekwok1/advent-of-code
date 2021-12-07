package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

// OPTOMIZATIONS
// Only check for bingos after size of grid
// Only check for bingos of boards that don't have bingo yet
// Mark bingo function can be optomized with a break
// Is it possible to create a boardmap from square value to position to reduce looping for finds?

func main() {
	file, bingoInfo := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	// Setting Up
	allBingoBoardInfo := bingoInfo[1:]
	bingoBoards := createBingoBoards(&allBingoBoardInfo, 5)
	numbers := strings.Split(bingoInfo[0], ",")

	// Winner
	winningBoard, winningNumber := getWinningBoardAndNumber(&numbers, &bingoBoards)
	score := calculateScore(&winningBoard, winningNumber)
	fmt.Println("Winning Score:", score)

	// Last
	lastBoard, lastNumber := getLastBoardAndNumber(&numbers, &bingoBoards)
	lastScore := calculateScore(&lastBoard, lastNumber)
	fmt.Println("Last Score:", lastScore)
}

func getWinningBoardAndNumber(bingoNumbers *[]string, bingoBoards *[]BingoBoard) (winningBoard BingoBoard, winningNumber int) {
Loop:
	for _, number := range *bingoNumbers {
		bingoNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Cannot parse int from string:", bingoNumber)
		}
		markBingoBoards(bingoBoards, bingoNumber)

		for _, bingoBoard := range *bingoBoards {
			bingo := checkBingo(&bingoBoard, 5)
			if bingo {
				winningBoard = bingoBoard
				winningNumber = bingoNumber
				break Loop
			}
		}
	}

	return
}

func getLastBoardAndNumber(bingoNumbers *[]string, bingoBoards *[]BingoBoard) (lastBoard BingoBoard, lastNumber int) {
	totalBoards := len(*bingoBoards)
	bingos := 0

Loop:
	for _, bingoNumber := range *bingoNumbers {
		bingoNumber, err := strconv.Atoi(bingoNumber)
		if err != nil {
			fmt.Println("Cannot parse int from string:", bingoNumber)
		}
		markBingoBoards(bingoBoards, bingoNumber)

		for i, bingoBoard := range *bingoBoards {
			bingo := checkBingo(&bingoBoard, 5)
			if bingo && !bingoBoard.bingo {
				(*bingoBoards)[i].bingo = true
				bingos++

				if bingos == totalBoards {
					lastBoard = bingoBoard
					lastNumber = bingoNumber
					break Loop
				}
			}
		}
	}

	return
}

func markBingoBoards(bingoBoards *[]BingoBoard, bingoNumber int) {
	for _, board := range *bingoBoards {
		for j, row := range board.board {
			for k, square := range row {
				if square.value == bingoNumber {
					board.board[j][k].selected = true
				}
			}
		}
	}
}

func createBingoBoards(allBingoRows *[]string, size int) (bingoBoards []BingoBoard) {
	rows := 1
	var bingoBoard [][]BingoSpace

	for _, row := range *allBingoRows {
		spaces := strings.Fields(row)

		var bingoRow []BingoSpace
		for _, space := range spaces {
			value, err := strconv.Atoi(space)
			if err != nil {
				fmt.Println("Cannot parse int from string:", space)
			}
			bingoRow = append(bingoRow, BingoSpace{value: value})
		}

		bingoBoard = append(bingoBoard, bingoRow)

		if rows == size {
			bingoBoards = append(bingoBoards, BingoBoard{board: bingoBoard})
			rows = 1
			bingoBoard = nil
			continue
		}

		rows++
	}

	return
}

func checkBingo(bingoBoard *BingoBoard, size int) bool {
	bingo := false

	// Check horizontal
	for i := 0; i < size; i++ {
		bingo = true

		for j := 0; j < size; j++ {
			if !bingoBoard.board[i][j].selected {
				bingo = false
				break
			}
		}

		if bingo {
			return bingo
		}
	}

	if !bingo {
		// Check vertical
		for i := 0; i < size; i++ {
			bingo = true

			for j := 0; j < size; j++ {
				if !bingoBoard.board[j][i].selected {
					bingo = false
					break
				}
			}

			if bingo {
				return bingo
			}
		}
	}

	middle := (1+size)/2 - 1
	if bingoBoard.board[middle][middle].selected {
		if !bingo {
			// Check top left diagonal
			for i := 0; i < size; i++ {
				bingo = true

				if !bingoBoard.board[i][i].selected {
					bingo = false
					break
				}
			}

			if bingo {
				return bingo
			}
		}

		if !bingo {
			// Check top right diagonal
			for i := size - 1; i >= 0; i-- {
				bingo = true

				if !bingoBoard.board[i][i].selected {
					bingo = false
					break
				}
			}

			if bingo {
				return bingo
			}
		}
	}

	return false
}

func calculateScore(bingoBoard *BingoBoard, winningNumber int) (score int) {
	for _, row := range bingoBoard.board {
		for _, square := range row {
			if !square.selected {
				score += square.value
			}
		}
	}
	return score * winningNumber
}

type BingoSpace struct {
	value    int
	selected bool
}

type BingoBoard struct {
	board [][]BingoSpace
	bingo bool
}
