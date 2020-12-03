package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part2() int {
	return loadAndCalculate("input.txt", []point{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}})
}

func part1() int {
	return loadAndCalculate("input.txt", []point{{3, 1}})
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

type point struct {
	x, y int
}
