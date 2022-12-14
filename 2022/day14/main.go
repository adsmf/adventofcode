package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	g := loadGrid()
	p1, p2 := g.dropSand(point{500, 0})
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func (g *scanGrid) dropSand(source point) (int, int) {
	p1 := 0
	step := 0
	for {
		sandPos := source
		for {
			moved := false
			switch {
			case g.clearAt(sandPos.down()):
				sandPos = sandPos.down()
				moved = true
			case g.clearAt(sandPos.downLeft()):
				sandPos = sandPos.downLeft()
				moved = true
			case g.clearAt(sandPos.downRight()):
				sandPos = sandPos.downRight()
				moved = true
			}
			if !moved {
				if sandPos == source {
					return p1, step + 1
				}
				g.data[sandPos.x][sandPos.y] = tileSand
				break
			}
			if sandPos.y > g.maxBound.y {
				if p1 == 0 {
					p1 = step
				}
				if sandPos.y == g.maxBound.y+1 {
					g.data[sandPos.x][sandPos.y] = tileSand
					break
				}
			}
		}
		step++
	}
}

func (g *scanGrid) clearAt(pos point) bool {
	return g.data[pos.x][pos.y] == tileEmpty
}

func loadGrid() scanGrid {
	g := scanGrid{
		minBound: point{500, 0},
		maxBound: point{500, 0},
	}

	newLine := true
	var newX, newY int
	prevPos := point{}
	for i := 0; i < len(input); {
		newX, i = getInt(input, i)
		newY, i = getInt(input, i+1)
		newPos := point{newX, newY}

		if newLine {
			prevPos = newPos
			newLine = false
			i += 4
			continue
		}

		from := minPoint(prevPos, newPos)
		to := maxPoint(prevPos, newPos)
		for x := from.x; x <= to.x; x++ {
			for y := from.y; y <= to.y; y++ {
				g.data[x][y] = tileRock
			}
		}

		g.minBound = minPoint(g.minBound, point{newX, newY})
		g.maxBound = maxPoint(g.maxBound, point{newX, newY})
		prevPos = newPos
		if input[i] == '\n' {
			newLine = true
			i++
			continue
		}
		i += 4
	}

	return g
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; pos < len(in) && in[pos]&0xf0 == 0x30; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

type point struct{ x, y int }

func (p point) down() point      { return point{p.x, p.y + 1} }
func (p point) downLeft() point  { return point{p.x - 1, p.y + 1} }
func (p point) downRight() point { return point{p.x + 1, p.y + 1} }

func minPoint(p1, p2 point) point {
	min := p1
	if p2.x < p1.x {
		min.x = p2.x
	}
	if p2.y < p1.y {
		min.y = p2.y
	}
	return min
}
func maxPoint(p1, p2 point) point {
	max := p1
	if p2.x > p1.x {
		max.x = p2.x
	}
	if p2.y > p1.y {
		max.y = p2.y
	}
	return max
}

type scanGrid struct {
	data     tileGrid
	minBound point
	maxBound point
}
type tileGrid [1000][200]tileType
type tileType byte

const (
	tileEmpty tileType = iota
	tileRock
	tileSand
)

var benchmark = false
