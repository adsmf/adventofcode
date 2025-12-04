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
	g := loadGrid()
	initial := 0
	removed := 0

	addAdjacent := func(counts []byte, idx int, delta int) {
		pos := g.fromIndex(idx)
		for xOff := -1; xOff <= 1; xOff++ {
			for yOff := -1; yOff <= 1; yOff++ {
				adjPos := point{pos.x + xOff, pos.y + yOff}
				adjIdx := g.toIndex(adjPos)
				if !g.inBound(adjPos) {
					continue
				}
				if adjIdx != idx && g.tiles[adjIdx] {
					counts[adjIdx] += byte(delta)
				}
			}
		}
	}

	counts := make([]byte, len(g.tiles))
	next := make([]byte, len(g.tiles))
	for idx, tile := range g.tiles {
		if !tile {
			continue
		}
		addAdjacent(counts, idx, 1)
	}

	done := false
	for !done {
		done = true
		copy(next, counts)
		for idx, tile := range g.tiles {
			if tile && counts[idx] < 4 {
				removed++
				done = false
				g.tiles[idx] = false
				addAdjacent(next, idx, -1)
			}
		}
		if initial == 0 {
			initial = removed
		}
		next, counts = counts, next
	}

	return initial, removed
}

type point struct{ x, y int }

func loadGrid() grid {
	g := grid{
		tiles: make([]bool, 0, 20000),
	}
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '@':
			g.tiles = append(g.tiles, true)
			x++
		case '.':
			g.tiles = append(g.tiles, false)
			x++
		case '\n':
			g.w = x
			x = 0
			y++
		}
	}
	g.h = y
	return g
}

type grid struct {
	h, w  int
	tiles []bool
}

func (g grid) inBound(p point) bool    { return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h }
func (g grid) toIndex(p point) int     { return p.x + p.y*g.w }
func (g grid) fromIndex(idx int) point { return point{idx % g.w, idx / g.w} }

var benchmark = false
