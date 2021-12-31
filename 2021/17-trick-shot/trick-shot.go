package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, targetArea := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	xMin, xMax, yMin, yMax := getCoordinateRange(targetArea[0])

	maxHeight := getMaxHeight(yMin)
	fmt.Println("Max Height:", maxHeight)

	hits := countHits(xMin, xMax, yMin, yMax)
	fmt.Println("Initial velocities that hit target:", hits)
}

func countHits(xMin, xMax, yMin, yMax int) (hits int) {
	minXVelocity := getMinXVelocity(xMin)
	maxXVelocity := xMax
	minYVelocity := yMin
	maxYVelocity := getMaxYVelocity(yMin)

	for xVelocity := minXVelocity; xVelocity <= maxXVelocity; xVelocity++ {
		for yVelocity := minYVelocity; yVelocity <= maxYVelocity; yVelocity++ {
			if hitsTarget(xVelocity, yVelocity, xMin, xMax, yMin, yMax) {
				hits++
			}
		}
	}

	return
}

func hitsTarget(xVelocity, yVelocity, xMin, xMax, yMin, yMax int) bool {
	xPosition, yPosition := 0, 0
	for !isLong(xPosition, yPosition, xMax, yMin) {
		xPosition += xVelocity
		yPosition += yVelocity
		if xVelocity != 0 {
			xVelocity--
		}
		yVelocity--

		if xPosition >= xMin && xPosition <= xMax && yPosition <= yMax && yPosition >= yMin {
			return true
		}
	}

	return false
}

func isLong(xPosition, yPosition, xMax, yMin int) bool {
	return xPosition > xMax || yPosition < yMin
}

func getMaxYVelocity(yMin int) int {
	return int(math.Abs(float64(yMin)) - 1)
}

func getMinXVelocity(xMin int) int {
	sum := 0
	i := 1
	for {
		sum += i

		if sum > xMin {
			break
		}

		i++
	}

	return i
}

func getMaxHeight(yMin int) (maxheight int) {
	maxYVelocity := getMaxYVelocity(yMin)
	for i := 1; i <= maxYVelocity; i++ {
		maxheight += i
	}

	return
}

func getCoordinateRange(targetArea string) (xMin, xMax, yMin, yMax int) {
	fmt.Sscanf(targetArea, "target area: x=%d..%d, y=%d..%d", &xMin, &xMax, &yMin, &yMax)
	return
}
