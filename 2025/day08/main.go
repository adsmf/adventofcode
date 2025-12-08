package main

import (
	_ "embed"
	"fmt"
	"math"
	"math/bits"
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
		idx1, idx2 int
		distance   distType
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
		return int(a.distance - b.distance)
	})
	part1 := 0
	part2 := 0
	groups := make(pointSets, 0, 350)
	for i := 0; i < len(sortedDists); i++ {
		pair := sortedDists[i]
		g1 := groups.indexOf(pair.idx1)
		g2 := groups.indexOf(pair.idx2)
		if g1 == -1 && g2 == -1 {
			newGroup := pointSet{}
			newGroup = newGroup.with(pair.idx1)
			newGroup = newGroup.with(pair.idx2)
			groups = append(groups, newGroup)
		} else if g1 == g2 {
			// Skip
		} else if g2 == -1 {
			groups[g1] = groups[g1].with(pair.idx2)
		} else if g1 == -1 {
			groups[g2] = groups[g2].with(pair.idx1)
		} else {
			groups[g1] = groups[g1].merge(groups[g2])
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

func (p pointSets) indexOf(check int) int {
	for idx, group := range p {
		if group.has(check) {
			return idx
		}
	}
	return -1
}

type pointSet [16]uint64

func (p pointSet) with(index int) pointSet {
	high := index >> 6
	low := index % 64
	p[high] |= 1 << low
	return p
}
func (p pointSet) merge(o pointSet) pointSet {
	for h := range p {
		p[h] |= o[h]
	}
	return p
}
func (p pointSet) has(index int) bool {
	high := index >> 6
	low := index % 64
	return (p[high]>>low)&1 == 1
}

func (p pointSet) count() int {
	c := 0
	for _, v := range p {
		c += bits.OnesCount64(v)
	}
	return c
}

type point struct{ x, y, z int }

func (p point) sub(q point) point { return point{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point) distance() distType {
	return distType(math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y) + float64(p.z*p.z)))
}

type distType int64

var benchmark = false
