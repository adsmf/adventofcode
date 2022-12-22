package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	g, cursor := load()
	p1 := part1(g, cursor)
	p2 := part2(g, cursor)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(g groveMap, cursor int) int {
	g = g.walkPath(cursor, false)
	return g.password()
}

func part2(g groveMap, cursor int) int {
	g = g.walkPath(cursor, true)
	return g.password()
}

type groveMap struct {
	tiles [maxX][maxY]tileContent
	loc   location
}

const (
	maxX = 150
	maxY = 200
)

func (g groveMap) password() int {
	return (g.loc.pos.y+1)*1000 + (g.loc.pos.x+1)*4 + int(g.loc.facing)
}

func (g groveMap) walkPath(cursor int, asCube bool) groveMap {
	for cursor < len(input) {
		ch := input[cursor]
		switch {
		case ch >= '0' && ch <= '9':
			steps := 0
			steps, cursor = getInt(input, cursor)
			if asCube {
				g = g.moveCube(steps)
			} else {
				g = g.move(steps)
			}
		case ch == 'L':
			g.loc.facing = g.loc.facing.turnLeft()
			cursor++
		case ch == 'R':
			g.loc.facing = g.loc.facing.turnRight()
			cursor++
		case ch == '\n':
			cursor++
		}
	}
	return g
}

func (g groveMap) moveCube(spaces int) groveMap {
	pos := g.loc.pos
	facing := g.loc.facing
	for step := 0; step < spaces; step++ {
		dir := directions[facing]
		newPos := pos.add(dir)
		newFacing := facing
		switch facing {
		case headingRight:
			switch {
			case newPos.x == 150: // 2->5
				newPos.y = 149 - newPos.y
				newPos.x = 99
				newFacing = headingLeft
			case newPos.x == 100 && newPos.y >= 50 && newPos.y < 100: // 3->2
				newPos.x = newPos.y + 50
				newPos.y = 49
				newFacing = headingUp
			case newPos.x == 100 && newPos.y >= 100 && newPos.y < 150: // 5->2
				newPos.y = 149 - newPos.y
				newPos.x = 149
				newFacing = headingLeft
			case newPos.x == 50 && newPos.y >= 150: // 6->5
				newPos.x = newPos.y - 100
				newPos.y = 149
				newFacing = headingUp
			}
		case headingLeft:
			switch {
			case newPos.x < 50 && newPos.y < 50: // 1->4
				newPos.y = 149 - newPos.y
				newPos.x = 0
				newFacing = headingRight
			case newPos.x < 50 && newPos.y >= 50 && newPos.y < 100: // 3->4
				newPos.x = newPos.y - 50
				newPos.y = 100
				newFacing = headingDown
			case newPos.x < 0 && newPos.y >= 100 && newPos.y < 150: // 4->1
				newPos.y = 149 - newPos.y
				newPos.x = 50
				newFacing = headingRight
			case newPos.x < 0 && newPos.y >= 150 && newPos.y < 200: // 6->1
				newPos.x = newPos.y - 100
				newPos.y = 0
				newFacing = headingDown
			}
		case headingUp:
			switch {
			case newPos.y < 0 && newPos.x < 100: // 1->6
				newPos.y = newPos.x + 100
				newPos.x = 0
				newFacing = headingRight
			case newPos.y < 0 && newPos.x >= 100 && newPos.x < 150: // 2->6
				newPos.x -= 100
				newPos.y = 199
				newFacing = headingUp
			case newPos.y < 100 && newPos.x < 50: // 4->3
				newPos.y = newPos.x + 50
				newPos.x = 50
				newFacing = headingRight
			}
		case headingDown:
			switch {
			case newPos.y >= 200: // 6->2
				newPos.x += 100
				newPos.y = 0
				newFacing = headingDown
			case newPos.y >= 150 && newPos.x >= 50 && newPos.x < 100: // 5->6
				newPos.y = newPos.x + 100
				newPos.x = 49
				newFacing = headingLeft
			case newPos.y >= 50 && newPos.x >= 100: // 2->3
				newPos.y = newPos.x - 50
				newPos.x = 99
				newFacing = headingLeft
			}
		}
		newTile := g.tiles[newPos.x][newPos.y]
		if newTile == tileWall {
			break
		}
		pos = newPos
		facing = newFacing
	}
	g.loc.facing = facing
	g.loc.pos = pos
	return g
}

func (g groveMap) move(spaces int) groveMap {
	dir := directions[g.loc.facing]
	pos := g.loc.pos

	for step := 0; step < spaces; step++ {
		newPos := g.fixBound(pos.add(dir))
		for g.tiles[newPos.x][newPos.y] == tileNothing {
			newPos = g.fixBound(newPos.add(dir))
		}
		newTile := g.tiles[newPos.x][newPos.y]
		if newTile == tileWall {
			break
		}
		pos = newPos
	}
	g.loc.pos = pos
	return g
}

func (g groveMap) fixBound(pos point) point {
	if pos.x >= maxX {
		pos.x = 0
	}
	if pos.x < 0 {
		pos.x = maxX - 1
	}
	if pos.y >= maxY {
		pos.y = 0
	}
	if pos.y < 0 {
		pos.y = maxY - 1
	}
	return pos
}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }

type location struct {
	pos    point
	facing heading
}

type heading int8

const (
	headingRight heading = iota
	headingDown
	headingLeft
	headingUp
)

func (h heading) turnRight() heading {
	h++
	if h > headingUp {
		return headingRight
	}
	return h
}
func (h heading) turnLeft() heading {
	h--
	if h < headingRight {
		return headingUp
	}
	return h
}

type tileContent int

const (
	tileNothing tileContent = iota
	tileOpen
	tileWall
)

func load() (groveMap, int) {
	g := groveMap{}
	cursor := 0

	lastWasNewline := false
	x, y := 0, 0
	for cursor = 0; cursor < len(input); cursor++ {
		ch := input[cursor]
		if lastWasNewline && ch == '\n' {
			break
		}
		switch ch {
		case '\n':
			lastWasNewline = true
			x = 0
			y++
		case '.':
			lastWasNewline = false
			g.tiles[x][y] = tileOpen
			x++
		case '#':
			lastWasNewline = false
			g.tiles[x][y] = tileWall
			x++
		case ' ':
			lastWasNewline = false
			x++
		}
	}
	g = g.move(1)
	return g, cursor + 1
}

var directions = [...]point{
	headingUp:    {0, -1},
	headingDown:  {0, 1},
	headingLeft:  {-1, 0},
	headingRight: {1, 0},
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

var benchmark = false
