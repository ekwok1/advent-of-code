package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, binaries := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	powerConsumption := calcPowerConsumption(binaries)
	oxygenGeneratorRating := calcOxygenGeneratorRating(binaries)
	carbonDioxideScrubberRating := calcCarbonDioxideScrubberRating(binaries)
	lifeSupportRating := oxygenGeneratorRating * carbonDioxideScrubberRating

	fmt.Println("Power Consumption:", powerConsumption)
	fmt.Println("Oxygen Generator Rating:", oxygenGeneratorRating)
	fmt.Println("Carbon Dioxide Scrubber Rating:", carbonDioxideScrubberRating)
	fmt.Println("Life Support Rating:", lifeSupportRating)
}

// REFACTOR
// Method for getting most common based on index?
// Is there a way to do this in place without looping so many times?

func calcPowerConsumption(binaries []string) int {
	gammaBinary := ""
	epsilonBinary := ""

	binaryLength := len(binaries[0])

	for i := 0; i < binaryLength; i++ {
		ones := 0
		zeros := 0

		for _, binary := range binaries {
			if binary[i] == 48 {
				zeros++
			} else {
				ones++
			}
		}

		if ones > zeros {
			gammaBinary += "1"
			epsilonBinary += "0"
		} else {
			gammaBinary += "0"
			epsilonBinary += "1"
		}
	}

	gamma, _ := strconv.ParseInt(gammaBinary, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonBinary, 2, 64)
	return int(gamma * epsilon)
}

func calcOxygenGeneratorRating(binaries []string) int {
	oxygenGeneratorRatingBinary := ""

	binaryLength := len(binaries[0])
	binariesCopy := binaries

	for i := 0; i < binaryLength; i++ {
		ones := 0
		zeros := 0

		for _, binary := range binariesCopy {
			if binary[i] == 48 {
				zeros++
			} else {
				ones++
			}
		}

		temp := binariesCopy
		binariesCopy = nil
		if ones >= zeros {
			for _, binary := range temp {
				if binary[i] == 49 {
					binariesCopy = append(binariesCopy, binary)
				}
			}
		} else {
			for _, binary := range temp {
				if binary[i] == 48 {
					binariesCopy = append(binariesCopy, binary)
				}
			}
		}

		if len(binariesCopy) == 1 {
			oxygenGeneratorRatingBinary = binariesCopy[0]
		}
	}

	oxygenGeneratorRating, _ := strconv.ParseInt(oxygenGeneratorRatingBinary, 2, 64)
	return int(oxygenGeneratorRating)
}

func calcCarbonDioxideScrubberRating(binaries []string) int {
	carbonDioxideScrubberRatingBinary := ""

	binaryLength := len(binaries[0])
	binariesCopy := binaries

	for i := 0; i < binaryLength; i++ {
		ones := 0
		zeros := 0

		for _, binary := range binariesCopy {
			if binary[i] == 48 {
				zeros++
			} else {
				ones++
			}
		}

		temp := binariesCopy
		binariesCopy = nil
		if ones < zeros {
			for _, binary := range temp {
				if binary[i] == 49 {
					binariesCopy = append(binariesCopy, binary)
				}
			}
		} else {
			for _, binary := range temp {
				if binary[i] == 48 {
					binariesCopy = append(binariesCopy, binary)
				}
			}
		}

		if len(binariesCopy) == 1 {
			carbonDioxideScrubberRatingBinary = binariesCopy[0]
		}
	}

	carbonDioxideScrubberRating, _ := strconv.ParseInt(carbonDioxideScrubberRatingBinary, 2, 64)
	return int(carbonDioxideScrubberRating)
}
