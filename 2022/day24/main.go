package main

import (
	_ "embed"
	"fmt"
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
	v := load()
	p1 := 0
	initialState := searchEntry{
		pos:  point{0, 0},
		tick: 0,
	}
	const maxOpen = 1050
	openset, nextopen := &[maxOpen]searchEntry{initialState}, &[maxOpen]searchEntry{}
	openCount, nextCount := 1, 0
	visited := [1 << searchHashBits]bool{}
	visited[initialState.hash()] = true

	search := func(state searchEntry) {
		hash := state.hash()
		if visited[hash] {
			return
		}
		visited[hash] = true
		nextopen[nextCount] = state
		nextCount++
	}

	goals := []point{v.end, v.start, v.end}

	doTick := true
	for openCount > 0 {

		if doTick {
			v.tick()
		}
		doTick = true
		nextCount = 0
		for openIdx := 0; openIdx < openCount; openIdx++ {
			state := openset[openIdx]
			if state.pos == goals[0] {
				if p1 == 0 {
					p1 = state.tick
				}
				if len(goals) > 1 {
					goals = goals[1:]
					nextopen[0] = state
					nextCount = 1
					visited = [1 << searchHashBits]bool{}
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
		openCount = nextCount
	}
	return p1, -1
}

type searchHash uint32

type searchEntry struct {
	pos  point
	tick int
}

func (s searchEntry) hash() searchHash {
	hash := fnvo32
	hash ^= fnvHash(s.pos.x)
	hash *= fnvp32
	hash ^= fnvHash(s.pos.y)
	hash *= fnvp32
	hash ^= fnvHash(s.tick)
	hash *= fnvp32
	hash = (hash >> searchHashBits) ^ (hash & ((1 << searchHashBits) - 1))
	return searchHash(hash)
}

const (
	fnvp32 fnvHash = 0x01000193
	fnvo32 fnvHash = 0x811c9dc5
)
const searchHashBits = 22

type fnvHash uint32

type point struct{ x, y int }

func (p point) up() point    { return point{p.x, p.y - 1} }
func (p point) down() point  { return point{p.x, p.y + 1} }
func (p point) left() point  { return point{p.x - 1, p.y} }
func (p point) right() point { return point{p.x + 1, p.y} }

type blizList struct {
	data [55]blizzardInfo
	len  int
}

type valleyData struct {
	blizzardH     [35]blizList
	blizzardV     [100]blizList
	width, height int
	start, end    point
}

func (v *valleyData) tick() {
	for y := 0; y < v.height; y++ {
		for i := 0; i < v.blizzardH[y].len; i++ {
			b := v.blizzardH[y].data[i]
			next := b.start + b.offset
			if next >= v.width {
				next = 0
			} else if next < 0 {
				next = v.width - 1
			}
			v.blizzardH[y].data[i].start = next
		}
	}
	for x := 0; x < v.width; x++ {
		for i := 0; i < v.blizzardV[x].len; i++ {
			b := v.blizzardV[x].data[i]
			next := b.start + b.offset
			if next >= v.height {
				next = 0
			} else if next < 0 {
				next = v.height - 1
			}
			v.blizzardV[x].data[i].start = next
		}
	}
}

func (v *valleyData) isBlizzard(pos point) bool {
	if pos.x < 0 || pos.x >= v.width || pos.y < 0 || pos.y >= v.height {
		if pos == v.end || pos == v.start {
			return false
		}
		return true
	}
	for i := 0; i < v.blizzardH[pos.y].len; i++ {
		b := v.blizzardH[pos.y].data[i]
		if b.start == pos.x {
			return true
		}
	}
	for i := 0; i < v.blizzardV[pos.x].len; i++ {
		b := v.blizzardV[pos.x].data[i]
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
	xMax := 0
	x, y := 0, 0
	for _, ch := range input {
		switch ch {
		case '>':
			v.blizzardH[y-1].data[v.blizzardH[y-1].len] = blizzardInfo{x - 1, 1}
			v.blizzardH[y-1].len++
			x++
		case '<':
			v.blizzardH[y-1].data[v.blizzardH[y-1].len] = blizzardInfo{x - 1, -1}
			v.blizzardH[y-1].len++
			x++
		case '^':
			v.blizzardV[x-1].data[v.blizzardV[x-1].len] = blizzardInfo{y - 1, -1}
			v.blizzardV[x-1].len++
			x++
		case 'v':
			v.blizzardV[x-1].data[v.blizzardV[x-1].len] = blizzardInfo{y - 1, 1}
			v.blizzardV[x-1].len++
			x++
		case '\n':
			x = 0
			y++
		default:
			x++
		}
		if x > xMax {
			xMax = x
		}
	}
	v.height = y - 2
	v.width = xMax - 2
	v.start = point{0, -1}
	v.end = point{v.width - 1, v.height}
	if v.height != 35 || v.width != 100 || (v.start != point{0, -1}) || (v.end != point{99, 35}) {
		panic(fmt.Sprint(v.height, v.width, v.start, v.end))
	}
	return v
}

var benchmark = false
