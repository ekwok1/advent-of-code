package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, chunks := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	syntaxMap := initializeSyntaxMap()
	illegalCharacters, incompleteChunks := analyzeChunks(&chunks, &syntaxMap)

	syntaxErrorScore := calculateSyntaxErrorScore(&illegalCharacters)
	fmt.Println("Illegal Syntax Score:", syntaxErrorScore)

	completionScore := calculateCompletionScore(&incompleteChunks, &syntaxMap)
	fmt.Println("Completion Score:", completionScore)
}

func calculateCompletionScore(incompleteChunks *[]string, syntaxMap *map[string]string) int {
	scoreMap := make(map[string]int)
	scoreMap[")"] = 1
	scoreMap["]"] = 2
	scoreMap["}"] = 3
	scoreMap[">"] = 4

	var scores []int

	for _, incompleteChunk := range *incompleteChunks {
		stack := strings.Split(incompleteChunk, "")

		for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
			stack[i], stack[j] = stack[j], stack[i]
		}

		completionString := ""
		for _, character := range stack {
			completionString += (*syntaxMap)[character]
		}

		closingCharacters := strings.Split(completionString, "")
		score := 0
		for _, closingCharacter := range closingCharacters {
			score *= 5
			score += scoreMap[closingCharacter]
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)
	middleIndex := len(scores) / 2
	return scores[middleIndex]
}

func calculateSyntaxErrorScore(illegalCharacters *[]string) (score int) {
	scoreMap := make(map[string]int)
	scoreMap[")"] = 3
	scoreMap["]"] = 57
	scoreMap["}"] = 1197
	scoreMap[">"] = 25137

	for _, illegalCharacter := range *illegalCharacters {
		score += scoreMap[illegalCharacter]
	}

	return
}

func analyzeChunks(chunks *[]string, syntaxMap *map[string]string) (illegalCharacters []string, incompleteChunks []string) {
	for _, chunk := range *chunks {
		illegalCharacter, incompleteChunk := analyzeChunk(chunk, syntaxMap)
		if illegalCharacter != "" {
			illegalCharacters = append(illegalCharacters, illegalCharacter)
		}

		if incompleteChunk != "" {
			incompleteChunks = append(incompleteChunks, incompleteChunk)
		}
	}

	return
}

func analyzeChunk(chunk string, syntaxMap *map[string]string) (illegalCharacter string, incompleteChunk string) {
	var stack []string

	for i := 0; i < len(chunk); i++ {
		c := string(chunk[i])
		if _, ok := (*syntaxMap)[c]; ok {
			stack = append(stack, c)
			continue
		}

		lastIndex := len(stack) - 1
		lastSyntaxCharacter := stack[lastIndex]
		stack = stack[:lastIndex]

		switch c {
		case ")":
			if lastSyntaxCharacter != "(" {
				// fmt.Printf("Expected %s, but found %s instead", (*syntaxMap)[lastSyntaxCharacter], ")")
				illegalCharacter = ")"
			}
		case "}":
			if lastSyntaxCharacter != "{" {
				// fmt.Printf("Expected %s, but found %s instead", (*syntaxMap)[lastSyntaxCharacter], "}")
				illegalCharacter = "}"
			}
		case "]":
			if lastSyntaxCharacter != "[" {
				// fmt.Printf("Expected %s, but found %s instead", (*syntaxMap)[lastSyntaxCharacter], "]")
				illegalCharacter = "]"
			}
		case ">":
			if lastSyntaxCharacter != "<" {
				// fmt.Printf("Expected %s, but found %s instead", (*syntaxMap)[lastSyntaxCharacter], ">")
				illegalCharacter = ">"
			}
		}
	}

	if illegalCharacter == "" {
		incompleteChunk = strings.Join(stack, "")
	}

	return
}

func initializeSyntaxMap() map[string]string {
	syntaxMap := make(map[string]string)

	syntaxMap["("] = ")"
	syntaxMap["["] = "]"
	syntaxMap["{"] = "}"
	syntaxMap["<"] = ">"

	return syntaxMap
}
