package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	initial, rules := load(input)
	p1, p2 := expandPolymers(initial, rules)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func expandPolymers(polymerString string, rules map[string]byte) (int, int) {
	pairs := make(map[string]int)
	firstCh, lastCh := string(polymerString[0]), string(polymerString[len(polymerString)-1])
	last := polymerString[0]
	for i := 1; i < len(polymerString); i++ {
		pair := string(last) + string(polymerString[i])
		pairs[pair]++
		last = polymerString[i]
	}
	p1 := 0
	for i := 0; i < 40; i++ {
		if i == 10 {
			p1 = diffElements(pairs, firstCh, lastCh)
		}
		nextPairs := make(map[string]int, len(pairs))
		for pair, count := range pairs {
			replace := rules[pair]
			rep1 := string(pair[0]) + string(replace)
			rep2 := string(replace) + string(pair[1])
			nextPairs[rep1] += count
			nextPairs[rep2] += count
		}

		pairs = nextPairs
	}
	return p1, diffElements(pairs, firstCh, lastCh)
}

func diffElements(pairs map[string]int, first, last string) int {
	counts := make(map[string]int, 26)
	counts[first]++
	counts[last]++
	for pair, count := range pairs {
		counts[string(pair[0])] += count
		counts[string(pair[1])] += count
	}
	min, max := 9999999999999, 0
	for _, count := range counts {
		if min > count {
			min = count
		}
		if max < count {
			max = count
		}
	}
	return (max - min) / 2
}

func load(in string) (string, map[string]byte) {
	blocks := strings.Split(strings.TrimSpace(in), "\n\n")
	polymer := blocks[0]
	ruleLines := strings.Split(blocks[1], "\n")

	rules := make(map[string]byte, len(ruleLines))
	for _, line := range ruleLines {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1][0]
	}

	return polymer, rules
}

var benchmark = false
