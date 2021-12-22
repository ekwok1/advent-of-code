package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, rebootInstructions := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	reactor := make(Reactor)
	initializeReactorTrivial(&rebootInstructions, &reactor)
	count := countOn(&reactor)
	fmt.Println("Count on:", count)
}

func countOn(reactor *Reactor) (count int) {
	for _, v := range *reactor {
		if v {
			count++
		}
	}

	return
}

func initializeReactorTrivial(rebootInstructions *[]string, reactor *Reactor) {
	for _, instruction := range *rebootInstructions {
		on, xMin, xMax, yMin, yMax, zMin, zMax := parseInstructions(instruction)

		if xMin < -50 || xMax > 50 || yMin < -50 || yMax > 50 || zMin < -50 || zMax > 50 {
			continue
		}

		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				for z := zMin; z <= zMax; z++ {
					(*reactor)[Vector3d{x, y, z}] = on
				}
			}
		}
	}
}

func parseInstructions(line string) (on bool, xMin, xMax, yMin, yMax, zMin, zMax int) {
	instructions := strings.Fields(line)

	on = instructions[0] == "on"
	fmt.Sscanf(instructions[1], "x=%d..%d,y=%d..%d,z=%d..%d",
		&xMin, &xMax, &yMin, &yMax, &zMin, &zMax)

	return
}

type Reactor map[Vector3d]bool

type Vector3d struct {
	x, y, z int
}
