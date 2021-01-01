package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"strings"

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
	scan := load("input.txt")
	risk := 0
	for y := 0; y <= scan.target.y; y++ {
		for x := 0; x <= scan.target.x; x++ {
			pos := point{x, y}
			risk += int(scan.getRegionType(pos))
		}
	}
	return risk
}

func part2() int {
	scan := load("input.txt")
	searcher := searcherState{point{0, 0}, equipTorch, 0}
	openStates := &searchHeap{searcher}
	heap.Init(openStates)
	visited := map[searcherState]bool{}
	for openStates.Len() > 0 {
		state := heap.Pop(openStates).(searcherState)
		if state.searchTime > 10000 {
			return -1
		}
		if state.position == scan.target && state.equipment == equipTorch {
			return state.searchTime
		}
		checkState := state
		checkState.searchTime = 0
		if visited[checkState] {
			continue
		}
		visited[checkState] = true

		for _, dir := range [4]point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			nextPos := point{state.position.x + dir.x, state.position.y + dir.y}
			if nextPos.x < 0 || nextPos.y < 0 {
				continue
			}
			nextType := scan.getRegionType(nextPos)
			if state.equipment.validIn(nextType) {
				nextState := searcherState{
					position:   nextPos,
					equipment:  state.equipment,
					searchTime: state.searchTime + 1,
				}
				heap.Push(openStates, nextState)
			}
		}

		curType := scan.getRegionType(state.position)
		for nextEquip := equippedItem(0); nextEquip < equipWRAP; nextEquip++ {
			if nextEquip == state.equipment {
				continue
			}
			if nextEquip.validIn(curType) {
				nextState := searcherState{
					position:   state.position,
					equipment:  nextEquip,
					searchTime: state.searchTime + 7,
				}
				heap.Push(openStates, nextState)
			}
		}
	}
	return -1
}

type searcherState struct {
	position   point
	equipment  equippedItem
	searchTime int
}

type searchHeap []searcherState

func (h searchHeap) Len() int            { return len(h) }
func (h searchHeap) Less(i, j int) bool  { return h[i].searchTime < h[j].searchTime }
func (h searchHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *searchHeap) Push(x interface{}) { *h = append(*h, x.(searcherState)) }
func (h *searchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type equippedItem int

const (
	equipTorch equippedItem = iota
	equipCliming
	equipNeither
	equipWRAP
)

var validEquipment = map[terrainType]map[equippedItem]bool{
	elRocky:  {equipCliming: true, equipTorch: true},
	elWet:    {equipCliming: true, equipNeither: true},
	elNarrow: {equipTorch: true, equipNeither: true},
}

func (ei equippedItem) validIn(tt terrainType) bool { return validEquipment[tt%3][ei] }

type erosionLevel int

type terrainType int

const (
	elRocky terrainType = iota
	elWet
	elNarrow
)

var terrainRepr = map[terrainType]byte{elRocky: '.', elWet: '=', elNarrow: '|'}

func (t terrainType) String() string {
	return string(terrainRepr[t%3])
}
func (t terrainType) asByte() byte {
	return terrainRepr[t%3]
}

type caveScan struct {
	geoIndex map[point]int
	terrain  map[point]terrainType
	depth    int
	target   point
}

func (c caveScan) String() string {
	sb := &strings.Builder{}
	sb.Grow((c.target.x + 2) * (c.target.y + 1))
	for y := 0; y <= c.target.y; y++ {
		for x := 0; x <= c.target.x; x++ {
			if x == 0 && y == 0 {
				sb.WriteByte('M')
				continue
			}
			pos := point{x, y}
			if pos == c.target {
				sb.WriteByte('T')
				continue
			}
			sb.WriteByte(c.getRegionType(pos).asByte())
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (c *caveScan) getGeoIndex(pos point) int {
	if gi, found := c.geoIndex[pos]; found {
		return gi
	}
	var gi int
	if pos.x == 0 && pos.y == 0 {
		gi = 0
	} else if pos.y == 0 {
		gi = pos.x * 16807
	} else if pos.x == 0 {
		gi = pos.y * 48271
	} else {
		gi = int(c.getErosionLevel(point{pos.x - 1, pos.y})) * int(c.getErosionLevel(point{pos.x, pos.y - 1}))
	}
	c.geoIndex[pos] = gi
	return gi
}

func (c *caveScan) getErosionLevel(pos point) erosionLevel {
	return erosionLevel((c.getGeoIndex(pos) + c.depth) % 20183)
}

func (c *caveScan) getRegionType(pos point) terrainType {
	if terrain, found := c.terrain[pos]; found {
		return terrain
	}
	if pos == c.target {
		return elRocky
	}
	terrain := terrainType(c.getErosionLevel(pos) % 3)
	c.terrain[pos] = terrain
	return terrain
}

type point struct{ x, y int }

func load(filename string) caveScan {
	inputBytes, _ := ioutil.ReadFile(filename)
	ints := utils.GetInts(string(inputBytes))
	depth, target := ints[0], point{ints[1], ints[2]}
	return caveScan{
		geoIndex: map[point]int{{0, 0}: 0, target: 0},
		terrain:  map[point]terrainType{},
		depth:    depth,
		target:   target,
	}
}

var benchmark = false
