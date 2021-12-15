package main

import (
	"container/heap"
	_ "embed"
	"fmt"
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
	return explore(g, size)
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
				}
			}
		}
	}
	return explore(bigGrid, point{size.x * 5, size.y * 5})
}

func explore(g grid, size point) int {
	start := point{0, 0}
	goal := point{size.x - 1, size.y - 1}
	visited := make(map[point]bool, len(g))
	open := &cellHeap{exploredCell{start, 0}}
	heap.Init(open)

	for len(*open) > 0 {
		cur := heap.Pop(open).(exploredCell)
		if cur.point == goal {
			return cur.totalRisk
		}
		if _, found := visited[cur.point]; found {
			continue
		}
		visited[cur.point] = true
		for _, n := range cur.point.neighbours() {
			if _, found := g[n]; !found {
				continue
			}
			next := exploredCell{
				point:     n,
				totalRisk: cur.totalRisk + g[n].value,
			}
			heap.Push(open, next)
		}
	}
	return -1
}

type cellHeap []exploredCell

func (c cellHeap) Len() int            { return len(c) }
func (c cellHeap) Less(i, j int) bool  { return c[i].totalRisk < c[j].totalRisk }
func (c cellHeap) Swap(i, j int)       { c[i], c[j] = c[j], c[i] }
func (c *cellHeap) Push(x interface{}) { *c = append(*c, x.(exploredCell)) }
func (c *cellHeap) Pop() interface{} {
	old := *c
	n := len(old)
	item := old[n-1]
	*c = old[0 : n-1]
	return item
}

type exploredCell struct {
	point     point
	totalRisk int
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
			g[pos] = cell{pos, int(ch - '0')}
			x++
		}
	}
	max := point{maxX, y}
	return g, max
}

type grid map[point]cell
type cell struct {
	point point
	value int
}

type point struct{ x, y int }

func (p point) neighbours() []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

var benchmark = false
