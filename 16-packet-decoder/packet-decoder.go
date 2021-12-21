package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, hexadecimal := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	hexadecimalMap := initializeHexadecimalMap()
	binary := convertHexdecimalToBinary(hexadecimal[0], &hexadecimalMap)

	versionCounter := parsePacketVersions(binary, -1)
	fmt.Println("Version counter:", versionCounter)

	value, _ := parseAndOperate(binary, 0)
	fmt.Println("Value of operations:", value)
}

func parseAndOperate(packet string, index int) (int, int) {
	typeId, _ := strconv.ParseInt(packet[index+3:index+6], 2, 64)

	if typeId == 4 {
		index += 6
		value := 0

		for {
			number, _ := strconv.ParseInt(packet[index+1:index+5], 2, 64)
			value = value*16 + int(number)
			index += 5

			if packet[index-5] == '0' {
				return value, index
			}
		}
	} else {
		lengthTypeId := packet[index+6]
		values := make([]int, 0)

		if lengthTypeId == '0' {
			bitLength, _ := strconv.ParseInt(packet[index+7:index+22], 2, 64)
			startIndex := index + 22
			index = startIndex

			for {
				value, nextIndex := parseAndOperate(packet, index)
				values = append(values, value)
				index = nextIndex
				if nextIndex-startIndex == int(bitLength) {
					break
				}
			}
		} else {
			numberOfPackets, _ := strconv.ParseInt(packet[index+7:index+18], 2, 64)
			index += 18

			for n := 0; n < int(numberOfPackets); n++ {
				value, nextIndex := parseAndOperate(packet, index)
				values = append(values, value)
				index = nextIndex
			}
		}

		switch typeId {
		case 0:
			sum := 0
			for _, value := range values {
				sum += value
			}
			return sum, index
		case 1:
			product := 1
			for _, value := range values {
				product *= value
			}
			return product, index
		case 2:
			min := math.MaxInt
			for _, value := range values {
				if value < min {
					min = value
				}
			}
			return min, index
		case 3:
			max := 0
			for _, value := range values {
				if value > max {
					max = value
				}
			}
			return max, index
		case 5:
			if values[0] > values[1] {
				return 1, index
			} else {
				return 0, index
			}
		case 6:
			if values[0] < values[1] {
				return 1, index
			} else {
				return 0, index
			}
		case 7:
			if values[0] == values[1] {
				return 1, index
			} else {
				return 0, index
			}
		}
	}

	return 0, index
}

func parsePacketVersions(packet string, counter int) int {
	if packet == "" {
		return 0
	}

	intPacket, _ := strconv.ParseInt(packet, 2, 64)
	if intPacket == 0 {
		return 0
	}

	if counter == 0 {
		return parsePacketVersions(packet, -1)
	}

	version, _ := strconv.ParseInt(packet[:3], 2, 64)
	typeId, _ := strconv.ParseInt(packet[3:6], 2, 64)

	if typeId == 4 {
		var bLiteral string
		i := 6
		for {
			goNext := string(packet[i])
			next := packet[i+1 : i+5]
			i += 5
			bLiteral += next
			if goNext != "1" {
				break
			}
		}

		return int(version) + parsePacketVersions(packet[i:], counter-1)
	}

	lengthTypeId, _ := strconv.Atoi(string(packet[6]))
	if lengthTypeId == 0 {
		bitLength, _ := strconv.ParseInt(packet[7:22], 2, 64)
		return int(version) + parsePacketVersions(packet[22:22+bitLength], -1) + parsePacketVersions(packet[22+bitLength:], counter-1)
	} else {
		numberOfSubpackets, _ := strconv.ParseInt(packet[7:18], 2, 64)
		return int(version) + parsePacketVersions(packet[18:], int(numberOfSubpackets))
	}
}

func convertHexdecimalToBinary(hexadecimal string, hexadecimalMap *map[string]string) (binary string) {
	for i := 0; i < len(hexadecimal); i++ {
		binary += (*hexadecimalMap)[string(hexadecimal[i])]
	}

	return
}

func initializeHexadecimalMap() map[string]string {
	hexadecimalMap := make(map[string]string)

	hexadecimalMap["0"] = "0000"
	hexadecimalMap["1"] = "0001"
	hexadecimalMap["2"] = "0010"
	hexadecimalMap["3"] = "0011"
	hexadecimalMap["4"] = "0100"
	hexadecimalMap["5"] = "0101"
	hexadecimalMap["6"] = "0110"
	hexadecimalMap["7"] = "0111"
	hexadecimalMap["8"] = "1000"
	hexadecimalMap["9"] = "1001"
	hexadecimalMap["A"] = "1010"
	hexadecimalMap["B"] = "1011"
	hexadecimalMap["C"] = "1100"
	hexadecimalMap["D"] = "1101"
	hexadecimalMap["E"] = "1110"
	hexadecimalMap["F"] = "1111"

	return hexadecimalMap
}
