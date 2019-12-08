package astar

import (
	"fmt"
	"math"
)

// Implementation of A* search algorithm
// Based on pseudocode on https://en.wikipedia.org/wiki/A*_search_algorithm

// Node represents any routable node
type Node interface {
	Paths() []Edge
	Heuristic(from Node) Cost
}

// Edge stores details for an edge in the graph
type Edge struct {
	To   Node
	Cost Cost
}

// Cost is, unsurprisingly, the cost associated with a graph edge
type Cost float64

type routedNode struct {
}

// Route finds a path from start to goal, returning the list of nodes to visit
func Route(start, goal Node) ([]Node, error) {
	openSet := nodeSet{}
	openSet.Add(start)

	cameFrom := parentSet{}

	gScore := costSet{}
	gScore.Set(start, 0)

	fScore := costSet{}
	fScore.Set(start, goal.Heuristic(start))

	for len(openSet) > 0 {
		current := openSet.getLowest(fScore)
		if current == goal {
			return reconstructPath(cameFrom, current), nil
		}

		openSet.Remove(current)

		for _, edge := range current.Paths() {
			neighbour := edge.To
			gScoreTentative := gScore.Cost(current) + edge.Cost

			if gScoreTentative <= gScore.Cost(neighbour) {
				cameFrom.Set(neighbour, current)
				gScore.Set(neighbour, gScoreTentative)
				fScore.Set(neighbour, gScoreTentative+goal.Heuristic(neighbour))
				openSet.Add(neighbour)
			}
		}
	}

	return nil, fmt.Errorf("Unable to find route from %#v to %#v", start, goal)
}

func reconstructPath(cameFrom parentSet, current Node) []Node {
	path := []Node{current}

	for cameFrom.Contains(current) {
		current = cameFrom.Get(current)
		path = append([]Node{current}, path...)
	}
	return path
}

type nodeMap map[Node]interface{}

func (nm nodeMap) Remove(node Node) {
	delete(nm, node)
}

func (nm nodeMap) Contains(node Node) bool {
	if _, found := nm[node]; found {
		return true
	}
	return false
}

func (nm nodeMap) getLowest(scores costSet) Node {
	bestScore := Cost(math.MaxFloat64)
	var bestNode Node
	for tryNode := range nm {
		score := scores.Cost(tryNode)
		if score <= bestScore {
			bestScore = score
			bestNode = tryNode
		}
	}
	return bestNode
}

type parentSet nodeMap

func (pm parentSet) Set(node, parent Node)   { pm[node] = parent }
func (pm parentSet) Get(node Node) Node      { return pm[node].(Node) }
func (pm parentSet) Contains(node Node) bool { return nodeMap(pm).Contains(node) }

type nodeSet nodeMap

func (ns nodeSet) Add(node Node)                 { ns[node] = true }
func (ns nodeSet) Remove(node Node)              { nodeMap(ns).Remove(node) }
func (ns nodeSet) getLowest(scores costSet) Node { return nodeMap(ns).getLowest(scores) }

type costSet nodeMap

func (cs costSet) Set(node Node, cost Cost) { cs[node] = cost }
func (cs costSet) Cost(node Node) Cost {
	if cost, found := cs[node]; found {
		return cost.(Cost)
	}
	return math.MaxFloat32
}
