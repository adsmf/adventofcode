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
	g := loadInput("input.txt")
	for i := 0; i < 100; i++ {
		g = g.step(false)
	}
	return g.countLit()
}

func part2() int {
	g := loadInput("input.txt")
	for i := 0; i < 100; i++ {
		g = g.step(true)
	}
	return g.countLit()
}

func loadInput(filename string) binaryGrid {
	grid := binaryGrid{}

	for y, line := range utils.ReadInputLines(filename) {
		for x, char := range line {
			if char == '#' {
				grid[point{x, y}] = true
			}
		}
	}

	return grid
}

type binaryGrid map[point]bool

func (g binaryGrid) step(setCorners bool) binaryGrid {
	ng := binaryGrid{}

	if setCorners {
		g[point{0, 0}] = true
		g[point{0, 99}] = true
		g[point{99, 0}] = true
		g[point{99, 99}] = true
	}

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			curPos := point{x, y}
			cur := g[curPos]

			numLit := 0
			for _, n := range curPos.neighbours() {
				if g[n] {
					numLit++
				}
			}

			if numLit == 3 || (numLit == 2 && cur) {
				ng[curPos] = true
			}

		}
	}

	if setCorners {
		ng[point{0, 0}] = true
		ng[point{0, 99}] = true
		ng[point{99, 0}] = true
		ng[point{99, 99}] = true
	}

	return ng
}

func (g binaryGrid) countLit() int {
	lit := 0
	for _, light := range g {
		if light {
			lit++
		}
	}
	return lit
}

type point struct {
	x, y int
}

func (p point) neighbours() []point {
	return []point{
		point{p.x - 1, p.y - 1},
		point{p.x + 0, p.y - 1},
		point{p.x + 1, p.y - 1},
		point{p.x - 1, p.y + 0},
		point{p.x + 1, p.y + 0},
		point{p.x - 1, p.y + 1},
		point{p.x + 0, p.y + 1},
		point{p.x + 1, p.y + 1},
	}
}
