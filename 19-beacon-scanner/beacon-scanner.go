package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, allData := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	beacons := GetBeacons(&allData)

	mappedBeacons := Beacon{}
	for beacon := range beacons[0] {
		mappedBeacons[beacon] = true
	}

	scannerVectors := []Vector3d{{0, 0, 0}}

	beacons = beacons[1:]

	for len(beacons) > 0 {
	OuterLoop:
		for i := len(beacons) - 1; i >= 0; i-- {
			for orientation := 0; orientation < 24; orientation++ {
				diffs := make(map[Vector3d]int)
				for knownBeacon := range mappedBeacons {
					for vector := range beacons[i] {
						diff := vector.Rotate(orientation).Subtract(knownBeacon)
						diffs[diff]++
					}
				}

				for vector, count := range diffs {
					if count >= 12 {
						scannerVector := vector.Invert()
						scannerVectors = append(scannerVectors, scannerVector)

						for vector := range beacons[i] {
							beacon := vector.Rotate(orientation).Add(scannerVector)
							mappedBeacons[beacon] = true
						}

						beacons = append(beacons[:i], beacons[i+1:]...)

						continue OuterLoop
					}
				}
			}
		}
	}

	fmt.Println("Mapped beacons:", len(mappedBeacons))

	maxManhattanDistance := 0
	for i := 0; i < len(scannerVectors); i++ {
		for j := 0; j < len(scannerVectors); j++ {
			manhattanDistance := scannerVectors[i].ManhattanDistance(scannerVectors[j])
			if manhattanDistance > maxManhattanDistance {
				maxManhattanDistance = manhattanDistance
			}
		}
	}

	fmt.Println("Max Manhattan Distance:", maxManhattanDistance)
}

func GetBeacons(allData *[]string) (beacons []Beacon) {
	var beacon Beacon
	for _, row := range *allData {
		if strings.Contains(row, "scanner") {
			beacon = Beacon{}
			beacons = append(beacons, beacon)
			continue
		}

		vector := Vector3d{}
		fmt.Sscanf(row, "%d,%d,%d", &vector.x, &vector.y, &vector.z)
		beacon[vector] = true
	}

	return
}

type Beacon map[Vector3d]bool

type Vector3d struct {
	x, y, z int
}

func (a Vector3d) ManhattanDistance(b Vector3d) int {
	return int(math.Abs(float64(a.x-b.x))) +
		int(math.Abs(float64(a.y-b.y))) +
		int(math.Abs(float64(a.z-b.z)))
}

func (a Vector3d) Add(b Vector3d) Vector3d {
	return Vector3d{
		x: a.x + b.x,
		y: a.y + b.y,
		z: a.z + b.z,
	}
}

func (a Vector3d) Subtract(b Vector3d) Vector3d {
	return Vector3d{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
	}
}

func (v Vector3d) Invert() Vector3d {
	return Vector3d{-v.x, -v.y, -v.z}
}

func (v Vector3d) Rotate(orientation int) Vector3d {
	switch orientation {
	case 0:
		return Vector3d{v.x, v.y, v.z}
	case 1:
		return Vector3d{v.x, -v.z, v.y}
	case 2:
		return Vector3d{v.x, -v.y, -v.z}
	case 3:
		return Vector3d{v.x, v.z, -v.y}
	case 4:
		return Vector3d{-v.x, -v.y, v.z}
	case 5:
		return Vector3d{-v.x, -v.z, -v.y}
	case 6:
		return Vector3d{-v.x, v.y, -v.z}
	case 7:
		return Vector3d{-v.x, v.z, v.y}
	case 8:
		return Vector3d{v.y, v.x, -v.z}
	case 9:
		return Vector3d{v.y, -v.x, v.z}
	case 10:
		return Vector3d{v.y, v.z, v.x}
	case 11:
		return Vector3d{v.y, -v.z, -v.x}
	case 12:
		return Vector3d{-v.y, v.x, v.z}
	case 13:
		return Vector3d{-v.y, -v.x, -v.z}
	case 14:
		return Vector3d{-v.y, -v.z, v.x}
	case 15:
		return Vector3d{-v.y, v.z, -v.x}
	case 16:
		return Vector3d{v.z, v.x, v.y}
	case 17:
		return Vector3d{v.z, -v.x, -v.y}
	case 18:
		return Vector3d{v.z, -v.y, v.x}
	case 19:
		return Vector3d{v.z, v.y, -v.x}
	case 20:
		return Vector3d{-v.z, v.x, -v.y}
	case 21:
		return Vector3d{-v.z, -v.x, v.y}
	case 22:
		return Vector3d{-v.z, v.y, v.x}
	case 23:
		return Vector3d{-v.z, -v.y, -v.x}
	default:
		panic(fmt.Sprintf("Orientation %d does not exist", orientation))
	}
}
