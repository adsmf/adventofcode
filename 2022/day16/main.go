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

	openSet := []searchEntry{initialState}
	maxFlowRate := 0
	for _, valve := range valves {
		maxFlowRate += valve.rate
	}
	best := 0
	visited := map[dfsEntry]int{}
	for len(openSet) > 0 {
		var curState searchEntry
		curState, openSet = openSet[len(openSet)-1], openSet[0:len(openSet)-1]
		for _, nextState := range curState.nextStates(valves, routes, withElephant) {
			dfsState := dfsEntry{
				time:       nextState.time,
				othTime:    nextState.othTime,
				valve:      nextState.valve,
				othValve:   nextState.valve,
				valvesOpen: nextState.valvesOpen,
			}
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
			openSet = append(openSet, nextState)
		}
	}
	return best
}

type dfsEntry struct {
	valvesOpen      uint64
	valve, othValve int
	time, othTime   int
}

type searchEntry struct {
	valvesOpen      uint64
	valve, othValve int
	time, othTime   int
	cumulativeFlow  int
}

func (s searchEntry) nextStates(valves valveSet, routes routeCosts, withElephant bool) []searchEntry {
	nextStates := []searchEntry{}

	for nextValve, cost := range routes[s.valve] {
		if s.time+cost > 30 {
			continue
		}
		if s.valvesOpen&(1<<nextValve) > 0 {
			continue
		}
		nextState := s
		nextState.valve = nextValve
		stepTime := nextState.time + cost + 1
		nextState.valvesOpen |= 1 << nextValve
		nextState.time = stepTime
		nextState.cumulativeFlow += (30 - stepTime) * valves[nextValve].rate
		if withElephant && nextState.time > nextState.othTime {
			nextState.time, nextState.othTime = nextState.othTime, nextState.time
			nextState.valve, nextState.othValve = nextState.othValve, nextState.valve
		}
		nextStates = append(nextStates, nextState)
	}

	return nextStates
}

type routeCosts map[int]map[int]int

func calcRoutes(valves valveSet, startValve int) routeCosts {
	costs := routeCosts{}
	for fromId := range valves {
		openSet, nextOpen := []int{0}, []int{}
		if valves[fromId].rate == 0 && fromId != startValve {
			continue
		}
		openSet = openSet[0:0]
		nextOpen = nextOpen[0:0]
		visited := [60]bool{}
		visited[fromId] = true
		openSet = append(openSet, fromId)
		costs[fromId] = map[int]int{}
		for steps := 0; len(openSet) > 0; steps++ {
			for _, curValve := range openSet {
				if valves[curValve].rate > 0 {
					costs[fromId][curValve] = steps
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

type valveSet [60]valveInfo
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
