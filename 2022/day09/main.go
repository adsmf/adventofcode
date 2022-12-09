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
	positions := [10]point{}
	visited := visitMap{}
	var dist int
	for pos := 0; pos < len(input); pos++ {
		headMove := directions[input[pos]]
		dist, pos = getInt(input, pos+2)
		for i := 0; i < dist; i++ {
			positions[0] = positions[0].add(headMove)
			head := positions[0]
			for i := 1; i < 10; i++ {
				tail := positions[i]
				tailOffset := tail.sub(head)
				tailMove, moved := tailOffset.reduce()
				if !moved {
					break
				}
				positions[i] = tail.add(tailMove)
				head = positions[i]
			}
			visited.markVisited(positions[1], positions[9])
		}
	}
	return visited.counts()
}

const visitBounds = 350
const visitSize = visitBounds*2 + 1

type visitMap struct {
	vis      [visitSize][visitSize]byte
	min, max point
}

func (v *visitMap) markVisited(p1 point, p2 point) {
	v.min = minPoint(minPoint(v.min, p1), p2)
	v.max = maxPoint(maxPoint(v.max, p1), p2)
	v.vis[p1.x+visitBounds][p1.y+visitBounds] |= 1
	v.vis[p2.x+visitBounds][p2.y+visitBounds] |= 2
}

func (v visitMap) counts() (int, int) {
	c1, c2 := 0, 0
	min := point{v.min.x + visitBounds, v.min.y + visitBounds}
	max := point{v.max.x + visitBounds, v.max.y + visitBounds}
	for x := min.x; x <= max.x; x++ {
		for y := min.y; y <= max.y; y++ {
			vis := v.vis[x][y]
			if vis == 0 {
				continue
			}
			c1 += int(vis & 1)
			c2 += int((vis & 2) >> 1)
		}
	}
	return c1, c2
}

var directions = []point{'R': {1, 0}, 'L': {-1, 0}, 'U': {0, -1}, 'D': {0, 1}}

func minPoint(p1, p2 point) point { return point{min(p1.x, p2.x), min(p1.y, p2.y)} }
func maxPoint(p1, p2 point) point { return point{max(p1.x, p2.x), max(p1.y, p2.y)} }
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }
func (p point) sub(o point) point { return point{p.x - o.x, p.y - o.y} }

func (p point) reduce() (point, bool) {
	if p.x > 1 {
		return point{-1, crop1(-p.y)}, true
	}
	if p.x < -1 {
		return point{1, crop1(-p.y)}, true
	}
	if p.y > 1 {
		return point{crop1(-p.x), -1}, true
	}
	if p.y < -1 {
		return point{crop1(-p.x), 1}, true
	}
	return point{}, false
}

func crop1(in int) int {
	if in == 0 {
		return 0
	}
	if in > 0 {
		return 1
	}
	return -1
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; pos < len(in) && in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

var benchmark = false
