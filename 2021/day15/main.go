package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
	"github.com/adsmf/adventofcode/utils/pathfinding/astar"
)

//go:embed input.txt
var input string

func main() {
	g, size := load(input)
	p1 := part1(g, size)
	p2 := part2(g, size)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(g grid, size point) int {
	start := g[point{0, 0}]
	goal := g[point{size.x - 1, size.y - 1}]
	route, err := astar.Route(start, goal)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	cost := 0
	for _, edge := range route {
		cost += edge.(cell).value
	}
	return cost - start.value
}

func part2(g grid, size point) int {
	bigGrid := make(grid, len(g)*25)
	for pos, v := range g {
		for cX := 0; cX < 5; cX++ {
			for cY := 0; cY < 5; cY++ {
				bigPos := point{pos.x + cX*size.x, pos.y + cY*size.y}
				bigVal := (v.value + cY + cX)
				for bigVal > 9 {
					bigVal -= 9
				}
				bigGrid[bigPos] = cell{
					point: bigPos,
					value: bigVal,
					grid:  &bigGrid,
				}
			}
		}
	}
	start := bigGrid[point{0, 0}]
	goal := bigGrid[point{size.x*5 - 1, size.y*5 - 1}]
	route, err := astar.Route(start, goal)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	cost := 0
	for _, edge := range route {
		cost += edge.(cell).value
	}
	return cost - start.value
}

func load(in string) (grid, point) {
	g := make(grid, len(in))
	x, y := 0, 0
	maxX := 0
	for _, ch := range []byte(in) {
		switch ch {
		case '\n':
			y++
			maxX = x
			x = 0
		default:
			pos := point{x, y}
			g[pos] = cell{
				pos,
				&g,
				int(ch - '0'),
			}
			x++
		}
	}
	max := point{maxX, y}
	return g, max
}

type grid map[point]cell
type cell struct {
	point point
	grid  *grid
	value int
}

func (c cell) Heuristic(to astar.Node) astar.Cost {
	return astar.Cost(utils.IntAbs(c.point.x-to.(cell).point.x) + utils.IntAbs(c.point.y-to.(cell).point.y))
}
func (c cell) Paths() []astar.Edge {
	edges := []astar.Edge{}
	for _, n := range c.point.neighbours() {
		if target, found := (*c.grid)[n]; found {
			edges = append(edges, astar.Edge{
				To:   target,
				Cost: astar.Cost(target.value),
			})
		}
	}
	return edges
}

type point struct {
	x, y int
}

func (p point) neighbours() []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}
func (p point) Heuristic() {

}

var benchmark = false
