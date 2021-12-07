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
	fmt.Println("Power Consumption:", powerConsumption)

	oxygenGeneratorRating := getOxygenGeneratorRating(binaries, 0)
	fmt.Println("Oxygen Generator Rating:", oxygenGeneratorRating)

	carbonDioxideScrubberRating := getCarbonDioxideScrubberRating(binaries, 0)
	fmt.Println("Carbon Dioxide Scrubber Rating:", carbonDioxideScrubberRating)

	lifeSupportRating := oxygenGeneratorRating * carbonDioxideScrubberRating
	fmt.Println("Life Support Rating:", lifeSupportRating)
}

func calcPowerConsumption(binaries []string) int {
	binaryMap := make(map[string]int)
	for _, binary := range binaries {
		for i, bit := range binary {
			if bit == '0' {
				binaryMap[fmt.Sprintf("zero%d", i)]++
			} else {
				binaryMap[fmt.Sprintf("one%d", i)]++
			}
		}
	}

	gammaBinary := ""
	epsilonBinary := ""

	binaryLength := len(binaries[0])
	for i := 0; i < binaryLength; i++ {
		zeroes := binaryMap[fmt.Sprintf("zero%d", i)]
		ones := binaryMap[fmt.Sprintf("one%d", i)]

		if ones > zeroes {
			gammaBinary += "1"
			epsilonBinary += "0"
		} else {
			gammaBinary += "0"
			epsilonBinary += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaBinary, 2, 64)
	if err != nil {
		fmt.Println("Cannot parse int from binary:", gammaBinary)
	}

	epsilon, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		fmt.Println("Cannot parse int from binary:", epsilonBinary)
	}

	return int(gamma * epsilon)
}

func getOxygenGeneratorRating(binaries []string, index int) int {
	if len(binaries) == 1 {
		oxygenGeneratorRating, err := strconv.ParseInt(binaries[0], 2, 64)
		if err != nil {
			fmt.Println("Cannot parse int from binary:", binaries[0])
		}
		return int(oxygenGeneratorRating)
	}

	mostCommonBit := getMostCommonBit(&binaries, index)
	binaries = getRelevantBinaries(binaries, index, mostCommonBit)
	index++

	return getOxygenGeneratorRating(binaries, index)
}

func getCarbonDioxideScrubberRating(binaries []string, index int) int {
	if len(binaries) == 1 {
		carbonDioxideScrubberRating, err := strconv.ParseInt(binaries[0], 2, 64)
		if err != nil {
			fmt.Println("Cannot parse int from binary:", binaries[0])
		}
		return int(carbonDioxideScrubberRating)
	}

	leastCommonBit := getLeastCommonBit(&binaries, index)
	binaries = getRelevantBinaries(binaries, index, leastCommonBit)
	index++

	return getCarbonDioxideScrubberRating(binaries, index)
}

func getRelevantBinaries(binaries []string, index int, bit byte) (relevantBinaries []string) {
	for _, binary := range binaries {
		if binary[index] == bit {
			relevantBinaries = append(relevantBinaries, binary)
		}
	}
	return
}

func getMostCommonBit(binaries *[]string, index int) byte {
	bitMap := generateBitMap(binaries, index)

	if bitMap['1'] >= bitMap['0'] {
		return byte('1')
	} else {
		return byte('0')
	}
}

func getLeastCommonBit(binaries *[]string, index int) byte {
	bitMap := generateBitMap(binaries, index)

	if bitMap['1'] < bitMap['0'] {
		return byte('1')
	} else {
		return byte('0')
	}
}

func generateBitMap(binaries *[]string, index int) map[rune]int {
	bitMap := make(map[rune]int)

	for _, binary := range *binaries {
		if binary[index] == '1' {
			bitMap['1']++
		} else {
			bitMap['0']++
		}
	}

	return bitMap
}
