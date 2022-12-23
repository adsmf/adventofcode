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

func checkAll(g groveMap, elf point) byte {
	result := byte(0)
	for i := 1; i < 1<<8; i <<= 1 {
		pos := surrounding[i]
		if g[elf.add(pos)] {
			result |= byte(i)
		}
	}
	return result
}

func scatter(g groveMap, startChoice int) bool {
	targets := map[point][]point{}
	for elf := range g {
		neighbours := checkAll(g, elf)
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
			moveTo := elf.add(moveDir[choice])
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
