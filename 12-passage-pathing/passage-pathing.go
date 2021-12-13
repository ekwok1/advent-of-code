package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/ekwok1/aoc-2021/utilities"
)

func main() {
	file, connections := utilities.ScanStringsFromFile(os.Args[1])
	defer file.Close()

	caveSystem := initializeCaveSystem(&connections)

	start := caveSystem.caves["start"]
	end := caveSystem.caves["end"]

	pathsNoReturn := countCavePathsNoReturn(caveSystem, start, end, start)
	fmt.Println("Paths without returning to small caves:", pathsNoReturn)

	pathsOneReturn := countCavePathsOneReturn(caveSystem, start, end, start)
	fmt.Println("Paths allowing one return to a small cave:", pathsOneReturn)
}

func countCavePathsOneReturn(caveSystem *CaveSystem, start *Cave, end *Cave, currentCave *Cave) int {
	visited := make(map[*Cave]int)

	paths := 0
	paths = caveDFSOneReturn(caveSystem, start, end, currentCave, &visited, &paths)

	return paths
}

func caveDFSOneReturn(caveSystem *CaveSystem, start *Cave, end *Cave, currentCave *Cave, visited *map[*Cave]int, paths *int) int {
	if currentCave.name == end.name {
		(*paths)++
		return (*paths)
	}

	if isSmall(currentCave) {
		(*visited)[currentCave]++

		multipleVisits := 0
		for cave := range *visited {
			if (*visited)[cave] > 1 {
				multipleVisits++
			}

			if (*visited)[cave] > 2 {
				(*visited)[cave]--
				return (*paths)
			}
		}

		if multipleVisits > 1 {
			(*visited)[currentCave]--
			return (*paths)
		}
	}

	for connection := range currentCave.connections {
		if connection == start.name {
			continue
		}
		caveDFSOneReturn(caveSystem, start, end, caveSystem.caves[connection], visited, paths)
	}

	if isSmall(currentCave) {
		(*visited)[currentCave]--
	}

	return (*paths)
}

func countCavePathsNoReturn(caveSystem *CaveSystem, start *Cave, end *Cave, currentCave *Cave) int {
	visited := make(map[*Cave]bool)

	paths := 0
	paths = caveDFSNoReturn(caveSystem, start, end, currentCave, &visited, &paths)

	return paths
}

func caveDFSNoReturn(caveSystem *CaveSystem, start *Cave, end *Cave, currentCave *Cave, visited *map[*Cave]bool, paths *int) int {
	if currentCave.name == end.name {
		(*paths)++
		return (*paths)
	}

	if isSmall(currentCave) && (*visited)[currentCave] {
		return (*paths)
	}

	if isSmall(currentCave) {
		(*visited)[currentCave] = true
	}

	for connection := range currentCave.connections {
		if connection == start.name {
			continue
		}
		caveDFSNoReturn(caveSystem, start, end, caveSystem.caves[connection], visited, paths)
	}

	if isSmall(currentCave) {
		(*visited)[currentCave] = false
	}

	return *paths
}

func isSmall(cave *Cave) bool {
	r := rune(cave.name[0])
	return unicode.IsLower(r)
}

func initializeCaveSystem(connections *[]string) *CaveSystem {
	caveSystem := NewCaveSystem()

	for _, connection := range *connections {
		caves := strings.Split(connection, "-")
		caveSystem.AddCave(caves[0])
		caveSystem.AddCave(caves[1])
		caveSystem.AddConnection(caves[0], caves[1])
	}

	return caveSystem
}

type Cave struct {
	name        string
	connections map[string]*Cave
}

func NewCave(name string) *Cave {
	return &Cave{
		name:        name,
		connections: map[string]*Cave{},
	}
}

type CaveSystem struct {
	caves map[string]*Cave
}

func NewCaveSystem() *CaveSystem {
	return &CaveSystem{
		caves: map[string]*Cave{},
	}
}

func (cs *CaveSystem) AddCave(name string) {
	if _, ok := cs.caves[name]; ok {
		return
	}

	cave := NewCave(name)
	cs.caves[name] = cave
}

func (cs *CaveSystem) AddConnection(name1, name2 string) {
	cave1 := cs.caves[name1]
	cave2 := cs.caves[name2]

	if cave1 == nil || cave2 == nil {
		panic("Not all caves exist")
	}

	if _, ok := cave1.connections[cave2.name]; ok {
		return
	}

	cave1.connections[cave2.name] = cave2
	cave2.connections[cave1.name] = cave1

	cs.caves[cave1.name] = cave1
	cs.caves[cave2.name] = cave2
}
