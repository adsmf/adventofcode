package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	vals := utils.GetInts(input)
	minX, maxX, minY, maxY := vals[0], vals[1], vals[2], vals[3]
	p1 := part1(minY)
	p2 := part2(minX, maxX, minY, maxY)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(minY int) int {
	return -minY * (-minY - 1) / 2
}

func part2(minX, maxX, minY, maxY int) int {
	countValid := 0
	startX := 1
	for minX > startX*(startX+1)/2 {
		startX++
	}
	for dY := minY; dY < -minY; dY++ {
		for dX := startX; dX <= maxX; dX++ {
			hit := runSim(dX, dY, minX, maxX, minY, maxY)
			if hit {
				countValid++
			}
		}
	}
	return countValid
}

func runSim(dX, dY int, minX, maxX, minY, maxY int) bool {
	x, y := 0, 0
	for x <= maxX && y >= minY {
		if minX <= x && y <= maxY {
			return true
		}
		x += dX
		y += dY
		if dX > 0 {
			dX--
		}
		dY--
	}
	return false
}

var benchmark = false
