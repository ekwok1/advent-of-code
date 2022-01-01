package main

import (
	"fmt"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, lisp := scanner.ScanStringFromFile("./input.txt")
	defer file.Close()

	fmt.Println("Floor:", GetFloor(&lisp))
	fmt.Println("Position:", GetPosition(&lisp, -1))
}

func GetFloor(lisp *string) int {
	counts := make(map[rune]int)
	for _, r := range *lisp {
		counts[r]++
	}

	return counts['('] - counts[')']
}

func GetPosition(lisp *string, floor int) int {
	counts := make(map[rune]int)
	for i, r := range *lisp {
		counts[r]++
		if counts['(']-counts[')'] == floor {
			return i + 1
		}
	}

	panic(fmt.Sprintf("Never reached floor %d", floor))
}
