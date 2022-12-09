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
	p1 := make(map[point]bool, 3000)
	p2 := make(map[point]bool, 3000)
	var dir byte
	var dist int
	for pos := 0; pos < len(input); pos++ {
		dir = input[pos]
		dist, pos = getInt(input, pos+2)
		headMove := directions[dir]
		for i := 0; i < dist; i++ {
			positions[0] = positions[0].add(headMove)
			for i := 1; i < 10; i++ {
				tailOffset := positions[i].sub(positions[i-1])
				tailMove := tailOffset.reduce()
				positions[i] = positions[i].add(tailMove)
			}
			p1[positions[1]] = true
			p2[positions[9]] = true
		}
	}
	return len(p1), len(p2)
}

var directions = map[byte]point{
	'R': {1, 0},
	'L': {-1, 0},
	'U': {0, -1},
	'D': {0, 1},
}

type point struct {
	x, y int
}

func (p point) add(o point) point {
	return point{p.x + o.x, p.y + o.y}
}

func (p point) sub(o point) point {
	return point{p.x - o.x, p.y - o.y}
}

func (p point) reduce() point {
	if p.x > 1 {
		return point{-1, crop1(-p.y)}
	}
	if p.x < -1 {
		return point{1, crop1(-p.y)}
	}
	if p.y > 1 {
		return point{crop1(-p.x), -1}
	}
	if p.y < -1 {
		return point{crop1(-p.x), 1}
	}
	return point{}
}

func crop1(in int) int {
	if in > 1 {
		return 1
	}
	if in < -1 {
		return -1
	}
	return in
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
