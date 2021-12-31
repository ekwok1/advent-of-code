package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
	"github.com/ekwok1/aoc-2021/utilities/slice"
)

func main() {
	file, positionData := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	player1Position, player2Position := getStartingPositions(&positionData)

	diceRolls := initializeDicerolls()
	score1, score2, totalRolls := play(player1Position, player2Position, 0, 0, &diceRolls, 0)
	product := productLowestScoreTotalRolls(score1, score2, totalRolls)
	fmt.Println("Product:", product)

	states := make(map[GameState][]int)
	winCounter := countWins(player1Position, player2Position, 0, 0, &states)
	mostWins := getMostWins(winCounter)
	fmt.Println("Most Wins:", mostWins)
}

func getMostWins(winCounter []int) int {
	if winCounter[0] > winCounter[1] {
		return winCounter[0]
	} else {
		return winCounter[1]
	}
}

func countWins(position1, position2, score1, score2 int, states *map[GameState][]int) []int {
	if score1 >= 21 {
		return []int{1, 0}
	}

	if score2 >= 21 {
		return []int{0, 1}
	}

	if val, ok := (*states)[GameState{position1, position2, score1, score2}]; ok {
		return val
	}

	counter := []int{0, 0}
	rolls := []int{1, 2, 3}
	for _, d1 := range rolls {
		for _, d2 := range rolls {
			for _, d3 := range rolls {
				roll := d1 + d2 + d3
				newP1 := land(position1, roll)
				newS1 := score1 + newP1

				nextCounter := countWins(position2, newP1, score2, newS1, states)
				counter = []int{counter[0] + nextCounter[1], counter[1] + nextCounter[0]}
			}
		}
	}

	(*states)[GameState{position1, position2, score1, score2}] = counter
	return counter
}

func productLowestScoreTotalRolls(score1, score2, rolls int) int {
	lowerScore := int(math.Min(float64(score1), float64(score2)))
	return lowerScore * rolls
}

func play(position1, position2, score1, score2 int, dicerolls *[]int, rolls int) (finalscore1, finalscore2, totalRolls int) {
	if score1 >= 1000 || score2 >= 1000 {
		return score1, score2, rolls
	}

	if len(*dicerolls) < 3 {
		(*dicerolls) = resetDice(*dicerolls)
	}

	roll := (*dicerolls)[:3]
	sum := slice.SumInts(&roll)
	lands := land(position1, sum)
	newScore := score1 + lands
	newPosition := lands
	(*dicerolls) = (*dicerolls)[3:]
	rolls += 3

	return play(position2, newPosition, score2, newScore, dicerolls, rolls)
}

func land(start, roll int) int {
	return ((start + roll - 1) % 10) + 1
}

func resetDice(remainingRolls []int) (rolls []int) {
	additionalRolls := make([]int, 100)
	for i := 1; i <= 100; i++ {
		additionalRolls[i-1] = i
	}

	rolls = append(rolls, remainingRolls...)
	rolls = append(rolls, additionalRolls...)
	return
}

func initializeDicerolls() (dicerolls []int) {
	for i := 1; i <= 100; i++ {
		dicerolls = append(dicerolls, i)
	}

	return
}

func getStartingPositions(positionData *[]string) (player1Position, player2Position int) {
	fmt.Sscanf((*positionData)[0], "Player 1 starting position: %d", &player1Position)
	fmt.Sscanf((*positionData)[1], "Player 2 starting position: %d", &player2Position)
	return
}

type GameState struct {
	p1, p2, s1, s2 int
}
