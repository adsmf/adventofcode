package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	slopeMap := load("input.txt")

	return countRun(3, 1, slopeMap)
}

func part2() int {
	slopeMap := load("input.txt")
	count := countRun(1, 1, slopeMap)
	count *= countRun(3, 1, slopeMap)
	count *= countRun(5, 1, slopeMap)
	count *= countRun(7, 1, slopeMap)
	count *= countRun(1, 2, slopeMap)
	return count
}

func countRun(right, down int, slopeMap slopeTreeMap) int {
	x := 0
	count := 0
	for y := down; y <= slopeMap.height; y += down {
		x += right
		x %= slopeMap.width
		if slopeMap.trees[point{x, y}] {
			count++
		}
	}
	return count
}

func load(filename string) slopeTreeMap {
	slope := slopeTreeMap{
		trees: map[point]bool{},
	}
	lines := utils.ReadInputLines(filename)
	slope.height = len(lines) - 1
	slope.width = len(lines[0])
	for y, line := range lines {
		for x, symbol := range line {
			switch symbol {
			case '.':
			case '#':
				slope.trees[point{x, y}] = true
			default:
				panic(fmt.Sprintf("Unknown symbol '%c'", symbol))
			}
		}
	}
	return slope
}

type point struct {
	x, y int
}
type slopeTreeMap struct {
	width  int
	height int
	trees  map[point]bool
}
