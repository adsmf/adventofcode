package main

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
	"github.com/adsmf/adventofcode/utils/hashing/fnv"
)

//go:embed input.txt
var input string

func main() {
	initial := load(input)
	p1 := calcEnergy(initial)
	p2 := part2(initial)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func calcEnergy(initial grid) int {
	openStates := make(stateHeap, 1, 22000)
	openStates[0] = initial
	heap.Init(&openStates)
	visited := make(map[gridHash]bool, len(initial.rooms[0])*22000)

	for openStates.Len() > 0 {
		nextInt := heap.Pop(&openStates)
		next := nextInt.(grid)
		hash := next.hash()
		if visited[hash] {
			continue
		}
		visited[hash] = true
		done := true
		if next.clearPath(0, 10, true) {
			for room := range next.rooms {
				if !next.roomCorrect(tileType(room)) {
					done = false
					break
				}
			}
			if done {
				return next.energy
			}
		}
		moves := next.moves()
		for _, move := range moves {
			heap.Push(&openStates, move)
		}
	}
	return -1
}

func part2(initial grid) int {
	initial.rooms[0] = append([]tileType{initial.rooms[0][0]}, tileD, tileD, initial.rooms[0][1])
	initial.rooms[1] = append([]tileType{initial.rooms[1][0]}, tileB, tileC, initial.rooms[1][1])
	initial.rooms[2] = append([]tileType{initial.rooms[2][0]}, tileA, tileB, initial.rooms[2][1])
	initial.rooms[3] = append([]tileType{initial.rooms[3][0]}, tileC, tileA, initial.rooms[3][1])
	return calcEnergy(initial)
}

func load(input string) grid {
	g := grid{
		hallway: make([]tileType, 11),
		rooms:   make([][]tileType, 4),
	}
	for i := 0; i < 11; i++ {
		g.hallway[i] = tileOpen
	}
	for i := 0; i < 4; i++ {
		g.rooms[i] = make([]tileType, 2)
	}
	lines := utils.GetLines(input)
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			g.rooms[j][1-i] = tileType(lines[2+i][3+2*j] - 'A')
		}
	}
	return g
}

type stateHeap []grid

func (s stateHeap) Len() int           { return len(s) }
func (s stateHeap) Less(i, j int) bool { return s[i].energy < s[j].energy }
func (s stateHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s *stateHeap) Pop() interface{} {
	var popped interface{}
	popped, (*s) = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return popped
}
func (s *stateHeap) Push(entry interface{}) {
	(*s) = append((*s), entry.(grid))
}

type grid struct {
	hallway []tileType
	rooms   [][]tileType
	energy  int
}

func (g grid) moves() []grid {
	moves := make([]grid, 0, 3)
	// Check hallway candidates
	for i := 0; i < 11; i++ {
		tile := g.hallway[i]
		if tile != tileOpen {
			if !g.roomCorrect(tile) {
				continue
			}
			start := tileType(i)
			end := (tile + 1) * 2
			if !g.clearPath(tileType(i), tileType(end), false) {
				continue
			}

			_, targetIndex := g.firstOccupant(tile)
			targetIndex++

			option := g.copy()
			option.hallway[i] = tileOpen
			option.rooms[tile][targetIndex] = tile
			option.energy += (utils.IntAbs(int(start-end)) + (len(g.rooms[tile]) - targetIndex)) * costs[tile]

			moves = append(moves, option)
		}
	}
	// Check room occupants
	for i := tileType(0); i < 4; i++ {
		if g.roomCorrect(i) {
			continue
		}
		tile, startIndex := g.firstOccupant(i)
		if g.roomCorrect(tile) && g.pathBetweenRooms(i, tile) {
			// Can jump directly
			hallwayDist := utils.IntAbs(int(i-tile) * 2)

			_, targetIndex := g.firstOccupant(tile)
			targetIndex++
			opt := g.copy()

			opt.rooms[i][startIndex] = tileOpen
			opt.rooms[tile][targetIndex] = tile
			opt.energy += ((len(g.rooms[i]) - startIndex) + (hallwayDist) + (len(g.rooms[tile]) - targetIndex)) * costs[tile]
			moves = append(moves, opt)
		} else {
			// Needs to move to hallway
			start := (i + 1) * 2
			for _, haltAt := range haltablePositions {
				if g.clearPath(start, haltAt, true) {
					opt := g.copy()
					opt.rooms[i][startIndex] = tileOpen
					opt.hallway[haltAt] = tile
					opt.energy += ((len(g.rooms[i]) - startIndex) + utils.IntAbs(int(start-haltAt))) * costs[tile]
					moves = append(moves, opt)
				}
			}
		}
	}
	return moves
}

func (g grid) clearPath(start, end tileType, includeEnds bool) bool {
	if start > end {
		start, end = end, start
	}
	if !includeEnds {
		start++
		end--
	}
	for i := start; i <= end; i++ {
		if g.hallway[i] != tileOpen {
			return false
		}
	}
	return true
}

func (g grid) pathBetweenRooms(room1, room2 tileType) bool {
	start := (room1 + 1) * 2
	end := (room2 + 1) * 2
	return g.clearPath(start, end, false)
}

func (g grid) firstOccupant(room tileType) (tileType, int) {
	for i := len(g.rooms[room]) - 1; i >= 0; i-- {
		if g.rooms[room][i] == tileOpen {
			continue
		}
		return g.rooms[room][i], i
	}
	return tileOpen, -1
}

func (g grid) roomCorrect(room tileType) bool {
	for j := 0; j < len(g.rooms[room]); j++ {
		if g.rooms[room][j] != tileOpen && g.rooms[room][j] != room {
			return false
		}
	}
	return true
}

func (g grid) copy() grid {
	c := grid{
		hallway: make([]tileType, 11),
		rooms:   make([][]tileType, 4),
		energy:  g.energy,
	}
	copy(c.hallway, g.hallway)
	for i := 0; i < 4; i++ {
		c.rooms[i] = make([]tileType, len(g.rooms[i]))
		copy(c.rooms[i], g.rooms[i])
	}
	return c
}

type gridHash uint32

var hasher = &fnv.Hasher{}

func (g grid) hash() gridHash {
	hasher.Reset()
	for _, pos := range haltablePositions {
		hasher.AddByte(g.hallway[pos].asByte())
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < len(g.rooms[i]); j++ {
			hasher.AddByte(g.rooms[i][j].asByte())
		}
	}
	return gridHash(hasher.Sum())
}

func (g grid) String() string {
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(g.energy))
	sb.WriteByte('\n')
	for i := 0; i < 11; i++ {
		sb.WriteByte(g.hallway[i].asByte())
	}
	for j := len(g.rooms[0]) - 1; j >= 0; j-- {
		sb.WriteByte('\n')
		sb.WriteByte(' ')
		for i := 0; i < 4; i++ {
			sb.WriteByte(' ')
			sb.WriteByte(g.rooms[i][j].asByte())
		}
	}
	sb.WriteByte('\n')
	return sb.String()
}

type tileType int8

func (t tileType) String() string {
	if t == tileOpen {
		return "."
	}
	return string(t.asByte())
}
func (t tileType) asByte() byte {
	if t == tileOpen {
		return '.'
	}
	return 'A' + byte(t)
}

const (
	tileA = iota
	tileB
	tileC
	tileD
	tileOpen
)

var costs = []int{1, 10, 100, 1000}
var haltablePositions = []tileType{0, 1, 3, 5, 7, 9, 10}

var benchmark = false
