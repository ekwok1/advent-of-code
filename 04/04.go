package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, bingoInfo := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	bingoNumbers := strings.Split(bingoInfo[0], ",")
	bingoInfo = bingoInfo[1:]

	var bingoBoards []BingoBoard
	var bingoBoard [][]BingoSquare

	// Creating bingo boards
	rows := 1
	for _, info := range bingoInfo {
		if info == "" {
			continue
		}

		fields := strings.Fields(info)

		var bingoRow []BingoSquare
		for _, field := range fields {
			value, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Cannot parse int from string:", field)
			}
			bingoRow = append(bingoRow, BingoSquare{value: value})
		}

		bingoBoard = append(bingoBoard, bingoRow)

		if rows == 5 {
			bingoBoards = append(bingoBoards, BingoBoard{board: bingoBoard})
			rows = 1
			bingoBoard = nil
			continue
		}

		rows++
	}

	// Playing Bingo
	// Mark a number
	// Check if any bingos
	// Repeat

	var winningNumber int
	var winningBoard int
Loop:
	for _, bingoNumber := range bingoNumbers {
		bingoNumber, _ := strconv.Atoi(bingoNumber)
		bingoBoards = markBingoBoards(bingoBoards, bingoNumber)

		for i, bingoBoard := range bingoBoards {
			bingo := checkBingo(bingoBoard)
			if bingo {
				winningNumber = bingoNumber
				winningBoard = i
				break Loop
			}
		}
	}

	fmt.Println(winningBoard)
	fmt.Println(bingoBoards[winningBoard])

	sum := 0
	for _, row := range bingoBoards[winningBoard].board {
		for _, square := range row {
			if !square.checked {
				sum += square.value
			}
		}
	}

	fmt.Println(winningNumber)
	fmt.Println(sum)

	fmt.Println(winningNumber * sum)
}

func markBingoBoards(bingoBoards []BingoBoard, bingoNumber int) []BingoBoard {
	for _, board := range bingoBoards {
		for j, row := range board.board {
			for k, square := range row {
				if square.value == bingoNumber {
					board.board[j][k].checked = true
					// bingoBoards[i][j][k].checked = true
				}
			}
		}
	}

	return bingoBoards
}

func checkBingo(bingoBoard BingoBoard) bool {
	bingo := false

	// Check horizontal
	for i := 0; i < 5; i++ {
		bingo = true

		for j := 0; j < 5; j++ {
			if !bingoBoard.board[i][j].checked {
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
		for i := 0; i < 5; i++ {
			bingo = true

			for j := 0; j < 5; j++ {
				if !bingoBoard.board[j][i].checked {
					bingo = false
					break
				}
			}

			if bingo {
				return bingo
			}
		}
	}

	if bingoBoard.board[2][2].checked {
		if !bingo {
			// Check top left diagonal
			for i := 0; i < 5; i++ {
				bingo = true

				if !bingoBoard.board[i][i].checked {
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
			for i := 4; i >= 0; i-- {
				bingo = true

				if !bingoBoard.board[i][i].checked {
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

type BingoSquare struct {
	value   int
	checked bool
}

type BingoBoard struct {
	board [][]BingoSquare
	bingo bool
}
