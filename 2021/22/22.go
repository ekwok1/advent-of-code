package main

import (
	"fmt"
	"os"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, allData := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	for _, row := range allData {
		fmt.Println(row)
	}
}
