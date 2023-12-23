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
	gm := load()
	em := gm.calcEdges(tileDirections)
	em.optimaze()
	ps := make(pointSet, gm.max.x*gm.max.y)
	return findLongest(em, point{1, 0}, gm.max, ps, 0, 0)
}

func part2() int {
	gm := load()
	em := gm.calcEdges(tileDirections2)
	em.optimaze()
	ps := make(pointSet, gm.max.x*gm.max.y)
	return findLongest(em, point{1, 0}, gm.max, ps, 0, 0)
}

func findLongest(gm edgeMap, pos point, maxPos point, visited pointSet, maxDist int, dist int) int {
	if pos.y == maxPos.y-1 {
		return max(maxDist, dist)
	}
	index := func(p point) int {
		return p.x + p.y*maxPos.x
	}

	visited[index(pos)] = true
	for n, cost := range gm[pos] {
		if visited[index(n)] {
			continue
		}
		maxDist = findLongest(gm, n, maxPos, visited, maxDist, dist+cost)
	}
	visited[index(pos)] = false
	return maxDist

}

type pointSet []bool

func load() mapInfo {
	gm := mapInfo{}
	utils.EachLine(input, func(y int, line string) (done bool) {
		gm.max.y = y + 1
		gm.max.x = len(line)
		for _, ch := range line {
			t := tileSym[ch]
			gm.tiles = append(gm.tiles, t)
		}
		return false
	})
	return gm
}

type edgeMap map[point]map[point]int

func (em *edgeMap) optimaze() {
	for cPos, edges := range *em {
		switch len(edges) {
		case 2:
			n := [2]point{}
			l := [2]int{}
			i := 0
			for pos, pL := range edges {
				n[i] = pos
				l[i] = pL
				i++
			}
			n0, n1 := (*em)[n[0]], (*em)[n[1]]
			n0d, n1d := n0[cPos], n1[cPos]
			if n0d == 0 && n1d == 0 {
				continue
			}
			n0[n[1]] = n0[cPos] + l[1]
			n1[n[0]] = n1[cPos] + l[0]
			delete(n0, cPos)
			delete(n1, cPos)
			delete(*em, cPos)
		}
	}
}

type mapInfo struct {
	tiles []tileType
	max   point
	start point
}

func (m mapInfo) calcEdges(dirs [][]point) edgeMap {
	em := edgeMap{}
	visited := map[point]bool{}
	open := []point{{1, 0}}
	nextOpen := []point{}
	for len(open) > 0 {
		for _, pos := range open {
			cTile := m.tile(pos)

			for _, offset := range dirs[cTile] {
				n := pos.add(offset)
				nTile := m.tile(n)
				if nTile == tileNone || nTile == tileWall {
					continue
				}
				if em[pos] == nil {
					em[pos] = map[point]int{n: 1}
				} else {
					em[pos][n] = 1
				}
				if visited[n] {
					continue
				}
				visited[n] = true
				nextOpen = append(nextOpen, n)
			}
		}
		open, nextOpen = nextOpen, open[0:0]
	}
	return em
}

func (m mapInfo) tile(pos point) tileType {
	for pos.x < 0 || pos.x >= m.max.x || pos.y < 0 || pos.y >= m.max.y {
		return tileNone
	}
	index := m.pointIndex(pos)
	return m.tiles[index]
}
func (m mapInfo) pointIndex(pos point) int {
	return pos.x + pos.y*(m.max.x)
}

type point struct{ x, y int }

func (p point) add(a point) point { return point{x: p.x + a.x, y: p.y + a.y} }

type direction int

const (
	dirUp direction = iota
	dirRight
	dirDown
	dirLeft

	dirMAX
	dirNone = -1
)

var directionOffsets = [dirMAX]point{
	dirUp:    {0, -1},
	dirRight: {1, 0},
	dirDown:  {0, 1},
	dirLeft:  {-1, 0},
}

var (
	tileDirectionsAll = []point{
		directionOffsets[dirUp],
		directionOffsets[dirDown],
		directionOffsets[dirLeft],
		directionOffsets[dirRight],
	}

	tileDirections = [][]point{
		tileSpace:      tileDirectionsAll,
		tileSlopeUp:    {directionOffsets[dirUp]},
		tileSlopeDown:  {directionOffsets[dirDown]},
		tileSlopeLeft:  {directionOffsets[dirLeft]},
		tileSlopeRight: {directionOffsets[dirRight]},
	}

	tileDirections2 = [][]point{
		tileSpace:      tileDirectionsAll,
		tileSlopeUp:    tileDirectionsAll,
		tileSlopeDown:  tileDirectionsAll,
		tileSlopeLeft:  tileDirectionsAll,
		tileSlopeRight: tileDirectionsAll,
	}
)

type tileType rune

var tileSym = [...]tileType{
	'.': tileSpace,
	'#': tileWall,
	'^': tileSlopeUp,
	'v': tileSlopeDown,
	'<': tileSlopeLeft,
	'>': tileSlopeRight,
}

const (
	tileNone tileType = iota
	tileSpace
	tileWall
	tileSlopeUp
	tileSlopeDown
	tileSlopeLeft
	tileSlopeRight

	tileMAX
)

var benchmark = false
