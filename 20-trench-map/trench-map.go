package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, allData := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	inputImage := allData[1:]
	image := map[Coordinate]bool{}
	for row := 0; row < len(inputImage); row++ {
		for col := 0; col < len(inputImage[0]); col++ {
			image[Coordinate{col, row}] = (inputImage[row][col] == '#')
		}
	}

	minX, maxX, minY, maxY := getImageBounds(image)
	minX, maxX, minY, maxY = addImageBuffer(image, minX, maxX, minY, maxY, false)

	algorithm := allData[0]
	enhancements := 50
	for i := 0; i < enhancements; i++ {
		defaultValue := image[Coordinate{minX, maxY}]
		minX, maxX, minY, maxY = addImageBuffer(image, minX, maxX, minY, maxY, defaultValue)
		image = enhance(image, algorithm, defaultValue)
	}

	litPixels := countLitPixels(image)
	fmt.Println("Lit pixels:", litPixels)
}

func countLitPixels(image Image) (count int) {
	for _, lit := range image {
		if lit {
			count++
		}
	}

	return
}

func enhance(image map[Coordinate]bool, algorithm string, defaultValue bool) Image {
	enhancedImage := Image{}

	for key := range image {
		neighbors := []Coordinate{
			{key.x - 1, key.y - 1},
			{key.x, key.y - 1},
			{key.x + 1, key.y - 1},
			{key.x - 1, key.y},
			key,
			{key.x + 1, key.y},
			{key.x - 1, key.y + 1},
			{key.x, key.y + 1},
			{key.x + 1, key.y + 1},
		}

		index := 0
		for _, neighbor := range neighbors {
			value, ok := image[neighbor]
			index = index << 1
			if ok && value {
				index++
			} else if !ok && defaultValue {
				index++
			}
		}

		enhancedImage[key] = (algorithm[index] == '#')
	}

	return enhancedImage
}

func addImageBuffer(image Image, minX, maxX, minY, maxY int, value bool) (newMinX, newMaxX, newMinY, newMaxY int) {
	newMinX = minX - 1
	newMaxX = maxX + 1
	for x := newMinX; x <= newMaxX; x++ {
		image[Coordinate{x, minY - 1}] = value
		image[Coordinate{x, maxY + 1}] = value
	}

	newMinY = minY - 1
	newMaxY = maxY + 1
	for y := newMinY; y <= newMaxY; y++ {
		image[Coordinate{minX - 1, y}] = value
		image[Coordinate{maxX + 1, y}] = value
	}

	return newMinX, newMaxX, newMinY, newMaxY
}

func getImageBounds(image Image) (minX, maxX, minY, maxY int) {
	minX = math.MaxInt
	maxX = math.MinInt
	minY = math.MaxInt
	maxY = math.MinInt

	for key := range image {
		if key.x < minX {
			minX = key.x
		}
		if key.x > maxX {
			maxX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
		if key.y > maxY {
			maxY = key.y
		}
	}

	return
}

type Coordinate struct {
	x, y int
}

type Image map[Coordinate]bool
