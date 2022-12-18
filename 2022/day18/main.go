package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	allCubes := loadCubes()
	p1 := part1(allCubes)
	p2 := part2(allCubes)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(allCubes cubeSet) int {
	area := 0
	faces := [6]point{
		{1, 0, 0}, {-1, 0, 0},
		{0, 1, 0}, {0, -1, 0},
		{0, 0, 1}, {0, 0, -1},
	}
	for cube := range allCubes {
		for _, face := range faces {
			if allCubes[cube.add(face)] {
				continue
			}
			area++
		}
	}
	return area
}

func part2(allCubes cubeSet) int {
	area := 0
	min, max := point{999, 999, 999}, point{}
	for c := range allCubes {
		min = min.minBound(c)
		max = max.maxBound(c)
	}
	min = point{min.x - 1, min.y - 1, min.z - 1}
	max = point{max.x + 1, max.y + 1, max.z + 1}
	explored := map[point]bool{min: true}
	toExplore := []point{min}
	nextExplore := []point{}
	for len(toExplore) > 0 {
		nextExplore = nextExplore[0:0]
		for _, exploreFrom := range toExplore {
			for _, adj := range exploreFrom.adjacentPoints() {
				if allCubes[adj] {
					area++
					continue
				}
				if explored[adj] {
					continue
				}
				explored[adj] = true
				if min != min.minBound(adj) || max != max.maxBound(adj) {
					continue
				}
				nextExplore = append(nextExplore, adj)
			}
		}
		toExplore, nextExplore = nextExplore, toExplore
	}
	return area
}

type cubeSet map[point]bool

type point struct{ x, y, z int }

func (p point) minBound(o point) point { return point{min(p.x, o.x), min(p.y, o.y), min(p.z, o.z)} }
func (p point) maxBound(o point) point { return point{max(p.x, o.x), max(p.y, o.y), max(p.z, o.z)} }
func (p point) add(o point) point      { return point{p.x + o.x, p.y + o.y, p.z + o.z} }
func (p point) adjacentPoints() []point {
	adjacent := make([]point, 6)
	faces := [6]point{
		{1, 0, 0}, {-1, 0, 0},
		{0, 1, 0}, {0, -1, 0},
		{0, 0, 1}, {0, 0, -1},
	}
	for i, face := range faces {
		adjacent[i] = p.add(face)
	}
	return adjacent
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func loadCubes() cubeSet {
	cubes := cubeSet{}
	for _, line := range utils.GetLines(input) {
		points := utils.GetInts(line)
		cubes[point{points[0], points[1], points[2]}] = true
	}
	return cubes
}

var benchmark = false
