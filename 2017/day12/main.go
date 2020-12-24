package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	graphs := loadTrees("input.txt")
	p1 := part1(graphs)
	p2 := len(graphs)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(graphs []progSet) int {
	for _, graph := range graphs {
		if graph[0] {
			return len(graph)
		}
	}
	return -1
}

func loadTrees(filename string) []progSet {
	lines := utils.ReadInputLines(filename)
	graphs := []progSet{}

	for _, line := range lines {
		progs := utils.GetInts(line)
		matches := map[int]bool{}
		for _, prog := range progs {
			for gID, graph := range graphs {
				if graph[prog] {
					matches[gID] = true
				}
			}
		}
		switch len(matches) {
		case 0:
			newGraph := progSet{}
			for _, prog := range progs {
				newGraph[prog] = true
			}
			graphs = append(graphs, newGraph)
		case 1:
			first := 0
			for iter := range matches {
				first = iter
			}
			for _, prog := range progs {
				graphs[first][prog] = true
			}
		default:
			newGraph := progSet{}
			for _, prog := range progs {
				newGraph[prog] = true
			}
			for gID := range matches {
				for prog := range graphs[gID] {
					newGraph[prog] = true
				}
			}
			newGraphs := []progSet{newGraph}
			for gID := range graphs {
				if !matches[gID] {
					newGraphs = append(newGraphs, graphs[gID])
				}
			}
			graphs = newGraphs
		}
	}

	return graphs
}

type progSet map[int]bool

var benchmark = false
