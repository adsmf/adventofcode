package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

func setTargets(cavern *grid, targets []*creature) {
	for y, row := range *cavern {
		for x := range row {
			(*cavern)[y][x].isTarget = false
		}
	}
	for _, target := range targets {
		if target.location.occupiedBy != target {
			panic("Back reference broken")
		}
		target.location.isTarget = true
	}
	numTargets := 0
	for y, row := range *cavern {
		for x := range row {
			if (*cavern)[y][x].isTarget {
				numTargets++
			}
		}
	}
	if numTargets != len(targets) {
		panic("Failed to set targets")
	}
}

func moveTowardBestTarget(cavern *grid, start *gridSquare) {
	verticies := []*vertex{}
	targets := []*vertex{}

	for y, row := range *cavern {
		for x := range row {
			(*cavern)[y][x].cost = utils.MaxInt
			if !!!(*cavern)[y][x].isCavern {
				continue
			}
			isStart := (*cavern)[y][x] == start
			isTarget := (*cavern)[y][x].isTarget == true
			if (*cavern)[y][x].occupiedBy != nil && !!!(isStart || isTarget) {
				continue
			}
			newVertex := &vertex{
				point: point{
					x: x,
					y: y,
				},
				gridSquare: (*cavern)[y][x],
			}
			verticies = append(verticies, newVertex)
			if newVertex.gridSquare.isTarget {
				targets = append(targets, newVertex)
			}

		}
	}
	start.cost = 0

	if len(targets) == 0 {
		panic("Don't have any targets!")
	}

	for len(verticies) > 0 {
		sort.Slice(verticies, func(i, j int) bool {
			return compareVertices(verticies[i], verticies[j])
		})
		minNode := verticies[0]
		verticies = verticies[1:]
		for _, v := range verticies {
			if minNode.manhattenDistance(v) == 1 {
				newCost := minNode.gridSquare.cost + 1
				if newCost < v.gridSquare.cost {
					v.gridSquare.cost = newCost
					v.prev = minNode
				}
			}
		}
		haveAllTargets := true
		for _, target := range targets {
			if target.gridSquare.cost == utils.MaxInt {
				haveAllTargets = false
				break
			}
		}
		if haveAllTargets {
			break
		}
	}

	sort.Slice(targets, func(i, j int) bool {
		return compareVertices(targets[i], targets[j])
	})
	closestTarget := targets[0]

	if closestTarget.gridSquare.cost == utils.MaxInt {
		return
	}
	chain := closestTarget

	debug("Starting to loop through chain:\n\t%+v\n\t%+v\n\t%+v\n\t%+v\n", chain, chain.prev, chain.prev.gridSquare, start)
	for {
		previous := chain.prev
		if previous == nil {
			return
			// panic("Nil previous")
		}
		previousGrid := previous.gridSquare
		if previousGrid == nil {
			panic("Nil grid")
		}
		if previousGrid == start {
			break
		}
		chain = previous
	}
	debug("Found start: %+v\n", start.occupiedBy)
	occupant := start.occupiedBy
	start.occupiedBy = nil
	chain.gridSquare.occupiedBy = occupant
	occupant.location = chain.gridSquare
}

func compareVertices(i, j *vertex) bool {
	lowerCost := i.gridSquare.cost < j.gridSquare.cost
	equalCost := i.gridSquare.cost == j.gridSquare.cost
	lowerReadingOrder := i.y < j.y || (i.y == j.y && i.x < j.x)
	return lowerCost || (equalCost && lowerReadingOrder)
}

type vertex struct {
	point

	prev       *vertex
	gridSquare *gridSquare
}

type point struct {
	x, y int
}

func (p *point) manhattenDistance(t interface{}) int {
	xDist := utils.MaxInt
	yDist := utils.MaxInt
	switch t.(type) {
	case point:
		tPoint := t.(point)
		xDist = int(math.Abs(float64(p.x - tPoint.x)))
		yDist = int(math.Abs(float64(p.y - tPoint.y)))
	case vertex:
		tPoint := t.(vertex)
		xDist = int(math.Abs(float64(p.x - tPoint.x)))
		yDist = int(math.Abs(float64(p.y - tPoint.y)))
	case *vertex:
		tPoint := t.(*vertex)
		xDist = int(math.Abs(float64(p.x - tPoint.x)))
		yDist = int(math.Abs(float64(p.y - tPoint.y)))
	default:
		panic(fmt.Errorf("Unsupported target type: %v", reflect.TypeOf(t)))
	}

	return xDist + yDist
}
