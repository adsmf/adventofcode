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
	p1, p2 := -1, -1
	g, start, end := loadGrid()
	visited := map[point]bool{}
	openSet := map[point]bool{end: true}
	for step := 1; len(openSet) > 0; step++ {
		nextOpen := map[point]bool{}
		for pos := range openSet {
			neighbours := g.validRoutes(pos)
			for _, neighbour := range neighbours {
				if visited[neighbour] {
					continue
				}
				if g[pos] == 1 && p2 == -1 {
					p2 = step
				}
				if neighbour == start {
					p1 = step
					return p1, p2
				}
				visited[neighbour] = true
				nextOpen[neighbour] = true
			}
		}
		openSet = nextOpen
	}
	return p1, p2
}

func loadGrid() (gridData, point, point) {
	start, end := point{}, point{}
	g := gridData{}
	x, y := 0, 0

	for _, ch := range input {
		if ch == '\n' {
			y++
			x = 0
			continue
		}
		pos := point{x, y}
		if ch == 'S' {
			start = pos
			ch = 'a'
		} else if ch == 'E' {
			end = pos
			ch = 'z'
		}
		g[pos] = ch - 'a'
		x++
	}

	return g, start, end
}

type gridData map[point]uint8

func (g gridData) validRoutes(p point) []point {
	minH := g[p] - 1
	routes := []point{}
	for _, pos := range []point{
		p.up(),
		p.down(),
		p.left(),
		p.right(),
	} {
		posH, found := g[pos]
		if !found {
			continue
		}
		if posH >= minH {
			routes = append(routes, pos)
		}
	}
	return routes
}

type point struct {
	x, y int
}

func (p point) add(a point) point {
	return point{
		x: p.x + a.x,
		y: p.y + a.y,
	}
}

func (p point) up() point    { return p.add(point{0, 1}) }
func (p point) down() point  { return p.add(point{0, -1}) }
func (p point) left() point  { return p.add(point{-1, 0}) }
func (p point) right() point { return p.add(point{1, 0}) }

var benchmark = false
