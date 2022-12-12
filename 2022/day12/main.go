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
	grid, start, end := loadGrid()
	const maxOpenSet = 100
	openSet, nextOpen := [maxOpenSet]point{end}, [maxOpenSet]point{}
	openCount, nextOpenCount := 1, 0
	visited := [_grid_max_width][_grid_max_height]bool{}
	for step := 1; openCount > 0; step++ {
		nextOpenCount = 0
		for i := 0; i < openCount; i++ {
			pos := openSet[i]
			posHeight := grid.get(pos)
			if posHeight == 2 && p2 == -1 {
				p2 = step
			}
			neighbours, numNeighbours := grid.validRoutes(pos)
			for n := 0; n < numNeighbours; n++ {
				neighbour := neighbours[n]
				if visited[neighbour.x][neighbour.y] {
					continue
				}
				if neighbour == start {
					p1 = step
					return p1, p2
				}
				visited[neighbour.x][neighbour.y] = true
				nextOpen[nextOpenCount] = neighbour
				nextOpenCount++
			}
		}
		openSet, nextOpen = nextOpen, openSet
		openCount = nextOpenCount
	}
	return p1, p2
}

func loadGrid() (gridData, point, point) {
	start, end := point{}, point{}
	g := gridData{}
	x, y := axis(0), axis(0)

	for _, ch := range input {
		if ch == '\n' {
			y++
			g.w = x
			x = 0
			continue
		}
		pos := point{x, y}
		if ch == 'S' {
			start = pos
			ch = 'a'
		} else if ch == 'E' {
			end = pos
			ch = 'z' + 1
		}
		g.set(pos, height(ch-'a'+1))
		x++
	}
	g.h = y
	return g, start, end
}

type height = uint8

const (
	_grid_max_width  = 80
	_grid_max_height = 41
)

type gridData struct {
	h, w    axis
	heights [_grid_max_width][_grid_max_height]height
}

func (g *gridData) set(pos point, val height) { g.heights[pos.x][pos.y] = val }
func (g gridData) get(pos point) height       { return g.heights[pos.x][pos.y] }

func (g gridData) validRoutes(p point) ([4]point, int) {
	minH := g.get(p)
	if minH > 1 {
		minH -= 1
	}
	routes := [4]point{}
	numPoints := 0
	for _, pos := range [4]point{
		p.up(),
		p.down(),
		p.left(),
		p.right(),
	} {
		if pos.x >= g.w {
			continue
		}
		if pos.y >= g.h {
			continue
		}
		if g.get(pos) >= minH {
			routes[numPoints] = pos
			numPoints++
		}
	}
	return routes, numPoints
}

type axis uint8
type point struct{ x, y axis }

// func newPoint(x, y point) point { return x | (y << _grid_width_bits) }
func (p point) up() point    { return point{p.x, p.y - 1} }
func (p point) down() point  { return point{p.x, p.y + 1} }
func (p point) left() point  { return point{p.x - 1, p.y} }
func (p point) right() point { return point{p.x + 1, p.y} }

// func (p point) x() point        { return p & (1<<_grid_width_bits - 1) }
// func (p point) y() point        { return p >> _grid_width_bits }
func (p point) String() string { return fmt.Sprintf("(%d,%d)", p.x, p.y) }

var benchmark = false
