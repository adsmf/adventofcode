package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	book := loadBlueprints()
	p1, p2 := solve(book)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type visitMap map[fnvHash]bool

func solve(book blueprintBook) (int, int) {
	visited := make(visitMap, 500000)
	p1, p2 := 0, 1
	for i := 0; i < 3; i++ {
		bp := book.blueprints[i]
		g1, g2 := runBlueprint(bp, 24, 32, &visited)
		p1 += g1 * (i + 1)
		p2 *= g2
	}
	for i := 3; i < book.numBlueprints; i++ {
		bp := book.blueprints[i]
		_, g1 := runBlueprint(bp, 24, 24, &visited)
		p1 += g1 * (i + 1)
	}
	return p1, p2
}

func runBlueprint(bp blueprint, recordStep, maxStep int, visited *visitMap) (int, int) {
	initialState := searchState{
		robots: materialSet{
			ore: 1,
		},
	}
	bestInterim := 0
	bestGeodes := 0
	for key := range *visited {
		delete(*visited, key)
	}
	const maxOpen = 160000
	openSet, nextOpen := [maxOpen]searchState{initialState}, [maxOpen]searchState{}
	openCount, nextCount := 1, 0
	search := func(state searchState) {
		if int(state.holding.geode) < bestGeodes-1 {
			return
		}
		const foldSize = 21
		hash := state.fnvHash()
		hash = (hash ^ hash>>foldSize) & (1<<foldSize - 1)
		if (*visited)[hash] {
			return
		}
		nextOpen[nextCount] = state
		nextCount++
		(*visited)[hash] = true
	}
	maxCost := materialSet{}
	maxCost = maxCost.max(bp.ore)
	maxCost = maxCost.max(bp.obsidian)
	maxCost = maxCost.max(bp.clay)
	maxCost = maxCost.max(bp.geode)
	maxCostSquared := maxCost.square()
	for step := 0; step < maxStep; step++ {
		if step == recordStep {
			bestInterim = bestGeodes
		}
		for i := 0; i < openCount; i++ {
			base := openSet[i]
			initialHolding := base.holding
			base.holding = base.holding.add(base.robots)
			base.holding = base.holding.capResources(maxCostSquared)
			base.robots = base.robots.capResources(maxCost)
			if int(base.holding.geode) > bestGeodes {
				bestGeodes = int(base.holding.geode)
			}
			search(base)
			if initialHolding.greater(bp.ore) && base.robots.ore < maxCost.ore {
				buy := base
				buy.holding = buy.holding.sub(bp.ore)
				buy.robots.ore++
				search(buy)
			}

			if initialHolding.greater(bp.obsidian) && base.robots.obsidian < maxCost.obsidian {
				buy := base
				buy.holding = buy.holding.sub(bp.obsidian)
				buy.robots.obsidian++
				search(buy)
			}

			if initialHolding.greater(bp.clay) && base.robots.clay < maxCost.clay {
				buy := base
				buy.holding = buy.holding.sub(bp.clay)
				buy.robots.clay++
				search(buy)
			}

			if initialHolding.greater(bp.geode) {
				buy := base
				buy.holding = buy.holding.sub(bp.geode)
				buy.robots.geode++
				search(buy)
			}
		}
		openSet, nextOpen = nextOpen, openSet
		openCount, nextCount = nextCount, 0
	}
	return bestInterim, bestGeodes
}

type searchState struct {
	holding materialSet
	robots  materialSet
}

func (s searchState) fnvHash() fnvHash {
	hash := fnvo32
	hash ^= fnvHash(s.holding.hash())
	hash *= fnvp32
	hash ^= fnvHash(s.robots.hash())
	hash *= fnvp32
	return hash
}

type fnvHash uint32

const (
	fnvp32 fnvHash = 0x01000193
	fnvo32 fnvHash = 0x811c9dc5
)

func loadBlueprints() blueprintBook {
	book := blueprintBook{}
	bpNum := 0
	for pos := 0; pos < len(input); pos++ {
		bp := blueprint{}
		_, pos = getInt(input, pos+10)
		bp.ore.ore, pos = getInt(input, pos+23)
		bp.clay.ore, pos = getInt(input, pos+28)
		bp.obsidian.ore, pos = getInt(input, pos+32)
		bp.obsidian.clay, pos = getInt(input, pos+9)
		bp.geode.ore, pos = getInt(input, pos+30)
		bp.geode.obsidian, pos = getInt(input, pos+9)
		pos += 10
		book.blueprints[bpNum] = bp
		bpNum++
		book.numBlueprints = bpNum
	}
	return book
}

func getInt(in []byte, pos int) (materialCount, int) {
	accumulator := materialCount(0)
	for ; in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += materialCount(in[pos] & 0xf)
	}
	return accumulator, pos
}

type blueprintBook struct {
	blueprints    [30]blueprint
	numBlueprints int
}
type blueprint struct {
	ore      materialSet
	clay     materialSet
	obsidian materialSet
	geode    materialSet
}

type materialSet struct {
	ore      materialCount
	obsidian materialCount
	clay     materialCount
	geode    materialCount
}

type materialCount uint8

func (m materialSet) hash() uint32 {
	return uint32(m.ore) |
		uint32(m.obsidian)<<8 |
		uint32(m.clay)<<16 |
		uint32(m.geode)<<24
}

func (m materialSet) add(oth materialSet) materialSet {
	return materialSet{
		ore:      m.ore + oth.ore,
		obsidian: m.obsidian + oth.obsidian,
		clay:     m.clay + oth.clay,
		geode:    m.geode + oth.geode,
	}
}
func (m materialSet) sub(oth materialSet) materialSet {
	return materialSet{
		ore:      m.ore - oth.ore,
		obsidian: m.obsidian - oth.obsidian,
		clay:     m.clay - oth.clay,
		geode:    m.geode - oth.geode,
	}
}
func (m materialSet) greater(cost materialSet) bool {
	return m.ore >= cost.ore &&
		m.obsidian >= cost.obsidian &&
		m.clay >= cost.clay &&
		m.geode >= cost.geode
}

func (m materialSet) max(oth materialSet) materialSet {
	return materialSet{
		ore:      max(m.ore, oth.ore),
		obsidian: max(m.obsidian, oth.obsidian),
		clay:     max(m.clay, oth.clay),
		geode:    max(m.geode, oth.geode),
	}
}

func (m materialSet) capResources(cap materialSet) materialSet {
	return materialSet{
		ore:      min(m.ore, cap.ore),
		obsidian: min(m.obsidian, cap.obsidian),
		clay:     min(m.clay, cap.clay),
		geode:    m.geode,
	}
}
func (m materialSet) square() materialSet {
	return materialSet{
		ore:      m.ore * m.ore,
		obsidian: m.obsidian * m.obsidian,
		clay:     m.clay * m.clay,
		geode:    m.geode * m.geode,
	}
}

func min(a, b materialCount) materialCount {
	if a < b {
		return a
	}
	return b
}
func max(a, b materialCount) materialCount {
	if a > b {
		return a
	}
	return b
}

var benchmark = false
