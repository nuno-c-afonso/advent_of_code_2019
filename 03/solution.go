package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"os"
)

func assert(got int, expected int) {
	if got != expected {
		panic(fmt.Sprintf("ERROR: %d != %d", got, expected))
	}
}

// Helpers
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Point
type point struct {
	x, y int
}

func maxInLine(l, r point) point {
	if l.x > r.x || l.y > r.y {
		return l
	}
	return r
}

func minInLine(l, r point) point {
	max := maxInLine(l, r)
	if max == l {
		return r
	}
	return l
}

// Line
type line struct {
	start, end point
	isReversed bool
}

func newLineCoords(xS, yS, xE, yE int) line {
	start, end := point{xS, yS}, point{xE, yE}
	return newLinePoints(start, end)
}

func newLinePoints(start, end point) line {
	i, j := minInLine(start, end), maxInLine(start, end)
	return line{i, j, i != start}
}

func lineLength(l line) int {
	p := point{l.end.x - l.start.x, l.end.y - l.start.y}
	return p.x + p.y
}

// Wire
type wire struct {
	segments []line
}

func manhattanDistance(intersection point) int {
	return abs(intersection.x) + abs(intersection.y)
}

func lineVector(l line) point {
	return point{l.end.x - l.start.x, l.end.y - l.start.y}
}

func mulLineVectors(l, r line) point {
	lVector, rVector := lineVector(l), lineVector(r)
	return point{lVector.x * rVector.x, lVector.y * lVector.y}
}

func pointInLine(p point, l line) bool {
	return l.start.x <= p.x && p.x <= l.end.x && l.start.y <= p.y && p.y <= l.end.y
}

func lengthToPoint(w wire, intersection point) int {
	cost := 0
	for _, segment := range w.segments {
		if !pointInLine(intersection, segment) {
			cost += lineLength(segment)
		} else {
			start := segment.start
			if segment.isReversed {
				start = segment.end
			}
			cost += abs(intersection.x - start.x + intersection.y - start.y)
			return cost
		}
	}
	return 1000000000
}

// TODO: Optimize this check!
func intersection(sLeft, sRight line) point {
	center := point{0, 0}
	var possibleIntersection point
	if sLeft.start.x == sLeft.end.x {
		possibleIntersection = point{sLeft.start.x, sRight.start.y}
	} else {
		possibleIntersection = point{sRight.start.x, sLeft.start.y}
	}
	if pointInLine(possibleIntersection, sLeft) &&
		pointInLine(possibleIntersection, sRight) {
		return possibleIntersection
	}
	return center
}

// TODO: Optimize the search for of the intersection!
func closestIntersectionToCenter(wLeft, wRight wire) int {
	distance := 10000000000
	for _, sLeft := range wLeft.segments {
		for _, sRight := range wRight.segments {
			newDistance := manhattanDistance(intersection(sLeft, sRight))
			if newDistance > 0 && newDistance < distance {
				distance = newDistance
			}
		}
	}
	return distance
}

// TODO: Optimize the search for of the intersection!
func shortestIntersectionToCenter(wLeft, wRight wire) int {
	length := 10000000000
	center := point{0, 0}
	for _, sLeft := range wLeft.segments {
		for _, sRight := range wRight.segments {
			possibleIntersection := intersection(sLeft, sRight)
			if possibleIntersection != center {
				cost := lengthToPoint(wLeft, possibleIntersection) + lengthToPoint(wRight, possibleIntersection)
				if cost > 0 && cost < length {
					length = cost
				}
			}
		}
	}
	return length
}

// Read the input
func wireFromString(lineStr string) wire {
	nReg, _ := regexp.Compile("[0-9]+")
	turns := strings.Split(lineStr, ",") // Get the individual turns
	lineEnd := point{0, 0}
	var currentWire wire
	for _, el := range turns {
		length, _ := strconv.Atoi(nReg.FindString(el))
		direction := string(el[0])
		if direction == "U" {
			newEnd := point{lineEnd.x, lineEnd.y + length}
			currentWire.segments = append(currentWire.segments, newLinePoints(lineEnd, newEnd))
			lineEnd = newEnd
		} else if direction == "D" {
			newEnd := point{lineEnd.x, lineEnd.y - length}
			currentWire.segments = append(currentWire.segments, newLinePoints(lineEnd, newEnd))
			lineEnd = newEnd
		} else if direction == "R" {
			newEnd := point{lineEnd.x + length, lineEnd.y}
			currentWire.segments = append(currentWire.segments, newLinePoints(lineEnd, newEnd))
			lineEnd = newEnd
		} else if direction == "L" {
			newEnd := point{lineEnd.x - length, lineEnd.y}
			currentWire.segments = append(currentWire.segments, newLinePoints(lineEnd, newEnd))
			lineEnd = newEnd
		} else {
			panic(fmt.Sprintf("Unknown direction %s!\n", direction))
		}
	}
	return currentWire
}

func readPanel(filename string) []wire {
	f, _ := os.Open(filename)
	s := bufio.NewScanner(f)
	var wires []wire
	// Get the line
	for s.Scan() {
		wires = append(wires, wireFromString(s.Text()))
	}
	return wires
}

/*-------|
|- MAIN -|
|-------*/

func main() {	
	// Tests - Pt1
	assert(
		closestIntersectionToCenter(
			wireFromString("R8,U5,L5,D3"),
			wireFromString("U7,R6,D4,L4")),
		6)
	assert(
		closestIntersectionToCenter(
			wireFromString("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			wireFromString("U62,R66,U55,R34,D71,R55,D58,R83")),
		159)
	assert(
		closestIntersectionToCenter(
			wireFromString("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			wireFromString("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")),
			135)

	// Tests - Pt2
	assert(
		shortestIntersectionToCenter(
			wireFromString("R8,U5,L5,D3"),
			wireFromString("U7,R6,D4,L4")),
		30)
	assert(
		shortestIntersectionToCenter(
			wireFromString("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			wireFromString("U62,R66,U55,R34,D71,R55,D58,R83")),
		610)
	assert(
		shortestIntersectionToCenter(
			wireFromString("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			wireFromString("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")),
		410)

	// Problem
	filename := "input.txt"
	panel := readPanel(filename)
	// Pt1
	fmt.Println(closestIntersectionToCenter(panel[0], panel[1]))
	// Pt2
	fmt.Println(shortestIntersectionToCenter(panel[0], panel[1]))
}
