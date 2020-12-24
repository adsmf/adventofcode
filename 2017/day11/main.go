package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	route := load("input.txt")
	end, p2 := traceRoute(point{0, 0}, route)
	p1 := end.movesFrom(point{0, 0})
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func traceRoute(start point, route []direction) (point, int) {
	pos := start
	max := 0
	for _, dir := range route {
		pos = pos.move(dir)
		dist := pos.movesFrom(point{0, 0})
		if dist > max {
			max = dist
		}
	}
	return pos, max
}

func load(filename string) []direction {
	inputBytes, _ := ioutil.ReadFile(filename)
	return parseRoute(strings.TrimSpace(string(inputBytes)))
}

func parseRoute(input string) []direction {
	route := []direction{}
	for _, dirString := range strings.Split(input, ",") {
		dir, match := directionLookup[dirString]
		if !match {
			panic(fmt.Errorf("No match for %s", dirString))
		}
		route = append(route, dir)
	}
	return route
}

type direction int

const (
	dirN direction = iota
	dirNE
	dirSE
	dirS
	dirSW
	dirNW
	directionMax
	dirStart direction = 0
)

var directionLookup = map[string]direction{"n": dirN, "ne": dirNE, "se": dirSE, "s": dirS, "sw": dirSW, "nw": dirNW}
var directionStrings = map[direction]string{dirN: "n", dirNE: "ne", dirSE: "se", dirS: "s", dirSW: "sw", dirNW: "nw"}
var directionVectors = map[direction]point{
	dirN:  {0, -2},
	dirNE: {1, -1},
	dirSE: {1, 1},
	dirS:  {0, 2},
	dirSW: {-1, 1},
	dirNW: {-1, -1},
}

func (d direction) String() string {
	if d < dirStart || d >= directionMax {
		return "??"
	}
	return directionStrings[d]
}

type point struct{ x, y int }

func (p point) move(dir direction) point {
	vec := directionVectors[dir]
	return point{p.x + vec.x, p.y + vec.y}
}

func (p point) movesFrom(to point) int {
	xDist, yDist := p.x-to.x, p.y-to.y
	if xDist < 0 {
		xDist *= -1
	}
	if yDist < 0 {
		yDist *= -1
	}
	yDist -= xDist
	yDist /= 2
	return xDist + yDist
}

var benchmark = false
