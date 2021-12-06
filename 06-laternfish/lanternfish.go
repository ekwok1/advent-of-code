package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, timers := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	school := setupInitialState(timers[0])
	simulateDays(&school, 256)
	total := utilities.CountTotal(&school)

	fmt.Println("Total Lanternfish:", total)
}

func setupInitialState(timers string) (school []int) {
	school = make([]int, 9)
	initialState := strings.Split(timers, ",")

	for _, time := range initialState {
		time, _ := strconv.Atoi(time)
		school[time]++
	}

	return
}

func simulateDays(school *[]int, days int) {
	for d := 0; d < days; d++ {
		toReset := (*school)[0]
		*school = append((*school)[1:], (*school)[:1]...)
		(*school)[6] += toReset
	}
}
