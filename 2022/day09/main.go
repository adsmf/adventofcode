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

func solve() (counter, counter) {
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
	return visited.c1, visited.c2
}

const visitBounds = 350
const visitSize = visitBounds*2 + 1

type visitMap struct {
	vis    [visitSize * visitSize]byte
	c1, c2 counter
}

func (v *visitMap) markVisited(p1 point, p2 point) {
	p1id, p2id := p1.id(), p2.id()
	if v.vis[p1id]&1 == 0 {
		v.vis[p1id] |= 1
		v.c1++
	}
	if v.vis[p2id]>>1 == 0 {
		v.vis[p2id] |= 2
		v.c2++
	}
}

var directions = []point{'R': {1, 0}, 'L': {-1, 0}, 'U': {0, -1}, 'D': {0, 1}}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }
func (p point) sub(o point) point { return point{p.x - o.x, p.y - o.y} }
func (p point) id() int           { return (p.y+visitBounds)*visitSize + (p.x + visitBounds) }

func (p point) reduce() (point, bool) {
	var cX, cY bool
	p.x, cX = reduceAxis(p.x)
	p.y, cY = reduceAxis(p.y)
	return p, cX || cY
}
func reduceAxis(v int) (int, bool) {
	if v > 1 {
		return -1, true
	}
	if v < -1 {
		return 1, true
	}
	return -v, false
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; pos < len(in) && in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

type counter uint16

var benchmark = false
