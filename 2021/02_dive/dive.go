package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, commands := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	product := dive(commands)
	fmt.Println("Commands Product:", product)

	aimedProduct := aimedDive(commands)
	fmt.Println("Commands Aimed Product:", aimedProduct)
}

func dive(commands []string) int {
	horizontal := 0
	vertical := 0

	for _, command := range commands {
		direction, magnitude := parseMagnitude(command)

		switch direction {
		case "forward":
			horizontal += magnitude
		case "down":
			vertical += magnitude
		case "up":
			vertical -= magnitude
		}
	}

	return horizontal * vertical
}

func aimedDive(commands []string) int {
	aim := 0
	horizontal := 0
	vertical := 0

	for _, command := range commands {
		direction, magnitude := parseMagnitude(command)

		switch direction {
		case "forward":
			horizontal += magnitude
			vertical += aim * magnitude
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		}
	}

	return horizontal * vertical
}

func parseMagnitude(command string) (string, int) {
	commandParts := strings.Fields(command)

	direction := commandParts[0]
	magnitude, _ := strconv.Atoi(commandParts[1])

	return direction, magnitude
}
