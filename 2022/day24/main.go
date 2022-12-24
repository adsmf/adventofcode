package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	v := load()
	p1, p2 := solve(v)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(v valleyData) (int, int) {
	p1 := 0
	initialState := searchEntry{
		pos:  point{0, 0},
		tick: 0,
	}
	openset, nextopen := []searchEntry{initialState}, []searchEntry{}
	visited := map[searchEntry]bool{initialState: true}

	search := func(state searchEntry) {
		if visited[state] {
			return
		}
		visited[state] = true
		nextopen = append(nextopen, state)
	}

	goals := []point{v.end, v.start, v.end}

	for len(openset) > 0 {
		nextopen = nextopen[0:0]
		for _, state := range openset {
			if state.pos == goals[0] {
				if p1 == 0 {
					p1 = state.tick
				}
				if len(goals) > 1 {
					goals = goals[1:]
					nextopen = nextopen[0:0]
					nextopen = append(nextopen, state)
					for s := range visited {
						delete(visited, s)
					}
					break
				}
				return p1, state.tick
			}
			nextTick := state.tick + 1
			if !v.isBlizzard(state.pos, nextTick) {
				search(searchEntry{state.pos, nextTick})
			}
			if !v.isBlizzard(state.pos.up(), nextTick) {
				search(searchEntry{state.pos.up(), nextTick})
			}
			if !v.isBlizzard(state.pos.down(), nextTick) {
				search(searchEntry{state.pos.down(), nextTick})
			}
			if !v.isBlizzard(state.pos.left(), nextTick) {
				search(searchEntry{state.pos.left(), nextTick})
			}
			if !v.isBlizzard(state.pos.right(), nextTick) {
				search(searchEntry{state.pos.right(), nextTick})
			}
		}
		openset, nextopen = nextopen, openset
	}
	return p1, -1
}

type searchEntry struct {
	pos  point
	tick int
}

type point struct{ x, y int }

func (p point) up() point    { return point{p.x, p.y - 1} }
func (p point) down() point  { return point{p.x, p.y + 1} }
func (p point) left() point  { return point{p.x - 1, p.y} }
func (p point) right() point { return point{p.x + 1, p.y} }

type valleyData struct {
	blizzardH     [35][]blizzardInfo
	blizzardV     [100][]blizzardInfo
	width, height int
	start, end    point
}

func (v valleyData) isBlizzard(pos point, tick int) bool {
	if pos.x < 0 || pos.x >= v.width || pos.y < 0 || pos.y >= v.height {
		if pos == v.end || pos == v.start {
			return false
		}
		return true
	}
	for _, b := range v.blizzardH[pos.y] {
		bPos := (b.start + b.offset*tick) % v.width
		if bPos < 0 {
			bPos += v.width
		}
		if bPos == pos.x {
			return true
		}
	}
	for _, b := range v.blizzardV[pos.x] {
		bPos := (b.start + b.offset*tick) % v.height
		if bPos < 0 {
			bPos += v.height
		}
		if bPos == pos.y {
			return true
		}
	}
	return false
}

type blizzardInfo struct {
	start  int
	offset int
}

func load() valleyData {
	v := valleyData{}
	xMax, yMax := 0, 0
	for y, line := range utils.GetLines(input) {
		if y > yMax {
			yMax = y
		}
		for x, ch := range line {
			if x == 0 {
				continue
			}
			if x > xMax {
				xMax = x
			}
			switch ch {
			case '>':
				v.blizzardH[y-1] = append(v.blizzardH[y-1], blizzardInfo{x - 1, 1})
			case '<':
				v.blizzardH[y-1] = append(v.blizzardH[y-1], blizzardInfo{x - 1, -1})
			case '^':
				v.blizzardV[x-1] = append(v.blizzardV[x-1], blizzardInfo{y - 1, -1})
			case 'v':
				v.blizzardV[x-1] = append(v.blizzardV[x-1], blizzardInfo{y - 1, 1})
			}
		}
	}
	v.height = yMax - 1
	v.width = xMax - 1
	v.start = point{0, -1}
	v.end = point{v.width - 1, v.height}
	return v
}

var benchmark = false
