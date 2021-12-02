package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, commands := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	product := dive(commands)
	aimedProduct := aimedDive(commands)

	fmt.Println("Commands Product:", product)
	fmt.Println("Commands Aimed Product:", aimedProduct)
}

func dive(commands []string) int {
	horizontal := 0
	vertical := 0

	for _, command := range commands {
		commandParts := strings.Fields(command)

		i, err := strconv.Atoi(commandParts[1])
		if err != nil {
			fmt.Println("Cannot convert to int:", commandParts[1])
		}

		switch commandParts[0] {
		case "forward":
			horizontal += i
		case "down":
			vertical += i
		case "up":
			vertical -= i
		}
	}

	return horizontal * vertical
}

func aimedDive(commands []string) int {
	aim := 0
	horizontal := 0
	vertical := 0

	for _, command := range commands {
		commandParts := strings.Fields(command)

		i, err := strconv.Atoi(commandParts[1])
		if err != nil {
			fmt.Println("Cannot convert to int:", commandParts[1])
		}

		switch commandParts[0] {
		case "forward":
			horizontal += i
			vertical += aim * i
		case "down":
			aim += i
		case "up":
			aim -= i
		}
	}

	return horizontal * vertical
}
