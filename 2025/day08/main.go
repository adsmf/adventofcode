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
	allPoints := []point{}
	utils.EachLine[string](input, func(index int, line string) (done bool) {
		vals := utils.GetInts(line)
		pos := point{vals[0], vals[1], vals[2]}
		allPoints = append(allPoints, pos)
		return
	})
	type distInfo struct {
		p1, p2   point
		distance float64
	}
	distances := map[int]map[int]float64{}
	sortedDists := []distInfo{}
	for idx1, pos1 := range allPoints {
		for idx2, pos2 := range allPoints {
			if idx1 == idx2 {
				continue
			}
			h1, h2 := pos1.hash(), pos2.hash()
			upper, lower := pos1, pos2
			if h1 > h2 {
				h1, h2 = h2, h1
				upper, lower = lower, upper
			}
			if distances[h1] == nil {
				distances[h1] = map[int]float64{}
			}
			if _, found := distances[h1][h2]; found {
				continue
			}
			dist := upper.sub(lower).distance()
			distances[h1][h2] = dist
			sortedDists = append(sortedDists, distInfo{upper, lower, dist})
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

type point struct {
	x, y, z int
}

func (p point) hash() int {
	return p.x * p.y * p.z
}

func (p point) sub(q point) point { return point{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point) distance() float64 {
	return math.Sqrt(float64(p.x*p.x) + float64(p.y*p.y) + float64(p.z*p.z))
}

var benchmark = false
