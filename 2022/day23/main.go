package main

import (
	_ "embed"
	"fmt"
	"strings"
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
	p1 := 0
	g := load()
	for i := 0; ; i++ {
		moved := scatter(g, i%4)
		if !moved {
			return p1, i + 1
		}
		if i == 9 {
			min, max := point{9999, 9999}, point{-9999, -9999}
			for pos := range g {
				min = min.minBound(pos)
				max = max.maxBound(pos)
			}
			area := (max.x - min.x + 1) * (max.y - min.y + 1)
			p1 = area - len(g)
		}
	}
}

func checkDir(g groveMap, elf point, dirs [3]point) bool {
	for _, pos := range dirs {
		if g[elf.add(pos)] {
			return true
		}
	}
	return false
}

func checkAll(g groveMap, elf point) bool {
	for _, pos := range allAround {
		if g[elf.add(pos)] {
			return true
		}
	}
	return false
}

func scatter(g groveMap, startChoice int) bool {
	targets := map[point][]point{}
	for elf := range g {
		if !checkAll(g, elf) {
			continue
		}

		for i := 0; i < 4; i++ {
			choice := (i + startChoice)
			if choice >= 4 {
				choice -= 4
			}
			checkDirs := look[choice]
			found := checkDir(g, elf, checkDirs)
			if found {
				continue
			}
			moveTo := elf.add(checkDirs[1])
			targets[moveTo] = append(targets[moveTo], elf)
			break
		}
	}
	moved := false
	for to, from := range targets {
		if len(from) == 1 {
			delete(g, from[0])
			g[to] = true
			moved = true
		}
	}
	return moved
}

var look = [...][3]point{
	{{-1, -1}, {0, -1}, {1, -1}}, // North
	{{-1, 1}, {0, 1}, {1, 1}},    // South
	{{-1, -1}, {-1, 0}, {-1, 1}}, // West
	{{1, -1}, {1, 0}, {1, 1}},    // East
}
var allAround = [...]point{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }
func (p point) minBound(o point) point {
	return point{min(p.x, o.x), min(p.y, o.y)}
}
func (p point) maxBound(o point) point {
	return point{max(p.x, o.x), max(p.y, o.y)}
}

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

type groveMap map[point]bool

func (g groveMap) String() string {
	min, max := point{999999, 999999}, point{-99999, -99999}
	for pos := range g {
		min = min.minBound(pos)
		max = max.maxBound(pos)
	}
	sb := strings.Builder{}
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if g[point{x, y}] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func load() groveMap {
	g := make(groveMap, 75*75)
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '\n':
			y++
			x = 0
		case '#':
			g[point{x, y}] = true
			x++
		default:
			x++
		}
	}
	return g
}

var benchmark = false
