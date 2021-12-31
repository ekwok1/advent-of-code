package main

import (
	"fmt"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, data := scanner.ScanStringFromFile("./input.txt")
	defer file.Close()

	lisp := Lisp{data}
	fmt.Println("Floor:", lisp.GetFloor())
	fmt.Println("Position:", lisp.GetPosition(-1))
}

// Lisp
type Lisp struct {
	str string
}

func (lisp *Lisp) GetFloor() int {
	counts := make(map[rune]int)
	for _, r := range lisp.str {
		counts[r]++
	}

	return counts['('] - counts[')']
}

func (lisp *Lisp) GetPosition(floor int) int {
	counts := make(map[rune]int)
	for i, r := range lisp.str {
		counts[r]++
		if counts['(']-counts[')'] == floor {
			return i + 1
		}
	}

	panic(fmt.Sprintf("Never reached floor %d", floor))
}
