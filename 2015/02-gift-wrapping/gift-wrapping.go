package main

import (
	"fmt"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
	"github.com/ekwok1/aoc-2021/utilities/slice"
)

func main() {
	file, presents := scanner.ScanStringsFromFile("./input.txt")
	defer file.Close()

	wrappingPaper := 0
	ribbon := 0
	for _, present := range presents {
		l, w, h := parsePresent(present)
		p := Present{l, w, h}
		wrappingPaper += p.SurfaceArea() + p.SlackRequired()
		ribbon += p.RibbonLength() + p.BowLength()
	}

	fmt.Println("Total wrapping paper needed:", wrappingPaper)
	fmt.Println("Total ribbon needed:", ribbon)
}

// Present
type Present struct {
	length, width, height int
}

func (present *Present) SurfaceArea() int {
	l, w, h := present.length, present.width, present.height
	return 2*l*w + 2*w*h + 2*h*l
}

func (present *Present) SlackRequired() int {
	l, w, h := present.length, present.width, present.height
	areas := []int{l * w, w * h, h * l}
	return slice.MinInt(&areas)
}

func (present *Present) RibbonLength() int {
	l, w, h := present.length, present.width, present.height
	dimensions := []int{l, w, h}
	minInts := slice.MinInts(&dimensions, 2)

	return slice.SumInts(&minInts) * 2
}

func (present *Present) BowLength() int {
	l, w, h := present.length, present.width, present.height
	return l * w * h
}

// Parsing
func parsePresent(present string) (length, width, height int) {
	fmt.Sscanf(present, "%dx%dx%d", &length, &width, &height)
	return
}
