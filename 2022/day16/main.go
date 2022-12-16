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
	valves, startValve := loadValves()
	costs := calcRoutes(valves, startValve)

	p1 := part1(valves, startValve, costs)
	p2 := part2(valves, startValve, costs)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(valves valveSet, startValve int, costs routeCosts) int {
	initialState := searchEntry{
		valve: startValve,
	}
	return findBest(initialState, valves, startValve, costs, false)
}

func part2(valves valveSet, startValve int, costs routeCosts) int {
	initialState := searchEntry{
		valve:    startValve,
		othValve: startValve,
		time:     4,
		othTime:  4,
	}
	return findBest(initialState, valves, startValve, costs, true)
}

func findBest(initialState searchEntry, valves valveSet, startValve int, routes routeCosts, withElephant bool) int {
	openSet := searchSet{}
	openSet.push(initialState)
	maxFlowRate := 0
	for _, valve := range valves {
		maxFlowRate += valve.rate
	}
	best := 0
	cap := 50000
	if withElephant {
		cap *= 30
	}
	visited := make(map[dfsEntry]int, cap)
	nextStates := searchSet{}
	for openSet.numEntries > 0 {
		curState := openSet.pop()
		nextStates.reset()
		curState.nextStates(valves, routes, withElephant, &nextStates)
		for nextStates.numEntries > 0 {
			nextState := nextStates.pop()
			dfsState := toDFSEntry(nextState)
			if curState.cumulativeFlow > 0 && visited[dfsState] >= curState.cumulativeFlow {
				continue
			}
			visited[dfsState] = curState.cumulativeFlow
			if nextState.cumulativeFlow > best {
				best = nextState.cumulativeFlow
			}
			maxAchieveable := nextState.cumulativeFlow + maxFlowRate*(30-nextState.time)
			if maxAchieveable < best {
				continue
			}
			openSet.push(nextState)
		}
	}
	return best
}

type dfsEntry uint64

func toDFSEntry(s searchEntry) dfsEntry {
	return dfsEntry(s.valvesOpen) | dfsEntry(s.valve)<<40 | dfsEntry(s.othValve)<<48 | dfsEntry(s.time)<<56
}

type searchSet struct {
	entries    [100]searchEntry
	numEntries int
}

func (s *searchSet) reset() { s.numEntries = 0 }
func (s *searchSet) push(entry searchEntry) {
	s.entries[s.numEntries] = entry
	s.numEntries++
}
func (s *searchSet) pop() searchEntry {
	s.numEntries--
	return s.entries[s.numEntries]
}

type searchEntry struct {
	valvesOpen      uint64
	valve, othValve int
	time, othTime   int
	cumulativeFlow  int
}

func (s searchEntry) nextStates(valves valveSet, routes routeCosts, withElephant bool, nextStates *searchSet) {
	for i := 0; i < routes[s.valve].count; i++ {
		route := routes[s.valve].routes[i]
		if route.cost == 0 {
			continue
		}
		if s.time+route.cost > 30 {
			continue
		}
		if s.valvesOpen&(1<<route.target) > 0 {
			continue
		}
		nextState := s
		nextState.valve = route.target
		stepTime := nextState.time + route.cost + 1
		nextState.valvesOpen |= 1 << route.target
		nextState.time = stepTime
		nextState.cumulativeFlow += (30 - stepTime) * valves[route.target].rate
		if withElephant && nextState.time > nextState.othTime {
			nextState.time, nextState.othTime = nextState.othTime, nextState.time
			nextState.valve, nextState.othValve = nextState.othValve, nextState.valve
		}
		nextStates.push(nextState)
	}
}

const maxValve = 60
const maxRoute = 20

type routeCosts [maxValve]routesInfo
type routesInfo struct {
	count  int
	routes [maxRoute]routeInfo
}

func (r *routesInfo) add(route routeInfo) {
	r.routes[r.count] = route
	r.count++
}

type routeInfo struct {
	target int
	cost   int
}

func calcRoutes(valves valveSet, startValve int) routeCosts {
	costs := routeCosts{}
	for fromId := range valves {
		openSet, nextOpen := []int{0}, []int{}
		if valves[fromId].rate == 0 && fromId != startValve {
			continue
		}
		openSet = openSet[0:0]
		nextOpen = nextOpen[0:0]
		visited := [maxValve]bool{}
		visited[fromId] = true
		openSet = append(openSet, fromId)
		for steps := 0; len(openSet) > 0; steps++ {
			for _, curValve := range openSet {
				if valves[curValve].rate > 0 {
					costs[fromId].add(routeInfo{target: curValve, cost: steps})
				}
				for i := 0; i < int(valves[curValve].numNext); i++ {
					next := valves[curValve].nextValves[i]
					if !visited[next] {
						nextOpen = append(nextOpen, next)
						visited[next] = true
					}
				}
			}
			openSet, nextOpen = nextOpen, openSet
			nextOpen = nextOpen[0:0]
		}
	}
	return costs
}

type valveSet [maxValve]valveInfo
type valveInfo struct {
	rate       int
	nextValves [5]int
	numNext    byte
}

func loadValves() (valveSet, int) {
	valves := valveSet{}

	valveIDs := map[string]int{}
	nextValveList := [][]string{}
	startValve := 0
	for idx, line := range utils.GetLines(input) {
		parts := strings.SplitN(line, " ", 10)
		valveName := parts[1]
		if valveName == "AA" {
			startValve = idx
		}
		var flowRate int
		fmt.Sscanf(parts[4], "rate=%d;", &flowRate)
		toValveList := parts[9]
		toValves := strings.Split(toValveList, ", ")
		valves[idx] = valveInfo{
			rate: flowRate,
		}
		valveIDs[valveName] = idx
		nextValveList = append(nextValveList, toValves)
	}
	for from, toList := range nextValveList {
		valves[from].numNext = byte(len(toList))
		for idx, toName := range toList {
			valves[from].nextValves[idx] = valveIDs[toName]
		}
	}
	return valves, startValve
}

var benchmark = false
