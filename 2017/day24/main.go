package main

import (
	"fmt"
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
	adapters := load("input.txt")
	_, strength := bestChain(0, adapters)
	return strength
}

func part2() int {
	adapters := load("input.txt")
	longestChains := longestChain(0, adapters)
	best := 0
	for _, chain := range longestChains {
		strength := 0
		for _, adapter := range chain {
			strength += adapter.strength()
		}
		if strength > best {
			best = strength
		}
	}
	return best
}

func longestChain(left int, available []adapterInfo) [][]adapterInfo {
	if len(available) == 0 {
		return [][]adapterInfo{}
	}
	bestLength := 0
	longestChains := [][]adapterInfo{}
	for idx, adapter := range available {
		valid, leaves := adapter.joinTo(left)
		if !valid {
			continue
		}
		nextAvail := make([]adapterInfo, 0, len(available)-1)
		for nextIdx, next := range available {
			if idx != nextIdx {
				nextAvail = append(nextAvail, next)
			}
		}
		bestSubChains := longestChain(leaves, nextAvail)
		if len(bestSubChains) > 0 {
			length := len(bestSubChains[0]) + 1
			if length < bestLength {
				continue
			}
			chains := [][]adapterInfo{}
			for _, subChain := range bestSubChains {
				chain := append([]adapterInfo{adapter}, subChain...)
				chains = append(chains, chain)
			}
			if length > bestLength {
				bestLength = length
				longestChains = chains
				continue
			}
			for _, chain := range chains {
				longestChains = append(longestChains, chain)
			}
		} else {
			length := 1
			chain := []adapterInfo{adapter}
			if length > bestLength {
				bestLength = length
				longestChains = [][]adapterInfo{chain}
			} else if length == bestLength {
				longestChains = append(longestChains, chain)
			}
		}
	}
	return longestChains
}

func bestChain(left int, available []adapterInfo) ([]adapterInfo, int) {
	if len(available) == 0 {
		return []adapterInfo{}, 0
	}
	bestScore := 0
	best := []adapterInfo{}
	for idx, adapter := range available {
		valid, leaves := adapter.joinTo(left)
		if !valid {
			continue
		}
		score := adapter.strength()
		nextAvail := make([]adapterInfo, 0, len(available)-1)
		for nextIdx, next := range available {
			if idx != nextIdx {
				nextAvail = append(nextAvail, next)
			}
		}
		subChain, subScore := bestChain(leaves, nextAvail)
		score += subScore
		if score > bestScore {
			bestScore = score
			best = append([]adapterInfo{adapter}, subChain...)
		}
	}
	return best, bestScore
}

type adapterInfo struct{ p1, p2 int }

func (a adapterInfo) strength() int { return a.p1 + a.p2 }

func (a adapterInfo) joinTo(prev int) (bool, int) {
	if prev == a.p1 {
		return true, a.p2
	} else if prev == a.p2 {
		return true, a.p1
	}
	return false, -1
}

func load(filename string) []adapterInfo {
	adapters := []adapterInfo{}

	for _, line := range utils.ReadInputLines(filename) {
		parts := strings.Split(line, "/")
		a := adapterInfo{
			p1: utils.MustInt(parts[0]),
			p2: utils.MustInt(parts[1]),
		}
		adapters = append(adapters, a)
	}

	return adapters
}

var benchmark = false
