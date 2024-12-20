package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

const (
	gridAlloc = 21_000
)

type searchSet [gridAlloc]int32

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
	g.h = len(input) / (g.w + 1)
	var start point
	for i, ch := range input {
		switch ch {
		case 'S':
			start = g.fromIndex(i)
			input[i] = '.'
		case 'E':
			input[i] = '.'
		}
	}
	var dists = searchSet{}
	baseRoute(g, &dists, start)
	p1, p2 := 0, 0
	cheatSearch(g, &dists, 20, 100, func(dist int) {
		p2++
		if dist == 2 {
			p1++
		}
	})
	return p1, p2
}

func cheatSearch(g grid, dist *searchSet, maxCheat int, minSaving int, callback func(cheatDist int)) {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			cur := point{x, y}
			curDist := dist[g.index(cur)]
			if g.valAt(cur) != '#' {
				for cX := x - maxCheat; cX <= x+maxCheat; cX++ {
					xDiff := max(cX-x, x-cX)
					remCheat := maxCheat - xDiff
					for cY := y - remCheat; cY <= y+remCheat; cY++ {
						yDiff := max(cY-y, y-cY)
						cPos := point{cX, cY}
						cVal := g.valAt(cPos)
						if cVal == '.' {
							nDist := dist[g.index(cPos)]
							saving := nDist - curDist - int32(xDiff) - int32(yDiff)
							if saving >= int32(minSaving) {
								callback(xDiff + yDiff)
							}
						}
					}
				}
			}
		}
	}
}

func baseRoute(g grid, dist *searchSet, start point) {
	lastPoint := point{0, 0}
	done := point{-1, -1}
	for cur := start; cur != done; {
		nextPoint := done
		g.eachNeighbour(cur, func(neigh point) {
			if neigh == lastPoint {
				return
			}
			ni := g.index(neigh)
			if !g.inBound(neigh) {
				return
			}
			dist[ni] = dist[g.index(cur)] + 1
			nextPoint = neigh
		})
		cur, lastPoint = nextPoint, cur
	}
}

type grid struct {
	h, w int
}

func (g grid) inBound(p point) bool    { return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h }
func (g grid) index(p point) int       { return int(p.x) + int(p.y)*int(g.w+1) }
func (g grid) fromIndex(idx int) point { return point{idx % (g.w + 1), idx / (g.w + 1)} }
func (g grid) valAt(p point) byte {
	if !g.inBound(p) {
		return 0
	}
	return input[g.index(p)]
}

func (g grid) eachNeighbour(s point, callback func(next point)) {
	for _, dir := range dirs {
		next := s.add(dir)
		if g.valAt(next) != '#' {
			callback(next)
		}
	}
}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }

var dirs = [4]point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

var benchmark = false
