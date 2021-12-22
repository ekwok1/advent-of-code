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

	total := calculateTotalOutput(&patternData)
	fmt.Println("Total:", total)
}

func calculateTotalOutput(patternData *[]string) (total int) {
	for _, patternDataRow := range *patternData {
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
			if isOne(output) {
				stringifiedNumber += "1"
			} else if isFour(output) {
				stringifiedNumber += "4"
			} else if isSeven(output) {
				stringifiedNumber += "7"
			} else if isEight(output) {
				stringifiedNumber += "8"
			} else if isZero(output, &signalPatterns) {
				stringifiedNumber += "0"
			} else if isSix(output, &signalPatterns) {
				stringifiedNumber += "6"
			} else if isNine(output, &signalPatterns) {
				stringifiedNumber += "9"
			} else if isTwo(output, &signalPatterns) {
				stringifiedNumber += "2"
			} else if isThree(output, &signalPatterns) {
				stringifiedNumber += "3"
			} else if isFive(output, &signalPatterns) {
				stringifiedNumber += "5"
			}
		}

		number, _ := strconv.Atoi(stringifiedNumber)
		total += number
	}

	return
}

func countTrivialPatterns(patternData *[]string) (count int) {
	for _, patternDataRow := range *patternData {
		outputs := strings.Fields(strings.Split(patternDataRow, "|")[1])

		for _, output := range outputs {
			if isOne(output) || isFour(output) || isSeven(output) || isEight(output) {
				count++
			}
		}
	}

	return
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

func isZero(output string, patternMap *map[int]string) bool {
	zeroPattern := (*patternMap)[0]
	return len(output) == 6 && hasAllSegments(zeroPattern, output)
}

func isOne(output string) bool {
	return len(output) == 2
}

func isTwo(output string, patternMap *map[int]string) bool {
	twoPattern := (*patternMap)[2]
	return len(output) == 5 && hasAllSegments(twoPattern, output)
}

func isThree(output string, patternMap *map[int]string) bool {
	threePattern := (*patternMap)[3]
	return len(output) == 5 && hasAllSegments(threePattern, output)
}

func isFour(output string) bool {
	return len(output) == 4
}

func isFive(output string, patternMap *map[int]string) bool {
	fivePattern := (*patternMap)[5]
	return len(output) == 5 && hasAllSegments(fivePattern, output)
}

func isSix(output string, patternMap *map[int]string) bool {
	sixPattern := (*patternMap)[6]
	return len(output) == 6 && hasAllSegments(sixPattern, output)
}

func isSeven(output string) bool {
	return len(output) == 3
}

func isEight(output string) bool {
	return len(output) == 7
}

func isNine(output string, patternMap *map[int]string) bool {
	ninePattern := (*patternMap)[9]
	return len(output) == 6 && hasAllSegments(ninePattern, output)
}
