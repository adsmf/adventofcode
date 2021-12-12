package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	routes := load(input)
	p1, p2 := explore(routes)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func explore(routes caveRouteList) (int, int) {
	goodRoutesP1 := 0
	goodRoutesP2 := 0
	openSet := []routeState{}

	for _, next := range routes["start"] {
		openSet = append(openSet, routeState{pos: next, visited: "start"})
	}

	for len(openSet) > 0 {
		nextOpen := []routeState{}
		for _, cur := range openSet {
			if cur.pos == "end" {
				if !cur.smallVisited {
					goodRoutesP1++
				}
				goodRoutesP2++
				continue
			}
			for _, nextCave := range routes[cur.pos] {
				if nextCave == "start" {
					continue
				}
				smallVisited := cur.smallVisited
				if nextCave[0] >= 'a' && strings.Contains(cur.visited, nextCave) {
					if cur.smallVisited {
						continue
					}
					smallVisited = true
				}
				nextState := routeState{
					pos:          nextCave,
					visited:      cur.visited + "," + cur.pos,
					smallVisited: smallVisited,
				}
				nextOpen = append(nextOpen, nextState)
			}
		}
		openSet = nextOpen
	}

	return goodRoutesP1, goodRoutesP2
}

func load(in string) caveRouteList {
	routes := caveRouteList{}

	for _, line := range utils.GetLines(in) {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		if routes[a] == nil {
			routes[a] = []string{}
		}
		if routes[b] == nil {
			routes[b] = []string{}
		}
		routes[a] = append(routes[a], b)
		routes[b] = append(routes[b], a)
	}

	return routes
}

type caveRouteList map[string][]string

type routeState struct {
	pos          string
	visited      string
	smallVisited bool
}

var benchmark = false
