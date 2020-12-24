package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1, tiles := layTiles("input.txt")
	p2 := tileLife(tiles, 100)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func layTiles(filename string) (int, lobbyFloor) {
	tiles := lobbyFloor{}
	directions := load(filename)
	for _, line := range directions {
		pos := point{0, 0}
		for _, dir := range line {
			pos = pos.move(dir)
		}
		tiles[pos] = !tiles[pos]
	}
	count := 0
	for _, oddFlip := range tiles {
		if oddFlip {
			count++
		}
	}
	return count, tiles
}

func tileLife(tiles lobbyFloor, days int) int {
	for day := 0; day < days; day++ {
		tileCounts := lobbyFloorCount{}
		for pos, val := range tiles {
			if val {
				for dir := dirStart; dir < directionMax; dir++ {
					tileCounts[pos.move(dir)]++
				}
			}
		}
		next := lobbyFloor{}
		for pos, count := range tileCounts {
			if count == 2 || (count == 1 && tiles[pos]) {
				next[pos] = true
			}
		}
		tiles = next
	}
	return len(tiles)
}

type lobbyFloor map[point]bool
type lobbyFloorCount map[point]int

type point struct{ x, y int }

func (p point) move(dir direction) point {
	vec := directionVectors[dir]
	return point{p.x + vec.x, p.y + vec.y}
}

type direction int

const (
	dirE direction = iota
	dirSE
	dirSW
	dirW
	dirNW
	dirNE
	directionMax
	dirStart direction = 0
)

func (d direction) String() string {
	if d < dirStart || d >= directionMax {
		return "??"
	}
	return directionStrings[d]
}

var directionLookup = map[string]direction{"e": dirE, "se": dirSE, "sw": dirSW, "w": dirW, "nw": dirNW, "ne": dirNE}
var directionStrings = map[direction]string{dirE: "e", dirSE: "se", dirSW: "sw", dirW: "w", dirNW: "nw", dirNE: "ne"}
var directionVectors = map[direction]point{dirE: {2, 0}, dirSE: {1, 1}, dirSW: {-1, 1}, dirW: {-2, 0}, dirNW: {-1, -1}, dirNE: {1, -1}}

func load(filename string) [][]direction {
	directions := [][]direction{}
	for _, line := range utils.ReadInputLines(filename) {
		lineDirections := []direction{}
		for len(line) > 0 {
			if len(line) == 1 {
				lineDirections = append(lineDirections, directionLookup[line])
				break
			}
			if dir, match := directionLookup[line[0:2]]; match {
				lineDirections = append(lineDirections, dir)
				line = line[2:]
			} else if dir, match := directionLookup[line[0:1]]; match {
				lineDirections = append(lineDirections, dir)
				line = line[1:]
			}
		}
		directions = append(directions, lineDirections)
	}
	return directions
}

var benchmark = false
