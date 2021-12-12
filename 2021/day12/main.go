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
	openSet := make([]routeState, 0, 100)

	for _, next := range routes["start"] {
		openSet = append(openSet, routeState{pos: next, visited: "start"})
	}

	nextOpen := make([]routeState, 0, 100)
	for len(openSet) > 0 {
		nextOpen = nextOpen[0:0]
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
				route := strings.Builder{}
				route.Grow(len(cur.visited) + len(cur.pos) + 1)
				route.WriteString(cur.visited)
				route.WriteByte(',')
				route.WriteString(cur.pos)
				nextState := routeState{
					pos:          nextCave,
					visited:      route.String(),
					smallVisited: smallVisited,
				}
				nextOpen = append(nextOpen, nextState)
			}
		}
		openSet, nextOpen = nextOpen, openSet
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
