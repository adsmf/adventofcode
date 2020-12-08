package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	g := load()
	p1 := part1(g)
	p2 := part2(g)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2:%s\n", p2)
	}
}

func part1(g grid) int {
	count := 0
	for _, pixel := range g {
		if pixel {
			count++
		}
	}
	return count
}

func part2(g grid) string {
	return "\n" + g.sprint()
}

func load() grid {
	g := make(grid, 50*6)
	for _, line := range utils.ReadInputLines("input.txt") {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "rect":
			bounds := utils.GetInts(parts[1])
			g.rect(bounds[0], bounds[1])
		case "rotate":
			args := utils.GetInts(strings.Join(parts[2:], " "))
			switch parts[1] {
			case "row":
				g.rotRow(args[0], args[1])
			case "column":
				g.rotCol(args[0], args[1])
			}
		}
	}
	return g
}

type grid []bool

func (g grid) rect(w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g[y*50+x] = true
		}
	}
}

func (g grid) rotRow(row, amt int) {
	orig := make(grid, len(g))
	copy(orig, g)

	for i := 0; i < 50; i++ {
		src := row*50 + i - amt
		dest := row*50 + i
		if i-amt < 0 {
			src += 50
		}
		g[dest] = orig[src]
	}
}

func (g grid) rotCol(col, amt int) {
	orig := make(grid, len(g))
	copy(orig, g)
	for y := 0; y < 6; y++ {

		src := (y-amt)*50 + col
		if src < 0 {
			src += 50 * 6
		}
		dest := y*50 + col
		g[dest] = orig[src]
	}
}

func (g grid) sprint() string {
	count := 0
	output := ""
	for i := 0; i < len(g); i++ {
		if g[i] {
			output += fmt.Sprint("█")
			count++
		} else {
			output += fmt.Sprint(".")
		}
		if i%50 == 49 {
			output += fmt.Sprintln()
		}
	}
	return output
}
