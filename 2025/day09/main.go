package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	redTiles := make([]point, 0, 500)
	curPoint := point{}
	utils.EachInteger(input, func(index, value int) (done bool) {
		switch index % 2 {
		case 0:
			curPoint.x = value
		case 1:
			curPoint.y = value
			redTiles = append(redTiles, curPoint)
		}
		return
	})
	edges := make([]hullEdge, 0, 500)
	for i := range len(redTiles) - 1 {
		edges = append(edges, [2]point{redTiles[i], redTiles[i+1]})
	}
	edges = append(edges, [2]point{redTiles[len(redTiles)-1], redTiles[0]})
	maxP1, maxP2 := 0, 0
	for i := 0; i < len(redTiles)-1; i++ {
		for j := i + 1; j < len(redTiles); j++ {
			pos1, pos2 := redTiles[i], redTiles[j]
			area := redTiles[i].sub(redTiles[j]).area()
			maxP1 = max(maxP1, area)
			if !withinBounds(edges, pos1, pos2) {
				continue
			}
			maxP2 = max(maxP2, area)
		}
	}
	return maxP1, maxP2
}

func withinBounds(edges []hullEdge, pos1, pos2 point) bool {
	minPos := point{min(pos1.x, pos2.x), min(pos1.y, pos2.y)}
	maxPos := point{max(pos1.x, pos2.x), max(pos1.y, pos2.y)}

	for _, edge := range edges {
		minEdge := point{min(edge[0].x, edge[1].x), min(edge[0].y, edge[1].y)}
		maxEdge := point{max(edge[0].x, edge[1].x), max(edge[0].y, edge[1].y)}
		if minEdge.y == maxEdge.y {
			if minPos.y >= maxEdge.y || maxPos.y <= minEdge.y {
				continue
			}
			if minEdge.x <= minPos.x && minPos.x < maxEdge.x {
				return false
			}
			if minEdge.x < maxPos.x && maxPos.x <= maxEdge.x {
				return false
			}
		} else {
			if minPos.x >= maxEdge.x || maxPos.x <= minEdge.x {
				continue
			}
			if minEdge.x <= minPos.y && minPos.y < maxEdge.x {
				return false
			}
			if minEdge.y < maxPos.y && maxPos.y <= maxEdge.y {
				return false
			}
		}
	}
	return true
}

type hullEdge [2]point

type point struct{ x, y int }

func (p point) sub(q point) point { return point{p.x - q.x, p.y - q.y} }
func (p point) area() int {
	return (utils.IntAbs(p.x) + 1) * (utils.IntAbs(p.y) + 1)
}

var benchmark = false
