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
	g := grid{}
	open := []search{}
	next := []search{}
	x, y := 0, 0
	for _, ch := range input {
		switch ch {
		case '\n':
			g.w = x
			y++
			x = 0
		case '0':
			open = append(open, search{point{x, y}, point{x, y}, 0})
			x++
		default:
			x++

		}
	}
	g.h = y
	trails := map[pointPair]int{}
	for len(open) > 0 {
		for _, cur := range open {
			nextVal := cur.val + 1
			for _, n := range cur.pos.neighbours() {
				if g.valAt(n) == nextVal {
					if nextVal == 9 {
						trails[pointPair{cur.start, n}]++
						continue
					}
					next = append(next, search{cur.start, n, nextVal})
				}
			}
		}
		open, next = next, open[0:0]
	}
	p2 := 0
	for _, c := range trails {
		p2 += c
	}
	return len(trails), p2
}

type grid struct {
	w, h int
}

func (g grid) inBound(p point) bool {
	return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h
}

func (g grid) valAt(p point) int {
	if !g.inBound(p) {
		return -1
	}
	return int(input[p.x+p.y*(g.w+1)] - '0')
}

type point struct{ x, y int }

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
	val   int
}

type pointPair struct{ a, b point }

var benchmark = false
