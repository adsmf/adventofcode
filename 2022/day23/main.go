package main

import (
	_ "embed"
	"fmt"

	"golang.org/x/exp/constraints"
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
			min, max := pointAt(9999, 9999), pointAt(-9999, -9999)
			for pos := range g {
				min = min.minBound(pos)
				max = max.maxBound(pos)
			}
			area := (max.x() - min.x() + 1) * (max.y() - min.y() + 1)
			p1 = int(area) - len(g)
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
	{pointAt(-1, -1), pointAt(0, -1), pointAt(1, -1)}, // North
	{pointAt(-1, 1), pointAt(0, 1), pointAt(1, 1)},    // South
	{pointAt(-1, -1), pointAt(-1, 0), pointAt(-1, 1)}, // West
	{pointAt(1, -1), pointAt(1, 0), pointAt(1, 1)},    // East
}
var allAround = [...]point{
	pointAt(-1, -1),
	pointAt(0, -1),
	pointAt(1, -1),
	pointAt(-1, 0),
	pointAt(1, 0),
	pointAt(-1, 1),
	pointAt(0, 1),
	pointAt(1, 1),
}

type point uint32

const (
	axisBits = 11
	mask     = (1 << (axisBits + 1)) - 1
)

func pointAt(x, y int16) point {
	return point(x) + 1<<axisBits + (point(y+1<<axisBits) << (axisBits + 2))
}

func (p point) x() int16 { return int16((p & mask) - 1<<axisBits) }
func (p point) y() int16 { return int16((p>>(axisBits+2))&mask - 1<<axisBits) }

func (p point) add(o point) point { return pointAt(p.x()+o.x(), p.y()+o.y()) }
func (p point) minBound(o point) point {
	return pointAt(min(p.x(), o.x()), min(p.y(), o.y()))
}
func (p point) maxBound(o point) point {
	return pointAt(max(p.x(), o.x()), max(p.y(), o.y()))
}

func min[T constraints.Integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}
func max[T constraints.Integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type groveMap map[point]bool

func load() groveMap {
	g := make(groveMap, 75*75)
	x, y := int16(0), int16(0)
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '\n':
			y++
			x = 0
		case '#':
			g[pointAt(x, y)] = true
			x++
		default:
			x++
		}
	}
	return g
}

var benchmark = false
