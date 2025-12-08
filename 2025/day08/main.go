package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"

	"github.com/adsmf/adventofcode/utils"
)

// //go:embed example1.txt
// var input string

// const closestN = 10

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
		idx1, idx2 int
		distance   float64
	}
	sortedDists := make([]distInfo, 0, 500000)
	for idx1 := 0; idx1 < len(allPoints)-1; idx1++ {
		pos1 := allPoints[idx1]
		for idx2 := idx1 + 1; idx2 < len(allPoints); idx2++ {
			pos2 := allPoints[idx2]
			dist := pos1.sub(pos2).distance()
			sortedDists = append(sortedDists, distInfo{idx1, idx2, dist})
		}
	}
	slices.SortFunc(sortedDists, func(a, b distInfo) int {
		return int(a.distance*1000 - b.distance*1000)
	})
	part1 := 0
	part2 := 0
	groups := make(pointSets, 0, 350)
	for i := 0; i < len(sortedDists); i++ {
		pair := sortedDists[i]
		g1 := groups.group(pair.idx1)
		g2 := groups.group(pair.idx2)
		if g1 == -1 && g2 == -1 {
			newGroup := pointSet{}
			newGroup[pair.idx1] = true
			newGroup[pair.idx2] = true
			groups = append(groups, newGroup)
		} else if g1 == g2 {
			// Skip
		} else if g2 == -1 {
			groups[g1][pair.idx2] = true
		} else if g1 == -1 {
			groups[g2][pair.idx1] = true
		} else {
			for pos, val := range groups[g2] {
				if val {
					groups[g1][pos] = true
				}
			}
			copy(groups[g2:], groups[g2+1:])
			groups = groups[:len(groups)-1]
			if len(groups) == 1 {
				part2 = allPoints[pair.idx1].x * allPoints[pair.idx2].x
			}
		}
		if i == closestN-1 {
			slices.SortFunc(groups[:], func(a, b pointSet) int {
				return b.count() - a.count()
			})
			part1 = groups[0].count() * groups[1].count() * groups[2].count()
		}
	}
	return part1, part2
}

type pointSets []pointSet

func (p pointSets) group(check int) int {
	for idx, group := range p {
		if group.has(check) {
			return idx
		}
	}
	return -1
}

type pointSet [1000]bool

func (p pointSet) has(check int) bool {
	return p[check]
}

func (p pointSet) count() int {
	c := 0
	for _, v := range p {
		if v {
			c++
		}
	}
	return c
}

type point struct{ x, y, z int }

func (p point) sub(q point) point { return point{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point) distance() float64 {
	return math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y) + float64(p.z*p.z))
}

var benchmark = false
