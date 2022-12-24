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
	visited := map[searchHash]bool{}
	visited[initialState.hash()] = true

	search := func(state searchEntry) {
		hash := state.hash()
		if visited[hash] {
			return
		}
		visited[hash] = true
		nextopen = append(nextopen, state)
	}

	goals := []point{v.end, v.start, v.end}

	doTick := true
	for len(openset) > 0 {

		if doTick {
			v.tick()
		}
		doTick = true
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
					doTick = false
					break
				}
				return p1, state.tick
			}
			nextTick := state.tick + 1
			for _, pos := range []point{
				state.pos,
				state.pos.up(),
				state.pos.down(),
				state.pos.left(),
				state.pos.right(),
			} {
				if !v.isBlizzard(pos) {
					search(searchEntry{pos, nextTick})
				}
			}
		}
		openset, nextopen = nextopen, openset
	}
	return p1, -1
}

type searchHash uint32

type searchEntry struct {
	pos  point
	tick int
}

func (s searchEntry) hash() searchHash {
	return searchHash(s.pos.x) + searchHash(s.pos.y)<<8 + searchHash(s.tick)<<16
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

func (v *valleyData) tick() {
	for y := 0; y < v.height; y++ {
		for i, b := range v.blizzardH[y] {
			next := b.start + b.offset
			if next >= v.width {
				next = 0
			} else if next < 0 {
				next = v.width - 1
			}
			v.blizzardH[y][i].start = next
		}
	}
	for x := 0; x < v.width; x++ {
		for i, b := range v.blizzardV[x] {
			next := b.start + b.offset
			if next >= v.height {
				next = 0
			} else if next < 0 {
				next = v.height - 1
			}
			v.blizzardV[x][i].start = next
		}
	}
}

func (v valleyData) isBlizzard(pos point) bool {
	if pos.x < 0 || pos.x >= v.width || pos.y < 0 || pos.y >= v.height {
		if pos == v.end || pos == v.start {
			return false
		}
		return true
	}
	for _, b := range v.blizzardH[pos.y] {
		if b.start == pos.x {
			return true
		}
	}
	for _, b := range v.blizzardV[pos.x] {
		if b.start == pos.y {
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
