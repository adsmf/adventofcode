package main

import (
	_ "embed"
	"fmt"
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
	g := loadGrid()
	w, _ := runSim(g)
	return len(w), searchLoops(g, w)
}

func runSim(g grid, addObst ...point) (walkSet, bool) {
	cur := walkDir{g.start, dirUp}
	w := walkSet{cur.p: cur.d}
	simObst := map[point]bool{}
	for _, obst := range addObst {
		simObst[obst] = true
	}
	for {
		next := cur.move()
		for simObst[next.p] || g.obstacle(next.p) {
			cur = cur.rotateRight()
			next = cur.move()
		}
		cur = next
		if !g.inBound(cur.p) {
			return w, false
		}
		if w[cur.p]&cur.d > 0 {
			return w, true
		}
		w[cur.p] |= cur.d
	}
}

func searchLoops(g grid, w walkSet) int {
	obst := map[point]bool{}
	for obstPos := range w {
		if _, found := obst[obstPos]; found || g.obstacle(obstPos) || obstPos == g.start {
			continue
		}
		_, loop := runSim(g, obstPos)
		obst[obstPos] = loop
	}
	count := 0
	for _, loop := range obst {
		if loop {
			count++
		}
	}
	return count
}

type walkDir struct {
	p point
	d direction
}

func (w walkDir) move() walkDir        { return walkDir{w.p.move(w.d), w.d} }
func (w walkDir) rotateRight() walkDir { return walkDir{w.p, w.d.rotateRight()} }

type walkSet map[point]direction

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
	pos := p.x + p.y*(g.w+1)
	if pos >= len(input) {
		panic(fmt.Sprintf("%#v %#v", p, g))
	}
	return input[p.x+p.y*(g.w+1)] == '#'
}

func (g grid) inBound(p point) bool {
	return p.x >= 0 && p.x < g.w && p.y >= 0 && p.y < g.h
}

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
