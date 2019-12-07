package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Constants
const SEP = ")"

// Structs
type spaceObject struct {
	attracts []*spaceObject
	visited bool
}

type spaceObjectWithLookup struct {
	graph *spaceObject
	matching map[string]*spaceObject
}

// It is an undirected acyclic graph
func linkSpaceObjects(from, to *spaceObject) {
	from.attracts = append(from.attracts, to)
	to.attracts = append(to.attracts, from)
}

// Functions
func assert (got, expected int) {
	if got != expected {
		panic(fmt.Sprintf("ERROR: got %d, expected %d!\n", got, expected))
	}
}

func strToSpaceObject(matching map[string]*spaceObject, s string) *spaceObject {
	object, exists := matching[s]
	// Return the existing object
	if exists {
		return object
	}
	// Create a new one
	newObject := &spaceObject{}
	matching[s] = newObject
	return newObject
}

func findNodeExcept(matching map[string]*spaceObject, s string) *spaceObject {
	object, exists := matching[s]
	// Return the existing object
	if !exists {
		panic(fmt.Sprintf("Object %s does not exist in map!\n", s))
	}
	return object
}

func fileToGraph(filename, rootID string) spaceObjectWithLookup {
	m := make(map[string]*spaceObject)
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), SEP)
		linkSpaceObjects(strToSpaceObject(m, names[0]), strToSpaceObject(m, names[1]))
	}
	return spaceObjectWithLookup{m[rootID], m}
}

func countOrbitsTail(node *spaceObject, from *spaceObject, counter int) int {
	result := counter + 1
	childArg := result
	for _, el := range node.attracts {
		if el != from {
			result += countOrbitsTail(el, node, childArg)
		}
	}
	return result
}

func countOrbits(root *spaceObject) int {
	result := 0
	for _, el := range root.attracts {
		result += countOrbitsTail(el, root, 0)
	}
	return result
}

func bfsSearch(node, target *spaceObject) int {
	dist := 0
	var level []*spaceObject; level = append(level, node)
	var queue [][]*spaceObject; queue = append(queue, level)

	for iLevel := 0; iLevel < len(queue); iLevel++ {
		var newLevel []*spaceObject
		for _, n := range queue[iLevel] {
			n.visited = true
			if n == target {
				return dist
			}
			for _, child := range n.attracts {
				if !child.visited {
					newLevel = append(newLevel, child)
				}
			}
		}
		queue = append(queue, newLevel)
		dist++
	}
	return 100000000000
}

// It considers the objects the src and dst are attached to
func findShortestPath(graphMap spaceObjectWithLookup, fromStr string, toStr string) int {
	return bfsSearch(
		findNodeExcept(graphMap.matching, fromStr).attracts[0],
		findNodeExcept(graphMap.matching, toStr).attracts[0])
}

func main() {
	// Tests
	assert(countOrbits(fileToGraph("test1.txt", "COM").graph), 42)
	assert(findShortestPath(fileToGraph("test2.txt", "COM"), "YOU", "SAN"), 4)
	
	// Problem
	graphMap := fileToGraph("input.txt", "COM")
	fmt.Printf("Part 1: %d\n", countOrbits(graphMap.graph))
	fmt.Printf("Part 2: %d\n", findShortestPath(graphMap, "YOU", "SAN"))
}