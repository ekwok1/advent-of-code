package main

func main() {
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
}
