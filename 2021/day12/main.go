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
	routes, smallCaves := load(input)
	p1, p2 := explore(routes, smallCaves)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func explore(routes caveRouteList, smallCaves []bool) (int, int) {
	goodRoutesP1 := 0
	goodRoutesP2 := 0
	openSet := make([]routeState, 0, 100)
	nextOpen := make([]routeState, 0, 100)

	for _, next := range routes[caveStart] {
		openSet = append(openSet, routeState{pos: next, visited: caveStart})
	}

	for len(openSet) > 0 {
		nextOpen = nextOpen[0:0]
		for _, cur := range openSet {
			if cur.pos == caveEnd {
				if !cur.smallVisited {
					goodRoutesP1++
				}
				goodRoutesP2++
				continue
			}
			for _, nextCave := range routes[cur.pos] {
				smallVisited := cur.smallVisited
				if smallCaves[nextCave] && (1<<nextCave)&cur.visited > 0 {
					if cur.smallVisited {
						continue
					}
					smallVisited = true
				}
				nextState := routeState{
					pos:          nextCave,
					visited:      cur.visited | 1<<cur.pos,
					smallVisited: smallVisited,
				}
				nextOpen = append(nextOpen, nextState)
			}
		}
		openSet, nextOpen = nextOpen, openSet
	}

	return goodRoutesP1, goodRoutesP2
}

const (
	caveStart = 1
	caveEnd   = 2
)

func load(in string) (caveRouteList, []bool) {
	routes := make(caveRouteList, 15)
	smallCaves := make([]bool, 15)

	caveIDs := map[string]int{
		"start": caveStart,
		"end":   caveEnd,
	}

	nextID := caveEnd + 1

	for _, line := range utils.GetLines(in) {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		idA, idB := caveIDs[a], caveIDs[b]
		if idA == 0 {
			idA = nextID
			caveIDs[a] = idA
			nextID++
		}
		if idB == 0 {
			idB = nextID
			caveIDs[b] = idB
			nextID++
		}
		if a[0] >= 'a' {
			smallCaves[idA] = true
		}
		if b[0] >= 'a' {
			smallCaves[idB] = true
		}
		if b != "start" {
			routes[idA] = append(routes[idA], idB)
		}
		if a != "start" {
			routes[idB] = append(routes[idB], idA)
		}
	}

	return routes, smallCaves
}

type caveRouteList [][]int

type routeState struct {
	pos          int
	visited      int
	smallVisited bool
}

var benchmark = false
