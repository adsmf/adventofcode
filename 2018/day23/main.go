package main

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	bots := load("input.txt")
	p1 := part1(bots)
	p2 := part2(bots)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(bots []nanobotInfo) int {
	inRange := [][]int{}
	bigR := 0
	bigRBot := 0
	for sID, sInfo := range bots {
		inRangeOfS := []int{}
		if sInfo.r > bigR {
			bigR = sInfo.r
			bigRBot = sID
		}
		for tID, tInfo := range bots {
			if sInfo.pos.distTo(tInfo.pos) <= sInfo.r {
				inRangeOfS = append(inRangeOfS, tID)
			}
		}
		inRange = append(inRange, inRangeOfS)
	}
	return len(inRange[bigRBot])
}

func part2(bots []nanobotInfo) int {
	overlaps := &overlapHeap{}
	heap.Init(overlaps)
	for _, info := range bots {
		d := info.pos.distTo(point{0, 0, 0})
		heap.Push(overlaps, overlapInfo{
			dist: int(math.Max(0, float64(d-info.r))),
			inc:  1,
		})
		heap.Push(overlaps, overlapInfo{
			dist: d + info.r + 1,
			inc:  -1,
		})
	}
	count, bestCount, minDist := 0, 0, 0
	for overlaps.Len() > 0 {
		info := heap.Pop(overlaps).(overlapInfo)
		count += info.inc
		if count > bestCount {
			bestCount = count
			minDist = info.dist
		}
	}
	return minDist
}

type overlapHeap []overlapInfo

func (h overlapHeap) Len() int            { return len(h) }
func (h overlapHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h overlapHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *overlapHeap) Push(x interface{}) { *h = append(*h, x.(overlapInfo)) }

func (h *overlapHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type overlapInfo struct {
	dist int
	inc  int // Marker for start/end of overlap
}

type nanobotInfo struct {
	pos point
	r   int
}

type point struct{ x, y, z int }

func (p point) distTo(a point) int {
	dX, dY, dZ := p.x-a.x, p.y-a.y, p.z-a.z

	if dX < 0 {
		dX *= -1
	}
	if dY < 0 {
		dY *= -1
	}
	if dZ < 0 {
		dZ *= -1
	}

	return dX + dY + dZ
}

func load(filename string) []nanobotInfo {
	bots := []nanobotInfo{}

	for _, line := range utils.ReadInputLines(filename) {
		ints := utils.GetInts(line)
		bots = append(bots, nanobotInfo{
			pos: point{ints[0], ints[1], ints[2]},
			r:   ints[3],
		})
	}

	return bots
}

var benchmark = false
