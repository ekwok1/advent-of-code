package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, positions := utilities.ScanIntsFromDelimitedString(os.Args[1], ",")
	defer file.Close()

	min, max := getPositionRange(&positions)

	basicFuel := calcMinimumFuel(min, max, &positions, calcBasicFuel)
	fmt.Println("Basic Fuel:", basicFuel)

	expensiveFuel := calcMinimumFuel(min, max, &positions, calcExpensiveFuel)
	fmt.Println("Expensive Fuel:", expensiveFuel)
}

func calcMinimumFuel(min int, max int, positions *[]int, calculator func(*[]int, int) int) int {
	minimumFuel := math.MaxInt

	for p := min; p <= max; p++ {
		fuel := calculator(positions, p)

		if fuel < minimumFuel {
			minimumFuel = fuel
		}
	}

	return minimumFuel
}

func calcBasicFuel(positions *[]int, position int) (fuel int) {
	for _, p := range *positions {
		fuel += int(math.Abs(float64(p - position)))
	}

	return
}

func calcExpensiveFuel(positions *[]int, position int) (fuel int) {
	fuelMap := make(map[int]int)

	for _, p := range *positions {
		distance := int(math.Abs(float64(p - position)))

		if val, ok := fuelMap[distance]; ok {
			fuel += val
		} else {
			var fuelNeeded int
			for i := 1; i <= distance; i++ {
				fuelNeeded += i
			}
			fuel += fuelNeeded
			fuelMap[distance] = fuelNeeded
		}
	}

	return
}

func getPositionRange(positions *[]int) (min int, max int) {
	for _, position := range *positions {
		if position < min {
			min = position
		} else if position > max {
			max = position
		}
	}

	return
}
