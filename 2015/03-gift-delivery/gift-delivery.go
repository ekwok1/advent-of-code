package main

import (
	"fmt"

	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, route := scanner.ScanStringFromFile("./input.txt")
	defer file.Close()

	deliveryRoute := DeliveryRoute{route}
	fmt.Println("Santa delivered presents to", deliveryRoute.SoloDeliveryCount(), "houses")
	fmt.Println("Santa and Robo-Santa delivered presents to", deliveryRoute.TagTeamDeliveryCount(), "houses")
}

type Location struct {
	x, y int
}

type DeliveryRoute struct {
	route string
}

func (d *DeliveryRoute) TagTeamDeliveryCount() (houses int) {
	deliveryTracker := make(map[Location]bool)
	santaLocation := Location{0, 0}
	roboSantaLocation := Location{0, 0}
	deliveryTracker[santaLocation] = true
	houses = 1

	for index, instruction := range d.route {
		isSanta := index%2 == 0
		var nextLocation Location
		switch instruction {
		case '>':
			if isSanta {
				nextLocation = Location{santaLocation.x + 1, santaLocation.y}
			} else {
				nextLocation = Location{roboSantaLocation.x + 1, roboSantaLocation.y}
			}
		case 'v':
			if isSanta {
				nextLocation = Location{santaLocation.x, santaLocation.y - 1}
			} else {
				nextLocation = Location{roboSantaLocation.x, roboSantaLocation.y - 1}
			}
		case '<':
			if isSanta {
				nextLocation = Location{santaLocation.x - 1, santaLocation.y}
			} else {
				nextLocation = Location{roboSantaLocation.x - 1, roboSantaLocation.y}
			}
		case '^':
			if isSanta {
				nextLocation = Location{santaLocation.x, santaLocation.y + 1}
			} else {
				nextLocation = Location{roboSantaLocation.x, roboSantaLocation.y + 1}
			}
		default:
			panic(fmt.Sprintf("Could not understand instruction %s", string(instruction)))
		}

		if _, ok := deliveryTracker[nextLocation]; !ok {
			deliveryTracker[nextLocation] = true
			houses++
		}

		if isSanta {
			santaLocation = nextLocation
		} else {
			roboSantaLocation = nextLocation
		}
	}

	return
}

func (d *DeliveryRoute) SoloDeliveryCount() (houses int) {
	deliveryTracker := make(map[Location]bool)
	santaLocation := Location{0, 0}
	deliveryTracker[santaLocation] = true
	houses = 1

	for _, instruction := range d.route {
		var nextLocation Location
		switch instruction {
		case '>':
			nextLocation = Location{santaLocation.x + 1, santaLocation.y}
		case 'v':
			nextLocation = Location{santaLocation.x, santaLocation.y - 1}
		case '<':
			nextLocation = Location{santaLocation.x - 1, santaLocation.y}
		case '^':
			nextLocation = Location{santaLocation.x, santaLocation.y + 1}
		default:
			panic(fmt.Sprintf("Could not understand instruction %s", string(instruction)))
		}

		if _, ok := deliveryTracker[nextLocation]; !ok {
			deliveryTracker[nextLocation] = true
			houses++
		}

		santaLocation = nextLocation
	}

	return
}
