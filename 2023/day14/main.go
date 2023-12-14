package main

import (
	_ "embed"
	"fmt"
	"hash/fnv"

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
	p1score := 0
	m := load()
	scores := make([]int, 0, 1000)
	seen := make(map[uint64]int, 1000)
	target := 1000000000
	for i := 0; i < target; i++ {
		m.tilt(dirNorth)
		if i == 0 {
			p1score = m.score()
		}
		m.tilt(dirWest)
		m.tilt(dirSouth)
		m.tilt(dirEast)
		h := m.hash()
		if prev, found := seen[h]; found {
			return p1score, scores[((target-prev)%(i-prev))+prev-1]
		}
		scores = append(scores, m.score())
		seen[h] = i
	}
	return p1score, -1
}

func load() mapInfo {
	gm := mapInfo{}
	utils.EachLine(input, func(y int, line string) (done bool) {
		gm.max.y = y + 1
		gm.max.x = len(line)
		for _, ch := range line {
			gm.items = append(gm.items, itemType(ch))
		}
		return false
	})
	return gm
}

type mapInfo struct {
	items []itemType
	max   point
}

func (m mapInfo) score() int {
	score := 0
	for pos, item := range m.items {
		if item != itemRoundRock {
			continue
		}
		score += m.max.y - pos/m.max.x
	}
	return score
}

func (m mapInfo) hash() uint64 {
	hash := fnv.New64a()
	hash.Write(m.items)
	return hash.Sum64()
}

func (m mapInfo) point(offset int, dir direction) (int, bool) {
	switch dir {
	case dirNorth:
		next := offset - m.max.x
		return next, next > 0
	case dirSouth:
		next := offset + m.max.x
		return next, next < m.max.x*m.max.y
	case dirEast:
		return offset + 1, (offset+1)%m.max.x != 0
	case dirWest:
		return offset - 1, (offset)%m.max.x != 0
	}
	return -1, false
}

func (m *mapInfo) tilt(dir direction) {
	for start, item := range m.items {
		switch item {
		case itemCubeRock, itemSpace:
			continue
		}
		lastGood := start
		valid := true
		for tryPos := start; valid; tryPos, valid = m.point(tryPos, dir) {
			newItem := m.items[tryPos]
			if newItem == itemNone || newItem == itemCubeRock {
				break
			}
			if newItem == itemRoundRock {
				continue
			}
			lastGood = tryPos
		}
		if lastGood != start {
			m.items[start] = itemSpace
			m.items[lastGood] = itemRoundRock
		}
	}
}

type point struct {
	x, y int
}

func (p point) add(a point) point { return point{x: p.x + a.x, y: p.y + a.y} }

type direction int

const (
	dirNorth direction = iota
	dirWest
	dirSouth
	dirEast
)

type itemType = byte

const (
	itemNone      itemType = 0
	itemCubeRock  itemType = '#'
	itemRoundRock itemType = 'O'
	itemSpace     itemType = '.'
)

var benchmark = false
