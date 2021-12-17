package main

import (
	"fmt"
	"strconv"
)

func main() {
	rawHex := "60A100"
	i, err := strconv.ParseUint(rawHex, 16, 32)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%024b\n", i)

	fmt.Println(asBits(i))

}

func asBits(val uint64) []uint64 {
	bits := []uint64{}
	for i := 0; i < 24; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		val = val >> 1
	}
	return bits
}

type Packet struct {
	version            int
	typeId             int
	literalValue       int
	lengthTypeId       int
	subPacketLength    int
	numberOfSubpackets int
	subpackets         []Packet
}
