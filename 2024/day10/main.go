package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	g := grid{}
	for ; input[g.w] != '\n'; g.w++ {
	}
	g.h = int8(len(input) / int(g.w+1))
	open := make([]search, 0, 50)
	next := make([]search, 0, 50)
	modified := make([]int, 0, 10)
	reset := func() {
		for _, pos := range modified {
			input[pos] &= 0x3f
		}
		modified = modified[0:0]
	}
	p1, p2 := 0, 0
	x, y := int8(0), int8(0)
	for _, ch := range input {
		switch ch {
		case '\n':
			y++
			x = 0
			continue
		case '0':
			reset()
			open = append(open, search{point{x, y}, '0'})
			for len(open) > 0 {
				for _, cur := range open {
					nextVal := cur.val + 1
					for _, n := range cur.pos.neighbours() {
						gVal := g.valAt(n)
						if gVal&0x3f == nextVal {
							switch gVal {
							case '9':
								p1++
								p2++
								input[g.index(n)] |= 0x80
								modified = append(modified, g.index(n))
							case '9' | 0x80:
								p2++
							default:
								next = append(next, search{n, nextVal})
							}

						}
					}
				}
				open, next = next, open[0:0]
			}
			reset()
		}
		x++
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
	pos point
	val byte
}

var benchmark = false
