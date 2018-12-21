package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"

	"github.com/adsmf/adventofcode2018/utils"
)

func moveToBestTarget(cavern grid, start *gridSquare) grid {
	verticies := []*vertex{}

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

		}
	}
	start.cost = 0

	for len(verticies) > 0 {
		sort.Slice(verticies, func(i, j int) bool {
			lowerCost := verticies[i].gridSquare.cost < verticies[j].gridSquare.cost
			equalCost := verticies[i].gridSquare.cost == verticies[j].gridSquare.cost
			lowerReadingOrder := verticies[i].y < verticies[j].y || (verticies[i].y == verticies[j].y && verticies[i].x < verticies[j].x)
			return lowerCost || (equalCost && lowerReadingOrder)
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
	}
	return cavern
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
