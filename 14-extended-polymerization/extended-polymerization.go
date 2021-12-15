package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, polymerData := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	polymerInsertionRules := polymerData[1:]
	polymerInsertionRuleMap := make(map[string]string)
	for _, rule := range polymerInsertionRules {
		insertionRule := strings.Split(rule, " -> ")
		polymerInsertionRuleMap[insertionRule[0]] = insertionRule[1]
	}

	elementCounter := make(map[string]int)
	polymerTemplate := polymerData[0]
	polymerTracker := make(map[string]int)

	for i := 0; i < len(polymerTemplate)-1; i++ {
		element1 := string(polymerTemplate[i])
		element2 := string(polymerTemplate[i+1])
		elementCounter[element1]++
		if i == len(polymerTemplate)-2 {
			elementCounter[element2]++
		}

		pair := element1 + element2
		polymerTracker[pair]++
	}

	for steps := 0; steps < 40; steps++ {
		newPolymerTracker := make(map[string]int)

		for k, v := range polymerTracker {
			insert := polymerInsertionRuleMap[k]
			elementCounter[insert] += v

			newPairLeft := string(k[0]) + insert
			newPairRight := insert + string(k[1])

			newPolymerTracker[newPairLeft] += v
			newPolymerTracker[newPairRight] += v
		}

		polymerTracker = newPolymerTracker
	}

	mostCommon := math.MinInt
	leastCommon := math.MaxInt
	for _, v := range elementCounter {
		if v > mostCommon {
			mostCommon = v
		} else if v < leastCommon {
			leastCommon = v
		}
	}

	fmt.Println(mostCommon)
	fmt.Println(leastCommon)

	fmt.Println(mostCommon - leastCommon)

	// NAIVE SOLUTION
	// polymerTemplate := polymerData[0]
	// polymer := polymerTemplate
	// newPolymer := ""
	// for step := 0; step < 10; step++ {
	// 	for i := 0; i < len(polymer)-1; i++ {
	// 		pair := string(polymer[i]) + string(polymer[i+1])
	// 		insert := polymerInsertionRuleMap[pair]

	// 		newPolymer += string(polymer[i]) + insert
	// 		if i == len(polymer)-2 {
	// 			newPolymer += string(polymer[i+1])
	// 		}
	// 	}
	// 	polymer = newPolymer
	// 	newPolymer = ""
	// }

	// polymerCountMap := make(map[string]int)
	// for i := 0; i < len(polymer); i++ {
	// 	polymerCountMap[string(polymer[i])]++
	// }

	// fmt.Println(polymerCountMap)
}
