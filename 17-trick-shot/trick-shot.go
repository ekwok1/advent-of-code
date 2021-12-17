package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, targetArea := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	xMin, xMax, yMin, yMax := getCoordinateRange(targetArea[0])
	fmt.Println("xMin:", xMin, "xMax:", xMax, "yMin:", yMin, "yMax:", yMax)

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

func getCoordinateRange(targetArea string) (int, int, int, int) {
	coordinateRanges := strings.Split(targetArea, "target area: ")
	xRange := strings.Split(strings.Split(strings.Split(coordinateRanges[1], ", ")[0], "x=")[1], "..")
	yRange := strings.Split(strings.Split(strings.Split(coordinateRanges[1], ", ")[1], "y=")[1], "..")

	xMin, err := strconv.Atoi(xRange[0])
	if err != nil {
		fmt.Println("Cannot parse int from string:", xRange[0])
	}
	xMax, err := strconv.Atoi(xRange[1])
	if err != nil {
		fmt.Println("Cannot parse int from string:", xRange[1])
	}
	yMin, err := strconv.Atoi(yRange[0])
	if err != nil {
		fmt.Println("Cannot parse int from string:", yRange[0])
	}
	yMax, err := strconv.Atoi(yRange[1])
	if err != nil {
		fmt.Println("Cannot parse int from string:", yRange[1])
	}

	return xMin, xMax, yMin, yMax
}
