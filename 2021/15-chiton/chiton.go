package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ekwok1/aoc-2021/utilities"
	"github.com/ekwok1/aoc-2021/utilities/scanner"
)

func main() {
	file, riskData := scanner.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	riskGrid := initialRiskGrid(&riskData)

	minimumTotalRisk := getMinimumTotalRisk(&riskGrid)
	fmt.Println("Minimum Risk:", minimumTotalRisk)

	minimumTotalRiskFullCave := getMinimumTotalRiskFullCave(&riskGrid)
	fmt.Println("Minimum Risk Full Cave:", minimumTotalRiskFullCave)
}

func getMinimumTotalRiskFullCave(riskGrid *[][]int) int {
	gridSize := len(*riskGrid)
	fullSize := gridSize * 5
	costGrid := utilities.CreateIntGrid(fullSize)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Location{risk: 0, coordinates: Coordinates{x: 0, y: 0}})

	visited := make(map[Coordinates]bool)

	for len(pq) > 0 {
		location := heap.Pop(&pq).(*Location)
		risk := location.risk
		coordinates := location.coordinates
		row := coordinates.y
		col := coordinates.x

		if visited[coordinates] {
			continue
		}
		visited[coordinates] = true

		costGrid[row][col] = risk

		if row == fullSize-1 && col == fullSize-1 {
			break
		}

		rowDir := [4]int{-1, 0, 0, 1}
		colDir := [4]int{0, -1, 1, 0}

		for dir := 0; dir < 4; dir++ {
			newRow := row + rowDir[dir]
			newCol := col + colDir[dir]
			if !isFullCaveSafe(riskGrid, newRow, newCol) {
				continue
			}

			heap.Push(&pq, &Location{risk: risk + getFullCaveCost(riskGrid, newRow, newCol), coordinates: Coordinates{x: newCol, y: newRow}})
		}
	}

	return costGrid[fullSize-1][fullSize-1]
}

func isFullCaveSafe(riskmap *[][]int, row int, column int) bool {
	totalRows := len(*riskmap) * 5
	totalColumns := len((*riskmap)[0]) * 5
	return row >= 0 && column >= 0 && row < totalRows && column < totalColumns
}

func getFullCaveCost(riskGrid *[][]int, row int, col int) int {
	gridSize := len(*riskGrid)
	basicCost := (*riskGrid)[row%gridSize][col%gridSize]
	incrementedCost := basicCost + (row / gridSize) + (col / gridSize)

	return (incrementedCost-1)%9 + 1
}

func getMinimumTotalRisk(riskGrid *[][]int) int {
	gridSize := len(*riskGrid)
	costGrid := utilities.CreateIntGrid(gridSize)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Location{risk: 0, coordinates: Coordinates{x: 0, y: 0}})

	visited := make(map[Coordinates]bool)

	for len(pq) > 0 {
		location := heap.Pop(&pq).(*Location)
		risk := location.risk
		coordinates := location.coordinates
		row := coordinates.y
		col := coordinates.x

		if visited[coordinates] {
			continue
		}
		visited[coordinates] = true

		costGrid[row][col] = risk

		if row == gridSize-1 && col == gridSize-1 {
			break
		}

		rowDir := [4]int{-1, 0, 0, 1}
		colDir := [4]int{0, -1, 1, 0}

		for dir := 0; dir < 4; dir++ {
			newRow := row + rowDir[dir]
			newCol := col + colDir[dir]
			if !isSafe(riskGrid, newRow, newCol) {
				continue
			}

			heap.Push(&pq, &Location{risk: risk + (*riskGrid)[newRow][newCol], coordinates: Coordinates{x: newCol, y: newRow}})
		}
	}

	return costGrid[gridSize-1][gridSize-1]
}

func isSafe(riskmap *[][]int, row int, column int) bool {
	totalRows := len(*riskmap)
	totalColumns := len((*riskmap)[0])
	return row >= 0 && column >= 0 && row < totalRows && column < totalColumns
}

func initialRiskGrid(initialEnergyGrid *[]string) (grid [][]int) {
	for _, row := range *initialEnergyGrid {
		initialRiskRow := strings.Split(row, "")

		var intRow []int
		for _, risk := range initialRiskRow {
			risk, _ := strconv.Atoi(risk)
			intRow = append(intRow, risk)
		}

		grid = append(grid, intRow)
	}

	return
}

type Location struct {
	risk        int
	coordinates Coordinates
}

type Coordinates struct {
	x, y int
}

type PriorityQueue []*Location

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].risk < pq[j].risk
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Location)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
