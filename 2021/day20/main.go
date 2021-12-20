package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(input string) (int, int) {
	im, enahncement := load(input)
	for i := 0; i < 2; i++ {
		im.enhance(enahncement)
	}
	p1 := len(im.grid)
	for i := 2; i < 50; i++ {
		im.enhance(enahncement)
	}
	return p1, len(im.grid)
}

func load(input string) (image, []bool) {
	blocks := strings.Split(input, "\n\n")

	enhancement := make([]bool, 0, len(blocks[0]))
	for _, ch := range blocks[0] {
		enhancement = append(enhancement, ch == '#')
	}

	lines := strings.Split(blocks[1], "\n")
	g := make(grid, len(lines)*len(lines[0]))
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				g[point{x, y}] = true
			}
		}
	}
	im := image{
		grid:     g,
		inverted: false,
		min:      point{0, 0},
		max:      point{len(lines[0]), len(lines)},
	}
	return im, enhancement
}

type image struct {
	grid     grid
	inverted bool
	min, max point
}

type grid map[point]bool

func (im *image) enhance(enhancement []bool) {
	newGrid := make(grid, len(im.grid))
	nextInverted := im.inverted
	if enhancement[0] {
		nextInverted = !im.inverted
	}

	for y := im.min.y - 1; y <= im.max.y+1; y++ {
		for x := im.min.x - 1; x <= im.max.x+1; x++ {
			pos := point{x, y}
			lookup := 0
			if im.inverted {
				for _, n := range pos.neighbours() {
					lookup *= 2
					if !im.grid[n] {
						lookup += 1
					}
				}
			} else {
				for _, n := range pos.neighbours() {
					lookup *= 2
					if im.grid[n] {
						lookup += 1
					}
				}
			}
			value := enhancement[lookup]
			if nextInverted {
				value = !value
			}
			if value {
				newGrid[pos] = value
			}
		}
	}

	im.grid = newGrid
	im.inverted = nextInverted
	im.min = point{im.min.x - 1, im.min.y - 1}
	im.max = point{im.max.x + 1, im.max.y + 1}
}

type point struct {
	x, y int
}

func (p point) neighbours() []point {
	return []point{
		{p.x - 1, p.y - 1}, {p.x + 0, p.y - 1}, {p.x + 1, p.y - 1},
		{p.x - 1, p.y}, {p.x, p.y}, {p.x + 1, p.y},
		{p.x - 1, p.y + 1}, {p.x, p.y + 1}, {p.x + 1, p.y + 1},
	}
}

var benchmark = false
