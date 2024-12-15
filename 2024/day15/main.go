package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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

func part1() int {
	g := grid{
		tiles: map[point]tileType{},
	}
	var rPos point
	utils.EachSectionMB(input, "\n\n", func(index int, section string) (done bool) {
		if index == 0 {
			rPos = g.load(section, false)
			return false
		}
		for _, ch := range section {
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
			switch g.tiles[next] {
			case tileEmpty:
				rPos = next
			case tileWall:
			case tileBox:
				for g.tiles[next] == tileBox {
					next = next.add(dir)
				}
				if g.tiles[next] == tileEmpty {
					g.tiles[next] = tileBox
					rPos = rPos.add(dir)
					g.tiles[rPos] = tileEmpty
				}
			}
		}
		return false
	})
	p1 := 0
	for pos, t := range g.tiles {
		if t == tileBox {
			p1 += 100*pos.y + pos.x
		}
	}
	return p1
}

func part2() int {
	g := grid{
		tiles: map[point]tileType{},
	}
	var rPos point
	utils.EachSectionMB(input, "\n\n", func(index int, section string) (done bool) {
		if index == 0 {
			rPos = g.load(section, true)
			return false
		}
		for _, ch := range section {
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
			open := []point{rPos.add(dir)}
			next := []point{}
			canMove := true
			toPush := []point{}
			for len(open) > 0 {
				for _, cur := range open {
					switch g.tiles[cur] {
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
					g.tiles[boxL] = tileEmpty
					g.tiles[boxL.add(point{1, 0})] = tileEmpty

				}
				for _, boxL := range toPush {
					g.tiles[boxL.add(dir)] = tileBoxL
					g.tiles[boxL.add(dir).add(point{1, 0})] = tileBoxR
				}
				rPos = rPos.add(dir)
			}
		}
		return false
	})
	p2 := 0
	for pos, t := range g.tiles {
		if t == tileBoxL {
			p2 += 100*pos.y + pos.x
		}
	}
	return p2
}

type grid struct {
	w, h  int
	tiles map[point]tileType
}

func (g *grid) load(section string, doubleWidth bool) point {
	x, y := 0, 0
	charSize := 1
	if doubleWidth {
		charSize = 2
	}
	var rPos point
	for _, ch := range section {
		pos := point{x, y}
		posR := point{x + 1, y}
		x += charSize
		switch ch {
		case '\n':
			y++
			g.w = max(g.w, x)
			x = 0
		case 'O':
			if doubleWidth {
				g.tiles[pos] = tileBoxL
				g.tiles[posR] = tileBoxR
			} else {
				g.tiles[pos] = tileBox
			}
		case '#':
			g.tiles[pos] = tileWall
			if doubleWidth {
				g.tiles[posR] = tileWall
			}
		case '@':
			rPos = pos
			fallthrough
		case '.':
			g.tiles[pos] = tileEmpty
			if doubleWidth {
				g.tiles[posR] = tileEmpty
			}
		}
	}
	g.h = y + 1
	return rPos
}

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

func (t tileType) String() string {
	switch t {
	case tileBox:
		return "O"
	case tileBoxL:
		return "["
	case tileBoxR:
		return "]"
	case tileEmpty:
		return "."
	case tileWall:
		return "#"
	default:
		return ""
	}
}

var benchmark = false
