package main

import (
	"fmt"
	"os"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, initialTimes := utilities.ScanIntsFromDelimitedString(os.Args[1], ",")
	defer file.Close()

	school := setupInitialState(initialTimes)
	simulateDays(&school, 256)
	sum := utilities.Sum(&school)

	fmt.Println("Total Lanternfish:", sum)
}

func simulateDays(school *[]int, days int) {
	for d := 0; d < days; d++ {
		toReset := (*school)[0]
		*school = append((*school)[1:], (*school)[:1]...)
		(*school)[6] += toReset
	}
}

func setupInitialState(initialTimes []int) (school []int) {
	school = make([]int, 9)

	for _, time := range initialTimes {
		school[time]++
	}

	return
}
