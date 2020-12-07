package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var benchmark = false

func main() {
	graph, reverseGraph := loadFile("input.txt")

	p1 := part1(reverseGraph)
	p2 := part2(graph)

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(reverseGraph containeeGraph) int {
	types := map[string]struct{}{}
	for containTypes, nextTypes := reverseGraph["shiny gold"], []string{}; len(containTypes) > 0; containTypes, nextTypes = nextTypes, []string{} {
		for _, cType := range containTypes {
			types[cType] = struct{}{}
			nextTypes = append(nextTypes, reverseGraph[cType]...)
		}
	}
	return len(types)
}

func part2(graph containmentGraph) int {
	contentCounts := map[string]int{}
	count := countContents(graph, contentCounts, "shiny gold")
	return count

}

func countContents(graph containmentGraph, knownCounts map[string]int, target string) int {
	if _, found := knownCounts[target]; !found {
		for innerType, innerCount := range graph[target] {
			knownCounts[target] += innerCount * (countContents(graph, knownCounts, innerType) + 1)
		}
	}
	return knownCounts[target]
}

func loadFile(filename string) (containmentGraph, containeeGraph) {
	forwardGraph := containmentGraph{}
	reverseGraph := containeeGraph{}

	inputBytes, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(inputBytes), "\n") {
		if line == "" || strings.HasSuffix(line, "no other bags.") {
			continue
		}
		specParts := strings.SplitN(line, "contain", 2)
		container, _ := parseDef(specParts[0])
		if forwardGraph[container] == nil {
			forwardGraph[container] = map[string]int{}
		}
		for _, contains := range strings.Split(specParts[1], ", ") {
			cType, cCount := parseDef(contains)
			forwardGraph[container][cType] = cCount
			reverseGraph[cType] = append(reverseGraph[cType], container)
		}
	}

	return forwardGraph, reverseGraph
}

func parseDef(def string) (desc string, count int) {
	parts := strings.Split(strings.TrimSpace(def), " ")
	if len(parts) == 4 {
		count, _ = strconv.Atoi(parts[0])
		parts = parts[1:]
	}
	desc = strings.Join(parts[0:2], " ")
	return desc, count
}

type containmentGraph map[string]map[string]int
type containeeGraph map[string][]string
