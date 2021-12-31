package main

import (
	"fmt"
	"os"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, data := scanner.ScanStringFromFile(os.Args[1])
	defer file.Close()

	lisp := lisp{data}
	fmt.Println("Floor:", lisp.GetFloor())
	fmt.Println("Position:", lisp.GetPosition(-1))
}

func (lisp *lisp) GetFloor() int {
	counts := make(map[rune]int)
	for _, r := range lisp.str {
		counts[r]++
	}

	return counts['('] - counts[')']
}

func (lisp *lisp) GetPosition(floor int) int {
	counts := make(map[rune]int)
	for i, r := range lisp.str {
		counts[r]++
		if counts['(']-counts[')'] == floor {
			return i + 1
		}
	}

	return -1
}

type lisp struct {
	str string
}
