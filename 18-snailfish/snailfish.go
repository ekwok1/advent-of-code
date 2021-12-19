package main

import (
	"fmt"
	"os"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, snailfishNumbers := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	for i, snailfishNumber := range snailfishNumbers {
		if i != 0 {
			continue
		}
		fmt.Println(snailfishNumber)
	}
}
