package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

var gridTiles = make([]tileType, 0, 5000)

func part1() int {
	g := grid{
		tiles: gridTiles[0:0],
	}
	rPos, startIdx := g.load(false)
	for i := startIdx; i < len(input); i++ {
		ch := input[i]
		var dir point
		switch ch {
		case '<':
			dir = point{-1, 0}
		case '^':
			dir = point{0, -1}
		case 'v':
			dir = point{0, 1}
		case '>':
			dir = point{1, 0}
		default:
			continue
		}
		next := rPos.add(dir)
		switch g.tile(next) {
		case tileEmpty:
			rPos = next
		case tileWall:
		case tileBox:
			for g.tile(next) == tileBox {
				next = next.add(dir)
			}
			if g.tile(next) == tileEmpty {
				g.setTile(next, tileBox)
				rPos = rPos.add(dir)
				g.setTile(rPos, tileEmpty)
			}
		}
	}
	p1 := 0
	for idx, t := range g.tiles {
		pos := g.fromIndex(idx)
		if t == tileBox {
			p1 += 100*pos.y + pos.x
		}
	}
	return p1
}

func part2() int {
	g := grid{
		tiles: gridTiles[0:0],
	}
	rPos, startIdx := g.load(true)
	open := make([]point, 0, 256)
	next := make([]point, 0, 256)
	toPush := make([]point, 450)
	for i := startIdx; i < len(input); i++ {
		ch := input[i]
		var dir point
		switch ch {
		case '<':
			dir = point{-1, 0}
		case '^':
			dir = point{0, -1}
		case 'v':
			dir = point{0, 1}
		case '>':
			dir = point{1, 0}
		default:
			continue
		}
		vertical := dir.y != 0
		open, next = open[0:0], next[0:0]
		open = append(open, rPos.add(dir))
		toPush = toPush[0:0]
		canMove := true
		for len(open) > 0 {
			for _, cur := range open {
				switch g.tile(cur) {
				case tileWall:
					canMove = false
				case tileBoxL:
					next = append(next, cur.add(dir))
					if vertical {
						next = append(next, cur.add(dir).add(point{1, 0}))
					}
					toPush = append(toPush, cur)
				case tileBoxR:
					next = append(next, cur.add(dir))
					if vertical {
						next = append(next, cur.add(dir).add(point{-1, 0}))
					}
					toPush = append(toPush, cur.add(point{-1, 0}))
				}
				if !canMove {
					break
				}
			}
			if !canMove {
				break
			}
			open, next = next, open[0:0]
		}
		if canMove {
			for _, boxL := range toPush {
				g.setTile(boxL, tileEmpty)
				g.setTile(boxL.add(point{1, 0}), tileEmpty)

			}
			for _, boxL := range toPush {
				g.setTile(boxL.add(dir), tileBoxL)
				g.setTile(boxL.add(dir).add(point{1, 0}), tileBoxR)
			}
			rPos = rPos.add(dir)
		}
	}
	p2 := 0
	for idx, t := range g.tiles {
		pos := g.fromIndex(idx)
		if t == tileBoxL {
			p2 += 100*pos.y + pos.x
		}
	}
	return p2
}

type grid struct {
	w, h  int
	tiles []tileType
}

func (g *grid) load(doubleWidth bool) (point, int) {
	x, y := 0, 0
	charSize := 1
	for ; input[g.w] != '\n'; g.w++ {
	}
	if doubleWidth {
		charSize = 2
		g.w <<= 1
	}
	var rPos point
	for i, ch := range input {
		pos := point{x, y}
		x += charSize
		switch ch {
		case '\n':
			y++
			if x == charSize {
				g.h = y
				return rPos, i
			}
			x = 0
		case 'O':
			if doubleWidth {
				g.tiles = append(g.tiles, tileBoxL)
				g.tiles = append(g.tiles, tileBoxR)
			} else {
				g.tiles = append(g.tiles, tileBox)
			}
		case '#':
			g.tiles = append(g.tiles, tileWall)
			if doubleWidth {
				g.tiles = append(g.tiles, tileWall)
			}
		case '@':
			rPos = pos
			fallthrough
		case '.':
			g.tiles = append(g.tiles, tileEmpty)
			if doubleWidth {
				g.tiles = append(g.tiles, tileEmpty)
			}
		}
	}
	return rPos, -1
}

func (g grid) tile(p point) tileType       { return g.tiles[g.index(p)] }
func (g grid) setTile(p point, t tileType) { g.tiles[g.index(p)] = t }
func (g grid) index(p point) int           { return p.x + p.y*g.w }
func (g grid) fromIndex(idx int) point     { return point{idx % g.w, idx / g.w} }

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }

const (
	tileWall tileType = iota
	tileBox
	tileBoxL
	tileBoxR
	tileEmpty
)

type tileType byte

var benchmark = false
