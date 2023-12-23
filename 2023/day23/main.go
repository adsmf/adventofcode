package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	gm := load()
	p1 := part1(gm)
	p2 := part2(gm)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(gm mapInfo) int {
	segments := gm.calcEdges(tileDirectionsP1)
	flattened, start, end := segments.flatten(gm.max.y - 1)
	return findLongest(flattened, start, end, make([]bool, len(flattened)), 0, 0)
}

func part2(gm mapInfo) int {
	segments := gm.calcEdges(tileDirectionsP2)
	flattened, start, end := segments.flatten(gm.max.y - 1)
	return findLongest(flattened, start, end, make([]bool, len(flattened)), 0, 0)
}

func findLongest(gm flatEdge, posID int, targetID int, visited []bool, maxDist int, dist int) int {
	if posID == targetID {
		return max(maxDist, dist)
	}
	visited[posID] = true
	for _, n := range gm[posID] {
		if !visited[n.end] {
			maxDist = findLongest(gm, n.end, targetID, visited, maxDist, dist+n.length)
		}
	}
	visited[posID] = false
	return maxDist
}

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

type flatEdge [][]edgeInfo

type edgeMap map[point]map[point]int

func (e edgeMap) flatten(finalRow int) (flatEdge, int, int) {
	startID, endID := -1, -1
	pointIDs := make(map[point]int, len(e))
	for pos := range e {
		id := len(pointIDs)
		pointIDs[pos] = id
		if pos.y == 0 {
			startID = id
		}
		if pos.y == finalRow {
			endID = id
		}
	}
	f := make(flatEdge, len(pointIDs))
	for pos, edges := range e {
		idx := pointIDs[pos]
		f[idx] = make([]edgeInfo, 0, len(edges))
		for to, dist := range edges {
			f[idx] = append(f[idx], edgeInfo{pointIDs[to], dist})
		}
	}
	return f, startID, endID
}

type edgeInfo struct {
	end    int
	length int
}

func (em *edgeMap) optimazeEdges() {
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
	em.optimazeEdges()
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

	tileDirectionsP1 = [][]point{
		tileSpace:      tileDirectionsAll,
		tileSlopeUp:    {directionOffsets[dirUp]},
		tileSlopeDown:  {directionOffsets[dirDown]},
		tileSlopeLeft:  {directionOffsets[dirLeft]},
		tileSlopeRight: {directionOffsets[dirRight]},
	}

	tileDirectionsP2 = [][]point{
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
