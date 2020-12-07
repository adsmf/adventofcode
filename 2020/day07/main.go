package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var benchmark = false

func main() {
	p1 := part1("input.txt")
	p2 := part2("input.txt")
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(filename string) int {
	graph := loadFile(filename)
	reverseGraph := map[string][]string{}
	for outerType, innerMap := range graph {
		for innerType := range innerMap {
			reverseGraph[innerType] = append(reverseGraph[innerType], outerType)
		}
	}
	types := map[string]struct{}{}
	containTypes := reverseGraph["shiny gold"]
	for len(containTypes) > 0 {
		nextTypes := []string{}
		for _, cType := range containTypes {
			if _, found := types[cType]; !found {
				types[cType] = struct{}{}
				nextTypes = append(nextTypes, reverseGraph[cType]...)
			}
		}
		containTypes = nextTypes
	}
	return len(types)
}

func part2(filename string) int {
	graph := loadFile(filename)
	contentCounts := map[string]int{}
	count := countContents(graph, contentCounts, "shiny gold")
	return count

}

func countContents(graph containmentGraph, knownCounts map[string]int, target string) int {
	if count, found := knownCounts[target]; found {
		return count
	}
	count := 0
	for innerType, innerCount := range graph[target] {
		count += innerCount * (countContents(graph, knownCounts, innerType) + 1)
	}
	knownCounts[target] = count
	return count
}

func loadFile(filename string) containmentGraph {
	inputBytes, _ := ioutil.ReadFile(filename)
	return load(string(inputBytes))
}
func load(input string) containmentGraph {
	graph := containmentGraph{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		sides := strings.SplitN(line, "contain", 2)
		leftType, _ := parseDef(sides[0])
		right := strings.TrimSpace(sides[1])
		if right == "no other bags." {
			continue
		}
		if graph[leftType] == nil {
			graph[leftType] = map[string]int{}
		}
		for _, contains := range strings.Split(right, ", ") {
			cType, cCount := parseDef(contains)
			graph[leftType][cType] = cCount
		}
	}
	return graph
}

func parseDef(def string) (string, int) {
	count := 0
	parts := strings.Split(strings.TrimSpace(def), " ")
	if len(parts) == 4 {
		count, _ = strconv.Atoi(parts[0])
		parts = parts[1:]
	}
	desc := strings.Join(parts[0:2], " ")
	return desc, count
}

type containmentGraph map[string]map[string]int
