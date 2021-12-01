package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("test-input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var measurements []int64

	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		measurements = append(measurements, i)
	}

	previous := math.MaxInt
	increases := 0
	for _, measurement := range measurements {
		if int(measurement) > previous {
			increases = increases + 1
		}

		previous = int(measurement)
	}

	fmt.Println(increases)
}
