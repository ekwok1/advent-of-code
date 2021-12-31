package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, polymerData := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	polymerInsertionRuleList := polymerData[1:]
	polymerInsertionRules := initializePolymerInsertionRuleMap(&polymerInsertionRuleList)

	polymerTemplate := polymerData[0]
	elementCounter, polymerTracker := initializeCounterAndTracker(&polymerTemplate)

	steps := 40
	step(steps, &elementCounter, &polymerTracker, &polymerInsertionRules)

	diff := calculateGreatestDiff(&elementCounter)
	fmt.Println("Diff:", diff)
}

func calculateGreatestDiff(elementCounter *map[string]int) int {
	mostCommon := math.MinInt
	leastCommon := math.MaxInt
	for _, v := range *elementCounter {
		if v > mostCommon {
			mostCommon = v
		} else if v < leastCommon {
			leastCommon = v
		}
	}

	return mostCommon - leastCommon
}

func step(steps int, elementCounter, polymerTracker *map[string]int, polymerInsertionRules *map[string]string) {
	for s := 0; s < steps; s++ {
		newPolymerTracker := make(map[string]int)

		for polymer, count := range *polymerTracker {
			insert := (*polymerInsertionRules)[polymer]
			(*elementCounter)[insert] += count

			newPairLeft := string(polymer[0]) + insert
			newPairRight := insert + string(polymer[1])

			newPolymerTracker[newPairLeft] += count
			newPolymerTracker[newPairRight] += count
		}

		polymerTracker = &newPolymerTracker
	}
}

func initializeCounterAndTracker(polymerTemplate *string) (map[string]int, map[string]int) {
	elementCounter := make(map[string]int)
	polymerTracker := make(map[string]int)

	for i := 0; i < len(*polymerTemplate)-1; i++ {
		element1 := string((*polymerTemplate)[i])
		element2 := string((*polymerTemplate)[i+1])

		elementCounter[element1]++
		if i == len((*polymerTemplate))-2 {
			elementCounter[element2]++
		}

		polymerTracker[element1+element2]++
	}

	return elementCounter, polymerTracker
}

func initializePolymerInsertionRuleMap(polymerInsertionRuleList *[]string) map[string]string {
	polymerInsertionRules := make(map[string]string)

	for _, rule := range *polymerInsertionRuleList {
		fromTo := strings.Split(rule, " -> ")
		from, to := fromTo[0], fromTo[1]
		polymerInsertionRules[from] = to
	}

	return polymerInsertionRules
}
