package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	discs := load("input.txt")
	p1 := part1(discs)
	p2 := part2(discs)
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(discs map[string]discInfo) string {
	containedByCount := map[string]int{}
	for _, disc := range discs {
		for cont := range disc.contains {
			containedByCount[cont]++
		}
	}
	for id := range discs {
		if containedByCount[id] == 0 {
			return id
		}
	}
	return ""
}

func part2(discs map[string]discInfo) int {
	weights := map[string]int{}
	root := part1(discs)
	calcWeights(root, discs, weights)

	offset := 0
	for {
		balanced, bad, newOffset := isBalanced(root, discs, weights)
		if balanced {
			target := discs[root]
			return target.weight - offset
		}
		offset = newOffset
		root = bad
	}
}

func isBalanced(id string, discs map[string]discInfo, weights map[string]int) (bool, string, int) {
	contents := discs[id].contains
	if len(contents) == 0 {
		return true, "", 0
	}

	contWeights := map[int]int{}
	w := 0
	total := 0
	for cont := range contents {
		w = weights[cont]
		total += w
		contWeights[w]++
	}
	if len(contWeights) == 1 {
		return true, "", w
	}
	bad := 0
	good := 0
	badNode := ""
	for cont := range contents {
		count := contWeights[weights[cont]]
		if count == 1 {
			bad = weights[cont]
			badNode = cont
		} else {
			good = weights[cont]
		}
	}
	return false, badNode, bad - good
}

func calcWeights(id string, discs map[string]discInfo, weights map[string]int) int {
	if weight, found := weights[id]; found {
		return weight
	}
	weight := discs[id].weight

	for cont := range discs[id].contains {
		weight += calcWeights(cont, discs, weights)
	}
	weights[id] = weight
	return weight
}

type discInfo struct {
	weight   int
	contains map[string]bool
}

func load(filename string) map[string]discInfo {
	discs := map[string]discInfo{}

	for _, line := range utils.ReadInputLines(filename) {
		parts := strings.Split(line, " -> ")
		prog, weight := "", 0
		fmt.Sscanf(parts[0], "%s (%d)", &prog, &weight)
		d := discInfo{
			weight:   weight,
			contains: map[string]bool{},
		}
		if len(parts) > 1 {
			for _, cont := range strings.Split(parts[1], ", ") {
				d.contains[cont] = true
			}
		}
		discs[prog] = d
	}
	return discs
}

var benchmark = false
