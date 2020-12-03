package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func mainAlt() {
	fmt.Printf("Part 1: %d\n", part1alt())
	fmt.Printf("Part 2: %d\n", part2alt())
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

func part2alt() int {
	return loadAndCalculate(
		"input.txt",
		[]point{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}},
	)
}

func part1alt() int {
	return loadAndCalculate(
		"input.txt",
		[]point{{3, 1}},
	)
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

func loadAndCalculate(filename string, slopes []point) int {
	lines := utils.ReadInputLines(filename)
	width := len(lines[0])
	slopeCounts := make([]int, len(slopes))
	for y, line := range lines {
		for slopeIndex, slope := range slopes {
			if y%slope.y == 0 {
				checkPos := (slope.x * (y / slope.y)) % width
				if line[checkPos] == '#' {
					slopeCounts[slopeIndex]++
				}
			}
		}
	}
	total := 1
	for _, count := range slopeCounts {
		total *= count
	}
	return total
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
