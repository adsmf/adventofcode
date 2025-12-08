package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

const closestN = 1000

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	allPoints := make([]point, 0, 1000)
	curPoint := point{}
	utils.EachInteger(input, func(index, value int) (done bool) {
		switch index % 3 {
		case 0:
			curPoint.x = value
		case 1:
			curPoint.y = value
		case 2:
			curPoint.z = value
			allPoints = append(allPoints, curPoint)
		}
		return
	})
	type distInfo struct {
		p1, p2   point
		distance float64
	}
	sortedDists := []distInfo{}
	for idx1 := 0; idx1 < len(allPoints)-1; idx1++ {
		pos1 := allPoints[idx1]
		for idx2 := idx1 + 1; idx2 < len(allPoints); idx2++ {
			pos2 := allPoints[idx2]
			dist := pos1.sub(pos2).distance()
			sortedDists = append(sortedDists, distInfo{pos1, pos2, dist})
		}
	}
	slices.SortFunc(sortedDists, func(a, b distInfo) int {
		return int(a.distance*1000 - b.distance*1000)
	})
	part1 := 0
	part2 := 0
	groups := pointSets{}
	for i := 0; i < len(sortedDists); i++ {
		pair := sortedDists[i]
		g1 := groups.group(pair.p1)
		g2 := groups.group(pair.p2)
		if g1 == -1 && g2 == -1 {
			newGroup := pointSet{
				pair.p1: true,
				pair.p2: true,
			}
			groups = append(groups, newGroup)
		} else if g1 == g2 {
			// Skip
		} else if g2 == -1 {
			groups[g1][pair.p2] = true
		} else if g1 == -1 {
			groups[g2][pair.p1] = true
		} else {
			for pos := range groups[g2] {
				groups[g1][pos] = true
			}
			copy(groups[g2:], groups[g2+1:])
			groups = groups[:len(groups)-1]
			if len(groups) == 1 {
				part2 = pair.p1.x * pair.p2.x
			}
		}
		if i == closestN-1 {

			slices.SortFunc(groups, func(a, b pointSet) int {
				return len(b) - len(a)
			})
			part1 = len(groups[0]) * len(groups[1]) * len(groups[2])
		}
	}
	return part1, part2
}

type pointSets []pointSet

func (p pointSets) group(check point) int {
	for idx, group := range p {
		if group.has(check) {
			return idx
		}
	}
	return -1
}

type pointSet map[point]bool

func (p pointSet) has(check point) bool {
	_, found := p[check]
	return found
}

type point struct{ x, y, z int }

func (p point) sub(q point) point { return point{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point) distance() float64 {
	return math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y) + float64(p.z*p.z))
}

var benchmark = false
