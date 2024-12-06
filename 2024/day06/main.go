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
	return countUnique(w), searchLoops(g, w)
}

func runSim(g grid, addObst ...point) (walkSet, bool) {
	facing := dirUp
	curPos := g.start
	w := walkSet{walkDir{curPos, facing}: true}
	simObst := map[point]bool{}
	for _, obst := range addObst {
		simObst[obst] = true
	}
	for {
		nextPos := curPos.move(facing)
		for simObst[nextPos] || g.obstacles[nextPos] {
			facing = facing.rotateRight()
			nextPos = curPos.move(facing)
		}
		curPos = nextPos
		if _, found := g.obstacles[curPos]; !found {
			return w, false
		}
		wd := walkDir{curPos, facing}
		if w[wd] {
			return w, true
		}
		w[wd] = true
	}
}

func countUnique(w walkSet) int {
	unique := map[point]bool{}
	for v := range w {
		unique[v.p] = true
	}
	return len(unique)
}

func searchLoops(g grid, w walkSet) int {
	obst := map[point]bool{}
	for v := range w {
		obstPos := v.p
		if _, found := obst[obstPos]; found || g.obstacles[obstPos] || obstPos == g.start {
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

type walkSet map[walkDir]bool

func loadGrid() grid {
	g := grid{
		obstacles: map[point]bool{},
	}
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '^':
			g.start = point{x, y}
			fallthrough
		case '.':
			g.obstacles[point{x, y}] = false
			x++
		case '#':
			g.obstacles[point{x, y}] = true
			x++
		case '\n':
			g.w = x
			x = 0
			y++
		}
	}
	g.h = y + 1
	return g
}

type grid struct {
	obstacles map[point]bool
	start     point
	h, w      int
}
type point struct{ x, y int }

var moves = []point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func (p point) move(dir direction) point {
	return point{p.x + moves[dir].x, p.y + moves[dir].y}
}

type direction int

func (d direction) rotateRight() direction {
	if d == dirLeft {
		return dirUp
	}
	return d + 1
}

const (
	dirUp direction = iota
	dirRight
	dirDown
	dirLeft
)

var benchmark = false
