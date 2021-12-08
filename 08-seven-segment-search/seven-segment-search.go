package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, patternData := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	count := countTrivialPatterns(&patternData)
	fmt.Println("Count of 1, 4, 7, 8:", count)

	// OPTOMIZE
	// If none of the output values have length of 5 or 6, then we don't need to deduce
	total := 0
	for _, patternDataRow := range patternData {
		// Deduce segment patterns
		signalPatterns := make(map[int]string)
		signalBuckets := make([][]string, 8)

		fields := strings.Split(patternDataRow, "|")
		signals := strings.Fields(fields[0])

		for _, signal := range signals {
			signalBuckets[len(signal)] = append(signalBuckets[len(signal)], signal)
		}

		signalPatterns[1] = signalBuckets[2][0]
		signalPatterns[7] = signalBuckets[3][0]
		signalPatterns[4] = signalBuckets[4][0]
		signalPatterns[8] = signalBuckets[7][0]

		for _, sixLengthSignal := range signalBuckets[6] {
			if !hasAllSegments(sixLengthSignal, signalPatterns[1]) {
				signalPatterns[6] = sixLengthSignal
			} else if hasAllSegments(sixLengthSignal, signalPatterns[4]) {
				signalPatterns[9] = sixLengthSignal
			} else {
				signalPatterns[0] = sixLengthSignal
			}
		}

		for _, fiveLengthSignal := range signalBuckets[5] {
			if hasAllSegments(fiveLengthSignal, signalPatterns[1]) {
				signalPatterns[3] = fiveLengthSignal
			} else if hasAllSegments(signalPatterns[6], fiveLengthSignal) {
				signalPatterns[5] = fiveLengthSignal
			} else {
				signalPatterns[2] = fiveLengthSignal
			}
		}

		stringifiedNumber := ""
		outputs := strings.Fields(fields[1])
		for _, output := range outputs {
			outputLength := len(output)
			if isOne(outputLength) {
				stringifiedNumber += "1"
			} else if isFour(outputLength) {
				stringifiedNumber += "4"
			} else if isSeven(outputLength) {
				stringifiedNumber += "7"
			} else if isEight(outputLength) {
				stringifiedNumber += "8"
			} else if outputLength == 6 {
				if isZero(output, &signalPatterns) {
					stringifiedNumber += "0"
				} else if isSix(output, &signalPatterns) {
					stringifiedNumber += "6"
				} else if isNine(output, &signalPatterns) {
					stringifiedNumber += "9"
				}
			} else if outputLength == 5 {
				if isTwo(output, &signalPatterns) {
					stringifiedNumber += "2"
				} else if isThree(output, &signalPatterns) {
					stringifiedNumber += "3"
				} else if isFive(output, &signalPatterns) {
					stringifiedNumber += "5"
				}
			}
		}

		number, err := strconv.Atoi(stringifiedNumber)
		if err != nil {
			fmt.Println("Cannot parse int from string:", stringifiedNumber)
		}
		total += number
	}

	fmt.Println("Total:", total)
}

func hasAllSegments(pattern string, testPattern string) bool {
	patternSegments := make(map[string]bool)
	for _, segment := range pattern {
		patternSegments[string(segment)] = true
	}

	hasAll := true

	testSegments := strings.Split(testPattern, "")
	for _, segment := range testSegments {
		if !patternSegments[segment] {
			hasAll = false
		}
	}

	return hasAll
}

func countTrivialPatterns(patternData *[]string) (count int) {
	for _, patternDataRow := range *patternData {
		outputs := strings.Fields(strings.Split(patternDataRow, "|")[1])
		for _, output := range outputs {
			outputLength := len(output)
			if isOne(outputLength) || isFour(outputLength) || isSeven(outputLength) || isEight(outputLength) {
				count++
			}
		}
	}
	return
}

func isZero(output string, patternMap *map[int]string) bool {
	zeroPattern := (*patternMap)[0]
	return len(output) == 6 && hasAllSegments(zeroPattern, output)
}

func isOne(length int) bool {
	return length == 2
}

func isTwo(output string, patternMap *map[int]string) bool {
	twoPattern := (*patternMap)[2]
	return len(output) == 5 && hasAllSegments(twoPattern, output)
}

func isThree(output string, patternMap *map[int]string) bool {
	threePattern := (*patternMap)[3]
	return len(output) == 5 && hasAllSegments(threePattern, output)
}

func isFour(length int) bool {
	return length == 4
}

func isFive(output string, patternMap *map[int]string) bool {
	fivePattern := (*patternMap)[5]
	return len(output) == 5 && hasAllSegments(fivePattern, output)
}

func isSix(output string, patternMap *map[int]string) bool {
	sixPattern := (*patternMap)[6]
	return len(output) == 6 && hasAllSegments(sixPattern, output)
}

func isSeven(length int) bool {
	return length == 3
}

func isEight(length int) bool {
	return length == 7
}

func isNine(output string, patternMap *map[int]string) bool {
	ninePattern := (*patternMap)[9]
	return len(output) == 6 && hasAllSegments(ninePattern, output)
}
