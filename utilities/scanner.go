package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanIntsFromFile(filepath string) (*os.File, []int) {
	file, scanner := getScanner(filepath)

	var ints []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Could not parse int:", scanner.Text())
		}

		ints = append(ints, int(i))
	}

	return file, ints
}

func getScanner(filepath string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Could not open file:", filepath)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return file, scanner
}
