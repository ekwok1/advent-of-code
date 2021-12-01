package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var measurements []int

	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		measurements = append(measurements, int(i))
	}

	var slidingMeasurements []int

	for i := 0; i < len(measurements)-2; i++ {
		sum := measurements[i] + measurements[i+1] + measurements[i+2]
		slidingMeasurements = append(slidingMeasurements, sum)
	}

	previous := math.MaxInt
	increases := 0
	for _, slidingMeasurement := range slidingMeasurements {
		if slidingMeasurement > previous {
			increases = increases + 1
		}

		previous = int(slidingMeasurement)
	}

	fmt.Println(increases)
}
