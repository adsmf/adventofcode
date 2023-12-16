package main

import (
	_ "embed"
	"fmt"
	"math/bits"
	"sync"

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
	tiles []tileType
	max   point
}

func (m mapInfo) maxEnergised() (int, int) {
	p1, best := 0, 0

	l := sync.Mutex{}
	record := func(result int) {
		l.Lock()
		if result > best {
			best = result
		}
		l.Unlock()
	}

	queue := make(chan searchItem, 100)
	wg := sync.WaitGroup{}
	for id := 0; id < 100; id++ {
		wg.Add(1)
		go func() {
			for s := range queue {
				record(m.countEnergised(s))
			}
			wg.Done()
		}()
	}

	for x := 0; x < m.max.x; x++ {
		queue <- searchItem{point{x, 0}, dirDown}
		queue <- searchItem{point{x, m.max.y - 1}, dirUp}
	}
	for y := 1; y < m.max.y; y++ {
		queue <- searchItem{point{0, y}, dirRight}
		queue <- searchItem{point{m.max.x - 1, y}, dirLeft}
	}
	p1 = m.countEnergised(searchItem{point{0, 0}, dirRight})
	record(p1)
	close(queue)
	wg.Wait()
	return p1, best
}

func (m *mapInfo) countEnergised(initial searchItem) int {
	seen := make([]direction, len(m.tiles))
	seen[m.pointIndex(initial.pos)] = initial.dir
	open := []searchItem{initial}
	nextOpen := []searchItem{}
	addSearch := func(curPos point, dir direction) {
		next := searchItem{
			m.pointNext(curPos, dir),
			dir,
		}
		posIdx := m.pointIndex(next.pos)
		if m.tile(next.pos) != tileNone && seen[posIdx]&dir == 0 {
			seen[posIdx] |= dir
			nextOpen = append(nextOpen, next)
		}
	}
	for len(open) > 0 {
		for _, cur := range open {
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
		nextOpen, open = open[0:0], nextOpen
	}
	pointsEnergised := 0
	for _, dirs := range seen {
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
