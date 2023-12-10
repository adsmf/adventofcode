package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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
	pipes, start, max := loadMap()
	visited := make(map[point]bool, max.x*max.y/2)
	step := -1
	current := []point{start}
	for len(current) > 0 {
		step++
		next := []point{}
		for _, pos := range current {
			for _, neighbour := range pipes.validNeighbours(pos) {
				if !visited[neighbour] {
					next = append(next, neighbour)
					visited[neighbour] = true
				}
			}
		}
		current = next
	}
	p1 := step
	insideCount := 0
	for y := 0; y <= max.y; y++ {
		inside := false
		startCorner := tileNone
		for x := 0; x <= max.x; x++ {
			pos := point{x, y}
			tile := pipes[pos]
			inLoop := visited[pos]
			if inside && !inLoop {
				insideCount++
				continue
			}
			if !inLoop {
				if inside {
					insideCount++
				}
				continue
			}
			switch tile {
			case tileVertical:
				inside = !inside
			case tileCornerNE, tileCornerSE:
				startCorner = tile
			case tileCornerNW:
				if startCorner != tileCornerNE {
					inside = !inside
				}
			case tileCornerSW:
				if startCorner != tileCornerSE {
					inside = !inside
				}
			}
		}
	}
	return p1, insideCount
}

func part2() int {
	return -1
}

func loadMap() (pipeMap, point, point) {
	start := point{}
	pipes := pipeMap{}
	max := point{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		max.y = index
		for x, ch := range line {
			pos := point{x, index}
			pipes[pos] = symbols[byte(ch)]
			if ch == 'S' {
				start = pos
			}
			if x > max.x {
				max.x = x
			}
		}
		return false
	})
	return pipes, start, max
}

type pipeMap map[point]tileType

func (p pipeMap) validNeighbours(pos point) []point {
	initial := p[pos]
	neighbours := []point{}
	if pipeTypeConnections[initial].up && pipeTypeConnections[p[pos.up()]].down {
		neighbours = append(neighbours, pos.up())
	}
	if pipeTypeConnections[initial].down && pipeTypeConnections[p[pos.down()]].up {
		neighbours = append(neighbours, pos.down())
	}
	if pipeTypeConnections[initial].left && pipeTypeConnections[p[pos.left()]].right {
		neighbours = append(neighbours, pos.left())
	}
	if pipeTypeConnections[initial].right && pipeTypeConnections[p[pos.right()]].left {
		neighbours = append(neighbours, pos.right())
	}
	return neighbours
}

type point struct{ x, y int }

func (p point) add(a point) point    { return point{x: p.x + a.x, y: p.y + a.y} }
func (p point) up() point            { return p.add(point{0, -1}) }
func (p point) right() point         { return p.add(point{1, 0}) }
func (p point) down() point          { return p.add(point{0, 1}) }
func (p point) left() point          { return p.add(point{-1, 0}) }
func (p point) neighbours() [4]point { return [4]point{p.up(), p.right(), p.down(), p.left()} }

type tileType byte

const (
	tileNone tileType = iota
	tileGround
	tileVertical
	tileHorizontal
	tileCornerNE
	tileCornerNW
	tileCornerSW
	tileCornerSE
	tileStart
)

type pipeConnections struct{ up, right, down, left bool }

var pipeTypeConnections = map[tileType]pipeConnections{
	tileVertical:   {true, false, true, false},
	tileHorizontal: {false, true, false, true},
	tileCornerNE:   {true, true, false, false},
	tileCornerNW:   {true, false, false, true},
	tileCornerSW:   {false, false, true, true},
	tileCornerSE:   {false, true, true, false},
	tileStart:      {true, true, true, true},
}

var symbols = map[byte]tileType{
	'.': tileGround,
	'-': tileHorizontal,
	'|': tileVertical,
	'L': tileCornerNE,
	'J': tileCornerNW,
	'7': tileCornerSW,
	'F': tileCornerSE,
	'S': tileStart,
}

var benchmark = false
