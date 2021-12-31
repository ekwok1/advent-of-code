package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, rebootInstructions := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	reactor := make(Reactor)
	initializeReactorTrivial(&rebootInstructions, &reactor)
	initializationAreaCount := countOn(&reactor)
	fmt.Println("Count on in init area:", initializationAreaCount)

	prisms := make([]Prism, 0)
	pointsOfInterestX, pointsOfInterestY, pointsOfInterestZ := getPointsOfInterest(&rebootInstructions, &prisms)
	compressedPointsMapX, compressedLengthMapX := compress(&pointsOfInterestX)
	compressedPointsMapY, compressedLengthMapY := compress(&pointsOfInterestY)
	compressedPointsMapZ, compressedLengthMapZ := compress(&pointsOfInterestZ)
	reactor2 := initializeReactor(&prisms, &compressedPointsMapX, &compressedPointsMapY, &compressedPointsMapZ)
	allCount := countAllOn(reactor2, &compressedLengthMapX, &compressedLengthMapY, &compressedLengthMapZ)
	fmt.Println("Count on all:", allCount)
}

func countAllOn(reactor *Reactor, compressedLengthMapX, compressedLengthMapY, compressedLengthMapZ *map[int]int) (count int) {
	for prism, value := range *reactor {
		if value {
			on := (*compressedLengthMapX)[prism.x] * (*compressedLengthMapY)[prism.y] * (*compressedLengthMapZ)[prism.z]
			count += on
		}
	}

	return
}

func initializeReactor(prisms *[]Prism, compressedPointsMapX, compressedPointsMapY, compressedPointsMapZ *map[int]int) *Reactor {
	reactor := make(Reactor)

	for _, prism := range *prisms {
		for x := (*compressedPointsMapX)[prism.xMin]; x < (*compressedPointsMapX)[prism.xMax+1]; x++ {
			for y := (*compressedPointsMapY)[prism.yMin]; y < (*compressedPointsMapY)[prism.yMax+1]; y++ {
				for z := (*compressedPointsMapZ)[prism.zMin]; z < (*compressedPointsMapZ)[prism.zMax+1]; z++ {
					reactor[Vector3d{x, y, z}] = prism.on
				}
			}
		}
	}

	return &reactor
}

func compress(pointMap *map[int]bool) (compressedPointsMap map[int]int, compressedLengthMap map[int]int) {
	compressedPoints := make([]int, 0)
	for point := range *pointMap {
		compressedPoints = append(compressedPoints, point)
	}
	sort.Ints(compressedPoints)

	compressedPointsMap, compressedLengthMap = make(map[int]int), make(map[int]int)
	for index, point := range compressedPoints {
		compressedPointsMap[point] = index
		if index < len(compressedPoints)-1 {
			compressedLengthMap[index] = compressedPoints[index+1] - point
		}
	}

	return
}

func getPointsOfInterest(rebootInstructions *[]string, prisms *[]Prism) (map[int]bool, map[int]bool, map[int]bool) {
	pointsOfInterestX := make(map[int]bool)
	pointsOfInterestY := make(map[int]bool)
	pointsOfInterestZ := make(map[int]bool)

	for _, instruction := range *rebootInstructions {
		on, xMin, xMax, yMin, yMax, zMin, zMax := parseInstructions(instruction)
		pointsOfInterestX[xMin] = true
		pointsOfInterestX[xMax+1] = true
		pointsOfInterestY[yMin] = true
		pointsOfInterestY[yMax+1] = true
		pointsOfInterestZ[zMin] = true
		pointsOfInterestZ[zMax+1] = true
		(*prisms) = append((*prisms), Prism{xMin, xMax, yMin, yMax, zMin, zMax, on})
	}

	return pointsOfInterestX, pointsOfInterestY, pointsOfInterestZ
}

type Prism struct {
	xMin, xMax, yMin, yMax, zMin, zMax int
	on                                 bool
}

func countOn(reactor *Reactor) (count int) {
	for _, v := range *reactor {
		if v {
			count++
		}
	}

	return
}

func initializeReactorTrivial(rebootInstructions *[]string, reactor *Reactor) {
	for _, instruction := range *rebootInstructions {
		on, xMin, xMax, yMin, yMax, zMin, zMax := parseInstructions(instruction)

		if xMin < -50 || xMax > 50 || yMin < -50 || yMax > 50 || zMin < -50 || zMax > 50 {
			continue
		}

		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				for z := zMin; z <= zMax; z++ {
					(*reactor)[Vector3d{x, y, z}] = on
				}
			}
		}
	}
}

func parseInstructions(line string) (on bool, xMin, xMax, yMin, yMax, zMin, zMax int) {
	instructions := strings.Fields(line)

	on = instructions[0] == "on"
	fmt.Sscanf(instructions[1], "x=%d..%d,y=%d..%d,z=%d..%d",
		&xMin, &xMax, &yMin, &yMax, &zMin, &zMax)

	return
}

type Reactor map[Vector3d]bool

type Vector3d struct {
	x, y, z int
}
