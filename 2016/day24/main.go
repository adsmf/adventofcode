package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
	"github.com/adsmf/adventofcode/utils/pathfinding/astar"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	duct := load("input.txt")
	duct.simplify()
	dists := duct.dists()
	items := []int{}
	for _, item := range duct.targets {
		items = append(items, int(item))
	}
	bestDist := utils.MaxInt
	for _, perm := range utils.PermuteInts(items) {
		dist := dists['0'][byte(perm[0])]
		for i := 0; i < len(perm)-1; i++ {
			dist += dists[byte(perm[i])][byte(perm[i+1])]
		}
		if dist < bestDist {
			bestDist = dist
		}
	}
	return bestDist
}

func part2() int {
	duct := load("input.txt")
	duct.simplify()
	dists := duct.dists()
	items := []int{}
	for _, item := range duct.targets {
		items = append(items, int(item))
	}
	bestDist := utils.MaxInt
	for _, perm := range utils.PermuteInts(items) {
		dist := dists['0'][byte(perm[0])]
		for i := 0; i < len(perm)-1; i++ {
			dist += dists[byte(perm[i])][byte(perm[i+1])]
		}
		dist += dists['0'][byte(perm[len(items)-1])]
		if dist < bestDist {
			bestDist = dist
		}
	}
	return bestDist
}

type tileType byte

const (
	tileBlocked tileType = iota
	tileEmpty
)

type blueprint struct {
	grid          map[point]tileType
	targets       map[point]byte
	pos           point
	width, height int
}

var activeBlueprint blueprint

func (b blueprint) dists() map[byte]map[byte]int {
	dists := map[byte]map[byte]int{'0': {}}
	items := []byte{'0'}
	positions := []point{b.pos}
	for pos, target := range b.targets {
		items = append(items, target)
		positions = append(positions, pos)
		dists[target] = map[byte]int{}
	}
	activeBlueprint = b
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			from := positions[i]
			to := positions[j]

			route, err := astar.Route(from, to)
			if err == nil {
				dist := len(route) - 1
				dists[items[i]][items[j]] = dist
				dists[items[j]][items[i]] = dist
			}
		}
	}
	return dists
}

func (b blueprint) simplify() {
	done := false

	for !done {
		done = true
		for pos := range b.grid {
			if pos == b.pos {
				continue
			}
			if _, found := b.targets[pos]; found {
				continue
			}
			countAdjacent := 0
			for _, adjacent := range pos.adjacent() {
				if b.grid[adjacent] != tileBlocked {
					countAdjacent++
				}
			}
			if countAdjacent < 2 {
				delete(b.grid, pos)
				done = false
			}
		}
	}
}

func (b blueprint) String() string {
	sb := strings.Builder{}
	sb.Grow(b.height * b.width)
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			pos := point{x, y}
			if b.pos == pos {
				sb.WriteByte('0')
			} else if char, isTarget := b.targets[pos]; isTarget {
				sb.WriteByte(char)
			} else {
				switch b.grid[pos] {
				case tileBlocked:
					sb.WriteByte('#')
				case tileEmpty:
					sb.WriteByte('.')
				}
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type point struct{ x, y int }

func (p point) Heuristic(toNode astar.Node) astar.Cost {
	to := toNode.(point)
	dX, dY := p.x-to.x, p.y-to.y
	if dX < 0 {
		dX *= -1
	}
	if dY < 0 {
		dY *= -1
	}
	return astar.Cost(dX + dY)
}

func (p point) Paths() []astar.Edge {
	edges := []astar.Edge{}

	for _, next := range p.adjacent() {
		if activeBlueprint.grid[next] != tileBlocked {
			edges = append(edges, astar.Edge{
				To:   next,
				Cost: 1,
			})
		}
	}

	return edges
}

func (p point) adjacent() []point {
	return []point{
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}
}

func load(filename string) blueprint {
	duct := blueprint{
		grid:    make(map[point]tileType),
		targets: make(map[point]byte),
	}
	lines := utils.ReadInputLines(filename)
	duct.height = len(lines)
	duct.width = len(lines[0])
	for y, line := range lines {
		for x, char := range line {
			pos := point{x, y}
			switch char {
			case '#':
				// Wall!
			case '.':
				duct.grid[pos] = tileEmpty
			case '0':
				duct.grid[pos] = tileEmpty
				duct.pos = pos
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				duct.grid[pos] = tileEmpty
				duct.targets[pos] = byte(char)
			}
		}
	}
	return duct
}

var benchmark = false
