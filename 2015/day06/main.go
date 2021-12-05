package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	g := binaryGrid{}
	g.runInstructions("input.txt")
	return g.countLit()
}

func part2() int {
	g := dimmableGrid{}
	g.runInstructions("input.txt")
	return g.sumLit()
}

type binaryGrid map[point]bool

func (g *binaryGrid) runInstructions(filename string) {
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

func (g binaryGrid) countLit() int {
	lit := 0
	for _, light := range g {
		if light {
			lit++
		}
	}
	return lit
}

type dimmableGrid map[point]int

func (g *dimmableGrid) runInstructions(filename string) {
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
					(*g)[point{x, y}] += 1

				case strings.HasPrefix(line, "turn off"):
					(*g)[point{x, y}] -= 1
					if (*g)[point{x, y}] < 0 {
						(*g)[point{x, y}] = 0
					}

				case strings.HasPrefix(line, "toggle"):
					(*g)[point{x, y}] += 2
				}
			}
		}
	}
}

func (g dimmableGrid) sumLit() int {
	lit := 0
	for _, light := range g {

		lit += light

	}
	return lit
}

type point struct {
	x, y int
}

var benchmark = false
