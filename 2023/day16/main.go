package main

import (
	_ "embed"
	"fmt"
	"math/bits"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	m := load()
	p1, p2 := m.maxEnergised()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type mapInfo struct {
	tiles    []tileType
	max      point
	open     []searchItem
	nextOpen []searchItem
	seen     []direction
}

func (m mapInfo) maxEnergised() (int, int) {
	p1, best := 0, 0
	for x := 0; x < m.max.x; x++ {
		down := m.countEnergised(searchItem{point{x, 0}, dirDown})
		up := m.countEnergised(searchItem{point{x, m.max.y - 1}, dirUp})
		if down > best {
			best = down
		}
		if up > best {
			best = up
		}
	}
	for y := 0; y < m.max.y; y++ {
		right := m.countEnergised(searchItem{point{0, y}, dirRight})
		left := m.countEnergised(searchItem{point{m.max.x - 1, y}, dirLeft})
		if y == 0 {
			p1 = right
		}
		if right > best {
			best = right
		}
		if left > best {
			best = left
		}
	}
	return p1, best
}

func (m *mapInfo) countEnergised(initial searchItem) int {
	if m.seen == nil {
		m.seen = make([]direction, len(m.tiles))
	} else {
		clear(m.seen)
	}
	m.seen[m.pointIndex(initial.pos)] = initial.dir
	m.open = append(m.open[0:0], initial)
	m.nextOpen = m.nextOpen[0:0]
	addSearch := func(curPos point, dir direction) {
		next := searchItem{
			m.pointNext(curPos, dir),
			dir,
		}
		posIdx := m.pointIndex(next.pos)
		if m.tile(next.pos) != tileNone && m.seen[posIdx]&dir == 0 {
			m.seen[posIdx] |= dir
			m.nextOpen = append(m.nextOpen, next)
		}
	}
	for len(m.open) > 0 {
		for _, cur := range m.open {
			tile := m.tile(cur.pos)
			outDirs := emitterMap[tile][cur.dir]
			if bits.OnesCount(uint(outDirs)) == 1 {
				addSearch(cur.pos, outDirs)
				continue
			}
			for dir := dirUp; dir < dirMAX; dir <<= 1 {
				if outDirs&dir > 0 {
					addSearch(cur.pos, dir)
				}
			}
		}
		m.nextOpen, m.open = m.open[0:0], m.nextOpen
	}
	pointsEnergised := 0
	for _, dirs := range m.seen {
		if dirs > 0 {
			pointsEnergised++
		}
	}
	return pointsEnergised
}

func (m mapInfo) tile(pos point) tileType {
	if pos.x < 0 || pos.y < 0 || pos.x >= m.max.x || pos.y >= m.max.y {
		return tileNone
	}
	return m.tiles[m.pointIndex(pos)]
}
func (m mapInfo) pointIndex(pos point) int {
	return pos.x + pos.y*(m.max.x)
}
func (m mapInfo) pointNext(pos point, dir direction) point {
	return pos.add(directionOffsets[dir])
}

type searchItem struct {
	pos point
	dir direction
}

func load() mapInfo {
	gm := mapInfo{}
	utils.EachLine(input, func(y int, line string) (done bool) {
		gm.max.y = y + 1
		gm.max.x = len(line)
		for _, ch := range line {
			gm.tiles = append(gm.tiles, tileSym[ch])
		}
		return false
	})
	return gm
}

var (
	emitterMap = [tileMAX][dirMAX]direction{
		tileSpace:      {dirUp: dirUp, dirDown: dirDown, dirLeft: dirLeft, dirRight: dirRight},
		tileMirrorNWSE: {dirUp: dirLeft, dirDown: dirRight, dirLeft: dirUp, dirRight: dirDown},
		tileMirrorNESW: {dirUp: dirRight, dirDown: dirLeft, dirLeft: dirDown, dirRight: dirUp},
		tileSplitH: {
			dirLeft:  dirLeft,
			dirRight: dirRight,
			dirUp:    dirLeft | dirRight,
			dirDown:  dirLeft | dirRight,
		},
		tileSplitV: {
			dirUp:    dirUp,
			dirDown:  dirDown,
			dirLeft:  dirUp | dirDown,
			dirRight: dirUp | dirDown,
		},
	}
	directionOffsets = [dirMAX]point{
		dirUp:    {0, -1},
		dirDown:  {0, 1},
		dirLeft:  {-1, 0},
		dirRight: {1, 0},
	}
)

type point struct {
	x, y int
}

func (p point) add(a point) point { return point{x: p.x + a.x, y: p.y + a.y} }

type direction int

const (
	dirUp direction = 1 << iota
	dirRight
	dirDown
	dirLeft

	dirMAX
	dirMask = dirMAX - 1
)

type tileType byte

var tileSym = [...]tileType{'.': tileSpace, '|': tileSplitV, '-': tileSplitH, '\\': tileMirrorNWSE, '/': tileMirrorNESW}

const (
	tileNone tileType = iota
	tileSpace
	tileSplitV
	tileSplitH
	tileMirrorNWSE
	tileMirrorNESW

	tileMAX
)

var benchmark = false
