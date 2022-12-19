package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	blueprints := loadBlueprints()
	p1, p2 := solve(blueprints)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(blueprints blueprintBook) (int, int) {
	p1, p2 := 0, 1
	for i, bp := range blueprints[:3] {
		g1, g2 := runBlueprint(bp, 24, 32)
		p1 += g1 * (i + 1)
		p2 *= g2
	}
	for i, bp := range blueprints[3:] {
		_, g1 := runBlueprint(bp, 24, 24)
		p1 += g1 * (i + 4)
	}
	return p1, p2
}

func runBlueprint(bp blueprint, recordStep, maxStep int) (int, int) {
	initialState := searchState{
		robots: materialSet{
			ore: 1,
		},
	}
	bestInterim := 0
	bestGeodes := 0
	visited := map[fnvHash]bool{}
	openSet, nextOpen := []searchState{initialState}, []searchState{}
	search := func(state searchState) {
		hash := state.fnvHash()
		if visited[hash] {
			return
		}
		if int(state.holding.geode) < bestGeodes-1 {
			return
		}
		nextOpen = append(nextOpen, state)
		visited[hash] = true
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
		nextOpen = nextOpen[0:0]
		for _, base := range openSet {
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
	blueprints := blueprintBook{}
	for _, line := range utils.GetLines(input) {
		vals := utils.GetInts(line)
		blueprints = append(blueprints, blueprint{
			ore:      materialSet{ore: materialCount(vals[1])},
			clay:     materialSet{ore: materialCount(vals[2])},
			obsidian: materialSet{ore: materialCount(vals[3]), clay: materialCount(vals[4])},
			geode:    materialSet{ore: materialCount(vals[5]), obsidian: materialCount(vals[6])},
		})
	}
	return blueprints
}

type blueprintBook []blueprint
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
