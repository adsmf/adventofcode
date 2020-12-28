package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
	"github.com/adsmf/adventofcode/utils/hashing/murmur3"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	grid := load("input.txt")
	return grid.part1pairs()
}

func part2() int {
	fullGrid := load("input.txt")
	maxX, maxY := 0, 0
	gapSize := 0
	gapPos := point{-1, -1}
	for pos, data := range fullGrid {
		if data.used == 0 {
			gapSize = data.size
			gapPos = pos
		}
		if pos.x > maxX {
			maxX = pos.x
		}

		if pos.y > maxY {
			maxY = pos.y
		}
	}
	gridState := simpleGridData{
		grid:   map[point]nodeState{},
		maxPos: point{maxX, maxY},
	}
	dataPos := point{maxX, 0}
	for pos, data := range fullGrid {
		if pos == dataPos {
			gridState.grid[pos] = stateTarget
			gridState.targetPos = pos
		} else if pos == gapPos {
			gridState.grid[pos] = stateEmpty
			gridState.gapPos = pos
		} else if data.used <= gapSize {
			gridState.grid[pos] = stateOccupied
		}
	}

	initialHash := gridState.hash()
	openStates := map[simpleGridHash]simpleGridData{initialHash: gridState}
	previousStates := map[simpleGridHash]bool{initialHash: true}

	// fmt.Println(gridState)

	finalTargetPos := point{0, 0}

	for turn := 0; len(openStates) > 0; turn++ {
		nextOpenStates := map[simpleGridHash]simpleGridData{}
		bestDist := 99999
		for _, state := range openStates {
			if state.targetPos == finalTargetPos {
				return turn
			}
			dist := state.targetPos.x + state.targetPos.y
			if dist < bestDist {
				bestDist = dist
			}
			for _, next := range state.next() {
				hash := next.hash()
				if previousStates[hash] {
					continue
				}
				previousStates[hash] = true
				nextOpenStates[hash] = next
			}
		}

		openStates = nextOpenStates
		fmt.Println(len(openStates), bestDist)
	}

	return -1
}

type nodeState int

const (
	stateUnavailable nodeState = iota
	stateEmpty
	stateOccupied
	stateTarget
)

type simpleGridHash uint32

type simpleGridData struct {
	grid      map[point]nodeState
	targetPos point
	gapPos    point
	maxPos    point
}

var hasher = murmur3.NewMurmer3_32(0)

var dirs = []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (g simpleGridData) next() []simpleGridData {
	nextStates := []simpleGridData{}
	gap := g.gapPos
	for _, dir := range dirs {
		nX, nY := gap.x+dir.x, gap.y+dir.y
		nP := point{nX, nY}
		if _, found := g.grid[nP]; !found {
			continue
		}
		nextState := g.copy()
		nextState.grid[gap] = g.grid[nP]
		nextState.grid[nP] = g.grid[gap]
		nextState.gapPos = nP
		if nP == g.targetPos {
			nextState.targetPos = gap
		}
		if nextState.grid[nextState.targetPos] != stateTarget {
			panic("Missed target")
		}
		if nextState.grid[nextState.gapPos] != stateEmpty {
			panic("Mind the gap")
		}
		nextStates = append(nextStates, nextState)
	}
	return nextStates
}

func (g simpleGridData) copy() simpleGridData {
	copy := simpleGridData{
		grid:      make(map[point]nodeState, len(g.grid)),
		targetPos: g.targetPos,
		gapPos:    g.gapPos,
		maxPos:    g.maxPos,
	}
	for pos, state := range g.grid {
		copy.grid[pos] = state
	}
	return copy
}

func (g simpleGridData) hash() simpleGridHash {
	hashData := make([]byte, (g.maxPos.x+1)*(g.maxPos.y)+2)
	hashData[0] = byte(g.targetPos.x)<<4 + byte(g.targetPos.y)
	hashData[1] = byte(g.gapPos.x)<<4 + byte(g.gapPos.y)
	for y := 0; y < g.maxPos.y; y++ {
		for x := 0; x < g.maxPos.x; x++ {
			pos := point{x, y}
			index := y*(g.maxPos.x+1) + x + 2
			hashData[index] = byte(g.grid[pos])
		}
	}
	return simpleGridHash(hasher.HashBytes(hashData))
}

type fullGridData map[point]nodeData

func (g fullGridData) part1pairs() int {
	count := 0
	for _, n1 := range g {
		for _, n2 := range g {
			if n1 == n2 {
				continue
			}
			if n1.used > 0 && (n1.used <= (n2.size - n2.used)) {
				count++
			}
		}
	}
	return count
}

type nodeData struct {
	size int
	used int
}

type point struct{ x, y int }

func load(filename string) fullGridData {
	gd := fullGridData{}
	for _, line := range utils.ReadInputLines(filename)[2:] {
		ints := utils.GetInts(line)
		pos := point{ints[0], ints[1]}
		node := nodeData{
			size: ints[2],
			used: ints[3],
		}
		gd[pos] = node
	}
	return gd
}

var benchmark = false
