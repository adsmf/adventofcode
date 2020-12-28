package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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
	gridSize := (maxX + 1) * (maxY + 1)
	gridState := simpleGridData{
		grid:   make([]nodeState, gridSize),
		width:  maxX + 1,
		height: maxY + 1,
	}
	dataPos := point{maxX, 0}
	for pos, data := range fullGrid {
		ind := pos.x + pos.y*gridState.width
		if pos == dataPos {
			gridState.grid[ind] = stateTarget
			gridState.targetIndex = ind
		} else if pos == gapPos {
			gridState.grid[ind] = stateEmpty
			gridState.gapIndex = ind
		} else if data.used <= gapSize {
			gridState.grid[ind] = stateOccupied
		}
	}

	initialHash := gridState.hash()
	openStates := map[simpleGridHash]simpleGridData{initialHash: gridState}
	previousStates := map[simpleGridHash]bool{initialHash: true}

	finalTargetPos := 0

	for turn := 0; len(openStates) > 0; turn++ {
		nextOpenStates := map[simpleGridHash]simpleGridData{}
		for _, state := range openStates {
			if state.targetIndex == finalTargetPos {
				return turn
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
	}

	return -1
}

type nodeState = byte

const (
	stateUnavailable nodeState = iota
	stateEmpty
	stateOccupied
	stateTarget
)

type simpleGridData struct {
	grid          []nodeState
	targetIndex   int
	gapIndex      int
	width, height int
}

var dirs = []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (g simpleGridData) next() []simpleGridData {
	nextStates := []simpleGridData{}
	gapInd := g.gapIndex
	for _, dir := range dirs {
		nInd := gapInd + dir.y*g.width + dir.x
		if nInd < 0 || nInd >= len(g.grid) {
			continue
		}
		if g.grid[nInd] == stateUnavailable {
			continue
		}
		if (nInd/g.width - gapInd/g.width) != dir.y {
			continue
		}

		nextState := g.copy()
		nextState.grid[gapInd] = g.grid[nInd]
		nextState.grid[nInd] = g.grid[gapInd]
		nextState.gapIndex = nInd
		if nInd == g.targetIndex {
			nextState.targetIndex = gapInd
		}
		nextStates = append(nextStates, nextState)
	}
	return nextStates
}

func (g simpleGridData) copy() simpleGridData {
	clone := simpleGridData{
		grid:        make([]nodeState, len(g.grid)),
		targetIndex: g.targetIndex,
		gapIndex:    g.gapIndex,
		width:       g.width,
		height:      g.height,
	}
	copy(clone.grid, g.grid)
	return clone
}

type simpleGridHash string

func (g simpleGridData) hash() simpleGridHash {
	return simpleGridHash(g.grid)
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
