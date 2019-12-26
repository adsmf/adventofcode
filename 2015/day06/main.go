package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	g := grid{}
	g.runInstructions("input.txt")
	return g.countLit()
}

func part2() int {
	return -1
}

type grid map[point]bool

func (g *grid) runInstructions(filename string) {
	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		ints := utils.GetInts(line)

		if len(ints) != 4 {
			fmt.Printf("Could not parse ints from %s (%#v)\n", line, ints)
			continue
		}
		for x := ints[0]; x <= ints[2]; x++ {
			for y := ints[1]; y <= ints[3]; y++ {
				switch {
				case strings.HasPrefix(line, "turn on"):
					(*g)[point{x, y}] = true

				case strings.HasPrefix(line, "turn off"):
					(*g)[point{x, y}] = false

				case strings.HasPrefix(line, "toggle"):
					(*g)[point{x, y}] = !(*g)[point{x, y}]
				}
			}
		}
	}
}

func (g grid) countLit() int {
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
