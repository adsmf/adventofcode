package main

import (
	"fmt"
	"math/bits"

	"github.com/adsmf/adventofcode/utils/pathfinding/astar"
)

const designerNumber = 1352

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	start := point{1, 1}
	goal := point{31, 39}
	route, err := astar.Route(start, goal)
	if err != nil {
		fmt.Println("Solving failed")
		return -1
	}
	return len(route) - 1
}

func part2() int {
	start := point{1, 1}
	visitable := map[point]bool{start: true}

	open := []point{start}
	for turn := 0; turn < 50; turn++ {
		nextOpen := []point{}
		for _, pos := range open {
			for _, dir := range dirs {
				tryPos := point{pos.x + dir.x, pos.y + dir.y}
				if tryPos.x < 0 || tryPos.y < 0 {
					continue
				}
				if visitable[tryPos] {
					continue
				}
				if tryPos.isSpace(designerNumber) {
					nextOpen = append(nextOpen, tryPos)
					visitable[tryPos] = true
				}
			}
		}
		open = nextOpen
	}

	return len(visitable)
}

type point struct{ x, y int }

var dirs = []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (p point) Paths() []astar.Edge {
	edges := make([]astar.Edge, 0, 4)
	for _, dir := range dirs {
		tryPos := point{p.x + dir.x, p.y + dir.y}
		if tryPos.x < 0 || tryPos.y < 0 {
			continue
		}
		if tryPos.isSpace(designerNumber) {
			edges = append(edges, astar.Edge{
				To:   tryPos,
				Cost: 1,
			})
		}
	}
	return edges
}

func (p point) Heuristic(toNode astar.Node) astar.Cost {
	to := toNode.(point)
	distX := to.x - p.x
	distY := to.y - p.y
	if distX < 0 {
		distX *= -1
	}
	if distY < 0 {
		distY *= -1
	}
	return astar.Cost(distX + distY)
}

func (p point) isSpace(offset uint) bool {
	val := uint(p.x*p.x+3*p.x+2*p.x*p.y+p.y+p.y*p.y) + offset
	return bits.OnesCount(val)%2 == 0
}

var benchmark = false
