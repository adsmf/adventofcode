package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	g1, g2 := loadInput()
	p1 := count(g1)
	p2 := count(g2)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func count(g grid) int {
	total := 0
	for _, count := range g {
		if count >= 2 {
			total++
		}
	}
	return total
}

func loadInput() (grid, grid) {
	g1 := grid{}
	g2 := grid{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		ints := utils.GetInts(line)
		x1, y1 := ints[0], ints[1]
		x2, y2 := ints[2], ints[3]
		dX, dY := x2-x1, y2-y1

		count := int(math.Max(math.Abs(float64(dX)), math.Abs(float64(dY))))
		dX /= count
		dY /= count
		incG1 := dX == 0 || dY == 0

		for i, x, y := 0, x1, y1; i <= count; i, x, y = i+1, x+dX, y+dY {
			if incG1 {
				g1[point{x, y}]++
			}
			g2[point{x, y}]++
		}
	}
	return g1, g2
}

type grid map[point]int
type point struct{ x, y int }

var benchmark = false
