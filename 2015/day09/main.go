package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	graph := load("input.txt")
	graph.enumerate()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", graph.min)
		fmt.Printf("Part 2: %d\n", graph.max)
	}
}

func load(filename string) graphData {
	graph := graphData{
		routes: map[pair]int{},
		cities: map[string]city{},
	}
	for _, line := range utils.ReadInputLines(filename) {
		parts := strings.Split(line, " ")
		a, b := parts[0], parts[2]
		graph.cities[a] = city{}
		graph.cities[b] = city{}
		dist, _ := strconv.Atoi(parts[4])
		graph.routes[pair{a, b}] = dist
		graph.routes[pair{b, a}] = dist
	}
	return graph
}

type graphData struct {
	cities   map[string]city
	routes   map[pair]int
	min, max int
}

func (g *graphData) enumerate() {
	cityList := []string{}
	for city := range g.cities {
		cityList = append(cityList, city)
	}
	possibleRoutes := utils.PermuteStrings(cityList)
	g.min, g.max = utils.MaxInt, 0
	for _, route := range possibleRoutes {
		routeDist := 0
		for i := 0; i < len(route)-1; i++ {
			routeDist += g.routes[pair{route[i], route[i+1]}]
		}
		if routeDist < g.min {
			g.min = routeDist
		}
		if routeDist > g.max {
			g.max = routeDist
		}
	}
}

type pair struct {
	a string
	b string
}

type city struct{}
