package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

// Ewww, magic numbers
const (
	gridAlloc = 2e4
	maxLoop   = 6e3
)

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	g := loadGrid()
	v := make(pointSet, gridAlloc)
	count, _ := runSim(g, v, point{-1, -1})
	return count, searchLoops(g, v)
}

func runSim(g grid, visit pointSet, simObst point) (int, bool) {
	cur := walkDir{g.start, dirUp}
	if visit != nil {
		visit[g.toMS(cur.p)] = true
	}
	unique := 1
	for i := 0; i < maxLoop; i++ {
		next := cur.move()
		for simObst == next.p || g.obstacle(next.p) {
			cur = cur.rotateRight()
			next = cur.move()
		}
		cur = next
		if !g.inBound(cur.p) {
			return unique, false
		}
		if visit != nil {
			ms := g.toMS(cur.p)
			if !visit[ms] {
				visit[ms] = true
				unique++
			}
		}
	}
	return 0, true
}

func searchLoops(g grid, w pointSet) int {
	obCount := 0
	for obstPos, visited := range w {
		if !visited {
			continue
		}
		_, loop := runSim(g, nil, g.fromMS(obstPos))
		if loop {
			obCount++
		}
	}
	return obCount
}

type walkDir struct {
	p point
	d direction
}

func (w walkDir) move() walkDir        { return walkDir{w.p.move(w.d), w.d} }
func (w walkDir) rotateRight() walkDir { return walkDir{w.p, w.d.rotateRight()} }

type pointSet []bool
type walkSet []direction

func loadGrid() grid {
	g := grid{}
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '^':
			g.start = point{x, y}
			g.h = len(input) / (g.w + 1)
			return g
		case '.', '#':
			x++
		case '\n':
			g.w = x
			x = 0
			y++
		}
	}
	return g
}

type grid struct {
	start point
	h, w  int
}

func (g grid) obstacle(p point) bool {
	if !g.inBound(p) {
		return false
	}
	return input[p.x+p.y*(g.w+1)] == '#'
}

func (g grid) inBound(p point) bool {
	return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h
}

func (g grid) toMS(p point) int    { return p.x + p.y*(g.w+1) + 1 }
func (g grid) fromMS(ms int) point { return point{(ms - 1) % (g.w + 1), (ms - 1) / (g.w + 1)} }

type point struct{ x, y int }

func (p point) move(dir direction) point {
	switch dir {
	case dirUp:
		return point{p.x, p.y - 1}
	case dirRight:
		return point{p.x + 1, p.y}
	case dirDown:
		return point{p.x, p.y + 1}
	default:
		return point{p.x - 1, p.y}
	}
}

type moveSet []direction

type direction byte

func (d direction) rotateRight() direction {
	if d == dirLeft {
		return dirUp
	}
	return d << 1
}

const (
	dirUp direction = 1 << iota
	dirRight
	dirDown
	dirLeft
)

var benchmark = false
