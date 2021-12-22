package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, measurements := utilities.ScanIntsFromFile(os.Args[1])
	defer file.Close()

	increases := countIncreases(measurements)
	fmt.Println("Increases:", increases)

	slidingIncreases := countSlidingIncreases(measurements)
	fmt.Println("Sliding increases:", slidingIncreases)
}

func countIncreases(measurements []int) (increases int) {
	previous := math.MaxInt

	for _, measurement := range measurements {
		if measurement > previous {
			increases++
		}

		previous = measurement
	}

	return
}

func countSlidingIncreases(measurements []int) int {
	var slidingMeasurements []int

	for i := 0; i < len(measurements)-2; i++ {
		sum := measurements[i] + measurements[i+1] + measurements[i+2]
		slidingMeasurements = append(slidingMeasurements, sum)
	}

	return countIncreases(slidingMeasurements)
}
