package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1, p2 := 0, 0
	g := grid{}
	for ; input[g.w] != '\n'; g.w++ {
	}
	g.h = len(input) / (g.w + 1)
	visited := map[point]bool{}
	x, y := 0, 0
	for _, ch := range input {
		if ch == '\n' {
			y++
			x = 0
			continue
		}
		pos := point{x, y}
		if visited[pos] {
			x++
			continue
		}
		visited[pos] = true
		open := []search{{pos, byte(ch)}}
		next := []search{}
		curRegion := map[point]bool{pos: true}
		type sideInfo struct {
			pos  point
			edge int
		}
		curPerimiterLen := 0
		perimiterPoints := map[sideInfo]bool{}
		for len(open) > 0 {
			for _, cur := range open {
				for dir, n := range cur.pos.neighbours() {
					nVal := g.valAt(n)
					if nVal == byte(ch) {
						if curRegion[n] {
							continue
						}
						visited[n] = true
						curRegion[n] = true
						next = append(next, search{n, byte(ch)})
					} else {
						curPerimiterLen++
						perimiterPoints[sideInfo{n, dir}] = true
					}
				}
			}
			open, next = next, open[0:0]
		}

		numSides := 0
		for p := range perimiterPoints {
			numSides++
			open := []sideInfo{p}
			next := []sideInfo{}
			for len(open) > 0 {
				for _, cur := range open {
					for _, n := range cur.pos.neighbours() {
						side := sideInfo{n, cur.edge}
						if perimiterPoints[side] {
							delete(perimiterPoints, side)
							next = append(next, side)
						}
					}
				}
				open, next = next, open[0:0]
			}
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

func (p point) neighbours() [4]point {
	return [4]point{
		{p.x - 1, p.y},
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
	}
}

var benchmark = false
