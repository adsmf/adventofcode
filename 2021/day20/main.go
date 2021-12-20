package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
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
	g, enahncement := load(input)
	inverted := false
	for i := 0; i < 2; i++ {
		g, inverted = g.enhance(enahncement, inverted)
	}
	p1 := len(g)
	for i := 2; i < 50; i++ {
		g, inverted = g.enhance(enahncement, inverted)
	}
	return p1, len(g)
}

func load(input string) (grid, []bool) {
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
	return g, enhancement
}

type grid map[point]bool

func (g grid) enhance(enhancement []bool, inverted bool) (grid, bool) {
	newGrid := make(grid, len(g))
	nextInverted := inverted
	if enhancement[0] {
		nextInverted = !inverted
	}

	min, max := g.bounds()

	for y := min.y - 1; y <= max.y+1; y++ {
		for x := min.x - 1; x <= max.x+1; x++ {
			pos := point{x, y}
			lookup := 0
			if inverted {
				for _, n := range pos.neighbours() {
					lookup *= 2
					if !g[n] {
						lookup += 1
					}
				}
			} else {
				for _, n := range pos.neighbours() {
					lookup *= 2
					if g[n] {
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

	return newGrid, nextInverted
}

func (g grid) bounds() (point, point) {
	minX, maxX := utils.MaxInt, -utils.MaxInt
	minY, maxY := utils.MaxInt, -utils.MaxInt

	for pos := range g {
		if pos.x < minX {
			minX = pos.x
		}
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y < minY {
			minY = pos.y
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	return point{minX, minY}, point{maxX, maxY}
}

func (g grid) String() string {
	sb := strings.Builder{}
	min, max := g.bounds()
	for y := min.y - 1; y <= max.y+1; y++ {
		for x := min.x - 1; x <= max.x+1; x++ {
			if g[point{x, y}] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
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
