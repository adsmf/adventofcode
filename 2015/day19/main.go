package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils"
	"github.com/adsmf/adventofcode/utils/pathfinding/astar"
	"strings"
)

var globalReplacements replacementOptions

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	l := loadInput("input.txt")
	return len(getOpts(l.initial, l.replacements))
}

func part2() int {
	l := loadInput("input.txt")
	return l.findTarget("e", l.initial)
}

type replacementOptions map[molecule][]molecule
type lab struct {
	replacements replacementOptions
	initial      molecule
}

func (l lab) findTarget(initial, target molecule) int {
	route, err := astar.Route(target, initial)
	if err != nil {
		panic(err)
	}
	return len(route) - 1
}

func getOpts(initial molecule, reps replacementOptions) moleculeList {
	opts := moleculeList{}
	for i := 0; i < len(initial); i++ {
		pre := string(initial[:i])
		sub := string(initial[i:])
		for rep, withOpts := range reps {
			if strings.HasPrefix(sub, string(rep)) {
				suf := strings.TrimPrefix(sub, string(rep))
				for _, with := range withOpts {
					opts[molecule(pre+string(with)+suf)] = struct{}{}
				}
			}
		}
	}
	return opts
}

type molecule string

func (m molecule) Heuristic(from astar.Node) astar.Cost {
	fromMolecule := from.(molecule)
	dist := levenshtein(string(m), string(fromMolecule))
	return astar.Cost(dist)
}

func (m molecule) Paths() []astar.Edge {
	paths := []astar.Edge{}

	for opt := range getOpts(m, globalReplacements) {
		paths = append(
			paths,
			astar.Edge{
				To:   opt,
				Cost: 1,
			},
		)
	}
	return paths
}

// https://www.rosettacode.org/wiki/Levenshtein_distance#Go
func levenshtein(s, t string) int {
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return d[len(s)][len(t)]
}

type moleculeList map[molecule]struct{}

func (m *moleculeList) add(item molecule) {
	(*m)[item] = struct{}{}
}

func loadInput(filename string) lab {
	l := lab{
		replacements: map[molecule][]molecule{},
	}
	globalReplacements = replacementOptions{}
	for _, line := range utils.ReadInputLines(filename) {
		if line == "" {
			continue
		}
		if strings.Contains(line, " => ") {
			parts := strings.Split(line, " => ")
			rep := molecule(parts[0])
			with := molecule(parts[1])
			// For part 1
			if l.replacements[rep] == nil {
				l.replacements[rep] = []molecule{}
			}
			l.replacements[rep] = append(l.replacements[rep], with)

			// For part 2
			if globalReplacements[rep] == nil {
				globalReplacements[rep] = []molecule{}
			}
			globalReplacements[rep] = append(globalReplacements[rep], with)

			if globalReplacements[with] == nil {
				globalReplacements[with] = []molecule{}
			}
			globalReplacements[with] = append(globalReplacements[with], rep)
			continue
		}
		l.initial = molecule(line)
	}
	return l
}
