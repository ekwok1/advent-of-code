package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanIntsFromFile(filepath string) (file *os.File, ints []int) {
	file, scanner := getScanner(filepath)

	for scanner.Scan() {
		var i int
		fmt.Sscanf(scanner.Text(), "%d", &i)
		ints = append(ints, i)
	}

	return
}

func ScanIntsFromDelimitedString(filepath string, delimiter string) (file *os.File, ints []int) {
	file, scanner := getScanner(filepath)

	var stringifiedInts []string

	for scanner.Scan() {
		delimitedString := scanner.Text()
		stringifiedInts = strings.Split(delimitedString, delimiter)
	}

	for _, stringifiedInt := range stringifiedInts {
		var i int
		fmt.Sscanf(stringifiedInt, "%d", &i)
		ints = append(ints, i)
	}

	return
}

func ScanStringFromFile(filepath string) (file *os.File, str string) {
	file, scanner := getScanner(filepath)

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			str = text
		}
	}

	return
}

func ScanStringsFromFile(filepath string) (file *os.File, strings []string) {
	file, scanner := getScanner(filepath)

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			strings = append(strings, text)
		}
	}

	return
}

func getScanner(filepath string) (file *os.File, scanner *bufio.Scanner) {
	file, _ = os.Open(filepath)

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return
}
