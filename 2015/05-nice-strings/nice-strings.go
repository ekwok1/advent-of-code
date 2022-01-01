package main

import (
	"fmt"
	"strings"

	_rune "github.com/ekwok1/aoc-2021/utilities/rune"
	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, strings := scanner.ScanStringsFromFile("./input.txt")
	defer file.Close()

	niceStrings := 0
	actuallyNiceStrings := 0
	for _, str := range strings {
		if isNice(str) {
			niceStrings++
		}

		if isNiceV2(str) {
			actuallyNiceStrings++
		}
	}

	fmt.Println("There are", niceStrings, "nice strings")
	fmt.Println("There are actually", actuallyNiceStrings, "nice strings")
}

func isNiceV2(str string) bool {
	return hasLetterSandwich(str) && hasMultiplePairs(str)
}

func hasLetterSandwich(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}

	return false
}

func hasMultiplePairs(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		pair := string(str[i]) + string(str[i+1])
		if strings.Contains(str[i+2:], pair) {
			return true
		}
	}

	return false
}

func isNice(str string) bool {
	return !hasInvalidPair(str) && hasConsecutiveLetters(str) && hasVowels(str, 3)
}

func hasConsecutiveLetters(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			return true
		}
	}

	return false
}

func hasInvalidPair(str string) bool {
	invalidPairs := []string{"ab", "cd", "pq", "xy"}
	for _, invalidPair := range invalidPairs {
		if strings.Contains(str, invalidPair) {
			return true
		}
	}

	return false
}

func hasVowels(str string, min int) bool {
	vowelCount := 0

	for _, r := range str {
		if _rune.IsVowel(r) {
			vowelCount++
		}
	}

	return vowelCount >= min
}
