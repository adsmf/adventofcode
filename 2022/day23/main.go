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
		moved := 0
		g, moved = scatter(g, i%4)
		if moved == 0 {
			return p1, i + 1
		}
		if i == 9 {
			min, max := pointAt(axisMax, axisMax), pointAt(axisMin, axisMin)
			for elfID := 0; elfID < g.numElves; elfID++ {
				elfPos := g.elves[elfID]
				min = min.minBound(elfPos)
				max = max.maxBound(elfPos)
			}
			area := (max.x() - min.x() + 1) * (max.y() - min.y() + 1)
			p1 = int(area) - g.numElves
		}
	}
}

func checkAll(g *groveMap, elf point) byte {
	result := byte(0)
	for i := 1; i < 1<<8; i <<= 1 {
		pos := surrounding[i]
		if g.occupied[elf.add(pos)] {
			result |= byte(i)
		}
	}
	return result
}

func scatter(g groveMap, startChoice int) (groveMap, int) {
	targets := [gridSize]byte{}
	newPositions := elfList{}
	for elfID := 0; elfID < g.numElves; elfID++ {
		elfPos := g.elves[elfID]
		neighbours := checkAll(&g, elfPos)
		if neighbours == 0 {
			continue
		}
		for i := 0; i < 4; i++ {
			choice := (i + startChoice)
			if choice >= 4 {
				choice -= 4
			}
			if neighbours&choices[choice] > 0 {
				continue
			}
			moveTo := elfPos.add(moveDir[choice])
			targets[moveTo]++
			newPositions[elfID] = moveTo
			break
		}
	}
	moved := 0
	for elfID := 0; elfID < g.numElves; elfID++ {
		fromPos := g.elves[elfID]
		toPos := newPositions[elfID]
		if targets[toPos] == 1 {
			g.occupied[fromPos] = false
			g.occupied[toPos] = true
			g.elves[elfID] = toPos
			moved++
		}
	}
	return g, moved
}

var choices = [...]byte{
	dirNW + dirN + dirNE, // North
	dirSW + dirS + dirSE, // South
	dirNW + dirW + dirSW, // West
	dirNE + dirE + dirSE, // East
}

type direction = byte

const (
	dirNW direction = 1 << iota
	dirN
	dirNE
	dirW
	dirE
	dirSW
	dirS
	dirSE
)

var moveDir = [...]point{surrounding[dirN], surrounding[dirS], surrounding[dirW], surrounding[dirE]}
var surrounding = [...]point{
	dirNW: pointAt(-1, -1), dirN: pointAt(0, -1), dirNE: pointAt(1, -1),
	dirW: pointAt(-1, 0) /*      Centre      */, dirE: pointAt(1, 0),
	dirSW: pointAt(-1, 1), dirS: pointAt(0, 1), dirSE: pointAt(1, 1),
}

type point uint16

const (
	axisBits = 8
	offset   = 1 << (axisBits - 1)
	mask     = (1 << axisBits) - 1
	axisMax  = (1 << (axisBits - 1)) - 1
	axisMin  = (1 << (axisBits - 1)) * -1
	gridSize = 1 << (axisBits * 2)
)

func pointAt(x, y int) point {
	return ((point(x+offset) & mask) | ((point(y+offset) & mask) << axisBits))
}

func (p point) x() int { return int((p & mask)) - offset }
func (p point) y() int { return int((p>>axisBits)&mask) - offset }

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

type elfList [3000]point
type groveMap struct {
	occupied [gridSize]bool
	elves    elfList
	numElves int
}

func load() groveMap {
	g := groveMap{}
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch input[pos] {
		case '\n':
			y++
			x = 0
		case '#':
			pos := pointAt(x, y)
			g.occupied[pos] = true
			g.elves[g.numElves] = pos
			g.numElves++
			x++
		default:
			x++
		}
	}
	return g
}

var benchmark = false
