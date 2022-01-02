package main

import (
	"fmt"

	"github.com/ekwok1/aoc-2021/utilities/grid"
	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, instructions := scanner.ScanStringsFromFile("./input.txt")
	defer file.Close()

	lightGrid := grid.NewBoolGrid(1000, 1000)
	lightGrid2 := grid.NewIntGrid(1000, 1000)

	for _, instruction := range instructions {
		action, x1, x2, y1, y2 := parseInstruction(instruction)

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				switch action {
				case "on":
					lightGrid[y][x] = true
					lightGrid2[y][x]++
				case "off":
					lightGrid[y][x] = false
					if lightGrid2[y][x] == 0 {
						continue
					}
					lightGrid2[y][x]--
				case "toggle":
					lightGrid[y][x] = !lightGrid[y][x]
					lightGrid2[y][x] += 2
				default:
					panic(fmt.Sprintf("Do not know how to take action on %s", action))
				}
			}
		}
	}

	fmt.Println("Part 1 - Christmas lights that are on:", grid.CountBool(&lightGrid, true))
	fmt.Println("Part 2 - Christmas light brightness:", grid.GridSum(&lightGrid2))
}

// Parsing
func parseInstruction(instruction string) (action string, x1, y1, x2, y2 int) {
	fmt.Sscanf(instruction, "%s %d,%d through %d,%d", &action, &x1, &x2, &y1, &y2)
	return
}
