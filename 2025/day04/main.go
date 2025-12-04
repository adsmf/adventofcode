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
	g := loadGrid()
	initial := 0
	removed := 0

	addAdjacent := func(idx int, delta int) {
		pos := g.fromIndex(idx)
		for xOff := -1; xOff <= 1; xOff++ {
			for yOff := -1; yOff <= 1; yOff++ {
				adjPos := point{pos.x + xOff, pos.y + yOff}
				adjIdx := g.toIndex(adjPos)
				if !g.inBound(adjPos) {
					continue
				}
				if adjIdx != idx && input[adjIdx] >= '0' {
					input[adjIdx] += byte(delta)
				}
			}
		}
	}

	toRemove := make([]int, 0, 1500)
	for idx, tile := range input {
		if tile < '0' {
			continue
		}
		addAdjacent(idx, 1)
	}

	for {
		toRemove = toRemove[0:0]
		for idx, ch := range input {
			if ch >= '0' && ch-'0' < 4 {
				toRemove = append(toRemove, idx)
			}
		}
		newRemovals := len(toRemove)
		if initial == 0 {
			initial = newRemovals
		}
		if newRemovals == 0 {
			break
		}
		removed += newRemovals
		for _, idx := range toRemove {
			input[idx] = 0
			addAdjacent(idx, -1)
		}
	}

	return initial, removed
}

type point struct{ x, y int }

func loadGrid() grid {
	g := grid{}
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '@':
			input[pos] = '0'
			x++
		case '.':
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
	h, w int
}

func (g grid) inBound(p point) bool    { return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h }
func (g grid) toIndex(p point) int     { return p.x + p.y*(g.w+1) }
func (g grid) fromIndex(idx int) point { return point{idx % (g.w + 1), idx / (g.w + 1)} }

var benchmark = false
