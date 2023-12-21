package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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
	m := load()
	p1 := 0
	open := []point{m.start}
	nextOpen := []point{}
	keyPoints := []point{}
	nextPos := map[point]bool{}
	for i := 1; len(open) > 0; i++ {
		clear(nextPos)
		for _, pos := range open {
			for _, off := range directionOffsets {
				nPos := pos.add(off)
				nTile := m.tile(nPos)
				if nTile == tileSpace || nTile == tileStart {
					nextPos[nPos] = true
				}
			}
		}
		for pos := range nextPos {
			nextOpen = append(nextOpen, pos)
		}
		if i == 64 {
			p1 = len(nextOpen)
		}
		if (i)%m.max.x == 65 {
			keyPoints = append(keyPoints, point{i, len(nextOpen)})
		}
		if len(keyPoints) == 3 {
			break
		}
		open, nextOpen = nextOpen, open[0:0]
	}
	return p1, interpolate(keyPoints, 26501365)
}

func interpolate(points []point, at int) int {
	estimate := float64(0)
	for i := 0; i < len(points); i++ {
		prod := float64(points[i].y)
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			prod = prod * (float64(at) - float64(points[j].x)) / float64(points[i].x-points[j].x)
		}
		estimate += prod
	}
	return int(estimate)
}

type searchEntry struct {
	pos       point
	remaining int
}

type mapInfo struct {
	tiles []tileType
	max   point
	start point
}

func (m mapInfo) tile(pos point) tileType {
	for pos.x < 0 {
		pos.x += m.max.x
	}
	for pos.x >= m.max.x {
		pos.x -= m.max.x
	}
	for pos.y < 0 {
		pos.y += m.max.y
	}
	for pos.y >= m.max.y {
		pos.y -= m.max.y
	}
	return m.tiles[m.pointIndex(pos)]
}
func (m mapInfo) pointIndex(pos point) int {
	return pos.x + pos.y*(m.max.x)
}

func load() mapInfo {
	gm := mapInfo{}
	utils.EachLine(input, func(y int, line string) (done bool) {
		gm.max.y = y + 1
		gm.max.x = len(line)
		for x, ch := range line {
			t := tileSym[ch]
			gm.tiles = append(gm.tiles, t)
			if t == tileStart {
				gm.start = point{x, y}
			}
		}
		return false
	})
	return gm
}

var directionOffsets = [dirMAX]point{
	dirUp:    {0, -1},
	dirDown:  {0, 1},
	dirLeft:  {-1, 0},
	dirRight: {1, 0},
}

type point struct {
	x, y int
}

func (p point) add(a point) point { return point{x: p.x + a.x, y: p.y + a.y} }

type direction int

const (
	dirUp direction = iota
	dirRight
	dirDown
	dirLeft

	dirMAX
)

type tileType rune

var tileSym = [...]tileType{'.': tileSpace, '#': tileWall, 'S': tileStart}

const (
	tileNone tileType = iota
	tileSpace
	tileWall
	tileStart

	tileMAX
)

var benchmark = false
