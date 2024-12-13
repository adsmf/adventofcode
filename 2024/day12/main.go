package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

const gridAlloc = 150 * 150

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1, p2 := 0, 0
	if gridAlloc < len(input) {
		panic("Insufficient grid allocation for input")
	}
	g := grid{}
	for ; input[g.w] != '\n'; g.w++ {
	}
	g.h = len(input) / (g.w + 1)
	visited := make([]bool, gridAlloc)
	open := make([]search, 0, 20)
	next := make([]search, 0, 20)
	// perimiterPoints := make(map[sideInfo]bool, 200)
	curRegion := make([]point, 0, 300)
	x, y := 0, 0
	for _, ch := range input {
		if ch == '\n' {
			y++
			x = 0
			continue
		}
		pos := point{x, y}
		if visited[g.index(pos)] {
			x++
			continue
		}
		visited[g.index(pos)] = true
		open = append(open, search{pos, byte(ch)})
		curRegion = curRegion[0:0]
		curRegion = append(curRegion, pos)
		curPerimiterLen := 0
		numSides := 0
		// clear(perimiterPoints)
		for len(open) > 0 {
			for _, cur := range open {
				outside := [8]bool{}
				for dir, n := range cur.pos.neighbours() {
					nVal := g.valAt(n)
					if dir < 4 {
						if nVal == byte(ch) {
							if visited[g.index(n)] {
								continue
							}
							visited[g.index(n)] = true
							curRegion = append(curRegion, n)
							next = append(next, search{n, byte(ch)})
						} else {
							curPerimiterLen++
						}
					}
					if nVal != byte(ch) {
						outside[dir] = true
					}
				}
				// Convex corner
				for i := range 4 {
					if outside[i] && outside[(i+1)&3] {
						numSides++
					}
				}
				// Concave corner
				for i := 4; i < 8; i++ {
					if outside[i] && !outside[(i-4)&3] && !outside[(i-3)&3] {
						numSides++
					}
				}
			}
			open, next = next, open[0:0]
		}
		p1 += len(curRegion) * curPerimiterLen
		p2 += len(curRegion) * (numSides)
		x++
	}
	return p1, p2
}

type search struct {
	pos   point
	plant byte
}

type grid struct {
	h, w int
}

func (g grid) inBound(p point) bool {
	return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h
}

func (g grid) valAt(p point) byte {
	if !g.inBound(p) {
		return 0
	}
	return input[g.index(p)]
}

func (g grid) index(p point) int {
	return int(p.x) + int(p.y)*int(g.w+1)
}

type point struct{ x, y int }

func (p point) neighbours() [8]point {
	return [8]point{
		// Adjacent
		{p.x - 1, p.y},
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
		// Diagonals
		{p.x - 1, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x + 1, p.y + 1},
		{p.x - 1, p.y + 1},
	}
}

type sideInfo struct {
	pos  point
	edge int
}

var benchmark = false
