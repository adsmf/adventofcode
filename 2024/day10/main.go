package main

import (
	_ "embed"
	"fmt"
	"slices"
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
	g := grid{}
	open := make([]search, 0, 1500)
	next := make([]search, 0, 1500)
	x, y := int8(0), int8(0)
	for _, ch := range input {
		switch ch {
		case '\n':
			g.w = x
			y++
			x = 0
		case '0':
			open = append(open, search{point{x, y}, point{x, y}, '0'})
			x++
		default:
			x++

		}
	}
	g.h = y
	trails := make([]pointPair, 0, 1500)

	p2 := 0
	for len(open) > 0 {
		for _, cur := range open {
			nextVal := cur.val + 1
			for _, n := range cur.pos.neighbours() {
				if g.valAt(n) == nextVal {
					if nextVal == '9' {
						p2++
						trails = append(trails, makePointPair(g, cur.start, n))
						continue
					}
					next = append(next, search{cur.start, n, nextVal})
				}
			}
		}
		open, next = next, open[0:0]
	}
	slices.SortFunc(trails, func(a, b pointPair) int {
		return int(a - b)
	})
	p1 := 0
	last := pointPair(0)
	for _, trail := range trails {
		if trail != last {
			last = trail
			p1++
		}
	}
	return p1, p2
}

type grid struct {
	w, h int8
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

type point struct{ x, y int8 }

func (p point) neighbours() [4]point {
	return [4]point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

type search struct {
	start point
	pos   point
	val   byte
}

type pointPair int32

func makePointPair(g grid, a, b point) pointPair {
	return pointPair(g.index(a)<<16 + g.index(b))
}

var benchmark = false
