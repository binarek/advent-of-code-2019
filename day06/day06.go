package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const com = "COM"

func main() {
	relations := readOrbitRelations()
	fmt.Println("Part 1:", SolvePart1(relations))
}

type spaceObject struct {
	center *spaceObject
	orbits []*spaceObject
	name   string
}

func readOrbitRelations() map[string][]string { // key: center, value: objects in orbit
	file, _ := os.Open(func() string {
		if len(os.Args) >= 2 {
			return os.Args[1]
		}
		return "./day06_input"
	}())
	scanner := bufio.NewScanner(file)
	relations := make(map[string][]string)
	for scanner.Scan() {
		objects := strings.Split(scanner.Text(), ")")
		if len(objects) != 2 {
			fmt.Println("Warn: invalid file line")
			continue
		}
		relations[objects[0]] = append(relations[objects[0]], objects[1])
	}
	return relations
}

func SolvePart1(orbitRelations map[string][]string) int {
	comNode := spaceObject{nil, nil, com}
	buildTree(&comNode, orbitRelations)
	return countAllOrbits(&comNode)
}

func buildTree(node *spaceObject, orbitRelations map[string][]string) {
	for _, orbitName := range orbitRelations[node.name] {
		orbitNode := spaceObject{node, nil, orbitName}
		node.orbits = append(node.orbits, &orbitNode)
		buildTree(&orbitNode, orbitRelations)
	}
}

func countAllOrbits(node *spaceObject) int {
	counter := countOrbits(node)
	for _, nodeOrbit := range node.orbits {
		counter += countAllOrbits(nodeOrbit)
	}
	return counter
}

func countOrbits(node *spaceObject) int {
	if node.center != nil {
		return 1 + countOrbits(node.center)
	}
	return 0
}
