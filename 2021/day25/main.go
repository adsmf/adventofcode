package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := settle(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func settle(input string) int {
	g := load(input)
	for moves := 1; moves < 1000; moves++ {
		movedEast := g.stepEast()
		movedSouth := g.stepSouth()
		if !(movedEast || movedSouth) {
			return moves
		}
	}
	return -1
}

func load(input string) grid {
	lines := utils.GetLines(input)
	h, w := len(lines), len(lines[0])

	g := grid{
		grid: make([]byte, 0, len(input)),
		h:    h,
		w:    w,
	}
	for _, line := range lines {
		g.grid = append(g.grid, []byte(line)...)
	}
	return g
}

type grid struct {
	grid []byte
	h, w int
}

func (g *grid) stepEast() bool {
	moved := false
	next := make([]byte, len(g.grid))
	copy(next, g.grid)
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			pos := x + y*g.w
			if g.grid[pos] != '>' {
				continue
			}
			nextPos := ((x + 1) % g.w) + y*g.w
			if g.grid[nextPos] == '.' {
				next[pos] = '.'
				next[nextPos] = '>'
				moved = true
			}
		}
	}
	g.grid = next
	return moved
}

func (g *grid) stepSouth() bool {
	moved := false
	next := make([]byte, len(g.grid))
	copy(next, g.grid)
	for i := 0; i < len(g.grid); i++ {
		if g.grid[i] != 'v' {
			continue
		}
		nextPos := (i + g.w) % len(g.grid)
		if g.grid[nextPos] == '.' {
			next[i] = '.'
			next[nextPos] = 'v'
			moved = true
		}
	}
	g.grid = next
	return moved
}

var benchmark = false
