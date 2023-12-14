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
			p1score, _ = m.eval()
		}
		m.tilt(dirWest)
		m.tilt(dirSouth)
		m.tilt(dirEast)
		score, h := m.eval()
		if prev, found := seen[h]; found {
			return p1score, scores[((target-prev)%(i-prev))+prev-1]
		}
		scores = append(scores, score)
		seen[h] = i
	}
	return p1score, -1
}

func load() mapInfo {
	gm := mapInfo{items: map[point]itemType{}}
	utils.EachLine(input, func(y int, line string) (done bool) {
		gm.maxScore = y + 1
		for x, ch := range line {
			pos := point{x: x, y: y}
			gm.items[pos] = itemType(ch)
		}
		return false
	})
	return gm
}

type mapInfo struct {
	items    map[point]itemType
	maxScore int
}

func (m mapInfo) eval() (int, uint64) {
	score := 0
	hash := fnv.New64a()
	for pos, item := range m.items {
		if item != itemRoundRock {
			continue
		}
		score += m.maxScore - pos.y
		hash.Write([]byte{byte(pos.x), byte(pos.y)})
	}
	return score, hash.Sum64()
}

func (m *mapInfo) tilt(dir direction) {
	tiltDir := directions[dir]
	for start, item := range m.items {
		switch item {
		case itemCubeRock, itemSpace:
			continue
		}
		lastGood := start
		for tryPos := start; ; tryPos = tryPos.add(tiltDir) {
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

var directions = map[direction]point{
	dirNorth: {0, -1},
	dirEast:  {1, 0},
	dirSouth: {0, 1},
	dirWest:  {-1, 0},
}

type itemType byte

const (
	itemNone      itemType = 0
	itemCubeRock  itemType = '#'
	itemRoundRock itemType = 'O'
	itemSpace     itemType = '.'
)

var benchmark = false
