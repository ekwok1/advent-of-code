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

func ScanStringsFromFile(filepath string) (*os.File, []string) {
	file, scanner := getScanner(filepath)

	var strings []string

	for scanner.Scan() {
		s := scanner.Text()

		strings = append(strings, s)
	}

	return file, strings
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
