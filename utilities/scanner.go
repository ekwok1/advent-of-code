package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ScanIntsFromFile(filepath string) (*os.File, []int) {
	file, scanner := getScanner(filepath)

	var ints []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Could not parse int from string:", scanner.Text())
		}

		ints = append(ints, int(i))
	}

	return file, ints
}

func ScanIntsFromDelimitedString(filepath string, delimiter string) (*os.File, []int) {
	file, scanner := getScanner(filepath)

	var ints []int
	var stringifiedInts []string

	for scanner.Scan() {
		delimitedString := scanner.Text()
		stringifiedInts = strings.Split(delimitedString, delimiter)
	}

	for _, stringifiedInt := range stringifiedInts {
		fmt.Println(stringifiedInt)
		i, err := strconv.Atoi(stringifiedInt)
		if err != nil {
			fmt.Println("Could not parse int from string:", stringifiedInt)
		}

		ints = append(ints, i)
	}

	return file, ints
}

func ScanStringsFromFile(filepath string) (*os.File, []string) {
	file, scanner := getScanner(filepath)

	var strings []string

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			strings = append(strings, text)
		}
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
