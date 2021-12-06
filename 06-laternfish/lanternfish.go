package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, fish := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	school := make([]int, 9)
	initialTimers := strings.Split(fish[0], ",")

	for _, initialTime := range initialTimers {
		initialTime, _ := strconv.Atoi(initialTime)
		school[initialTime]++
	}

	simulateDays(&school, 256)
	total := utilities.CountTotal(&school)

	fmt.Println("Total Lanternfish:", total)
}

func simulateDays(school *[]int, days int) {
	for d := 0; d < days; d++ {
		toReset := (*school)[0]
		*school = append((*school)[1:], (*school)[:1]...)
		(*school)[6] += toReset
	}
}
