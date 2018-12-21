package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

func moveTowardBestTarget(cavern grid, start *gridSquare) grid {
	verticies := []*vertex{}
	targets := []*vertex{}

	for y, row := range cavern {
		for x := range row {
			cavern[y][x].cost = utils.MaxInt
			if !!!cavern[y][x].isCavern {
				continue
			}
			isStart := &cavern[y][x] == start
			if cavern[y][x].occupiedBy != nil && !!!isStart {
				continue
			}
			newVertex := &vertex{
				point: point{
					x: x,
					y: y,
				},
				gridSquare: &cavern[y][x],
			}
			verticies = append(verticies, newVertex)
			if newVertex.gridSquare.isTarget {
				targets = append(targets, newVertex)
			}

		}
	}
	start.cost = 0

	for len(verticies) > 0 {
		sort.Slice(verticies, func(i, j int) bool {
			return vertexLess(verticies[i], verticies[j])
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
		return vertexLess(targets[i], targets[j])
	})
	closestTarget := targets[0]

	fmt.Println("Closest target: ", closestTarget)
	if closestTarget.gridSquare.cost == utils.MaxInt {
		return cavern
	}
	chain := closestTarget
	for chain.prev.gridSquare != start {
		fmt.Printf("Path %d,%d\n", chain.x, chain.y)
		chain = chain.prev
	}
	occupant := start.occupiedBy
	start.occupiedBy = nil
	chain.gridSquare.occupiedBy = occupant

	return cavern
}

func vertexLess(i, j *vertex) bool {
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
