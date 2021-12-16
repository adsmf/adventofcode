package main

import (
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
				bigPos := makePoint(pos.x()+cX*size.x(), pos.y()+cY*size.y())
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
	return explore(bigGrid, makePoint(size.x()*5, size.y()*5))
}

func explore(g grid, size point) int {
	start := makePoint(0, 0)
	goal := makePoint(size.x()-1, size.y()-1)
	visited := make(map[point]bool, len(g))
	open := make(cellStack, 1, 1000)
	open[0] = exploredCell{start, 0}

	for len(open) > 0 {
		cur := open.Pop()
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
			open.Insert(next)
		}
	}
	return -1
}

type cellStack []exploredCell

func (c *cellStack) Insert(ec exploredCell) {
	l, h := 0, len(*c)
	mid := 0
	for l != h {
		mid = (l + h) / 2
		if ec.totalRisk < (*c)[mid].totalRisk {
			h = mid
		} else {
			l = mid + 1
		}
	}
	if mid >= len(*c) {
		*c = append(*c, ec)
		return
	}
	*c = append((*c)[:mid+1], (*c)[mid:]...)
	(*c)[mid] = ec
}
func (c *cellStack) Pop() exploredCell {
	ec := (*c)[0]
	*c = (*c)[1:]
	return ec
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
			pos := makePoint(x, y)
			g[pos] = cell{pos, int(ch - '0')}
			x++
		}
	}
	max := makePoint(maxX, y)
	return g, max
}

type grid map[point]cell
type cell struct {
	point point
	value int
}

type point uint32

func makePoint(x, y int) point { return point(x | (y << 16)) }
func (p point) x() int         { return int(p & 0xffff) }
func (p point) y() int         { return int(p >> 16) }

func (p point) neighbours() []point {
	return []point{
		p - 1,
		p + 1,
		p - (1 << 16),
		p + (1 << 16),
	}
}

var benchmark = false
