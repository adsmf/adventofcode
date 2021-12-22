package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	r := loadRanges(input)
	p1, p2 := solve(r)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(ranges []gridRange) (int, int) {
	cubes := make([]gridRange, 0, len(ranges)*6)
	nextCubes := make([]gridRange, 0, len(ranges)*6)
	for _, r := range ranges {
		nextCubes = nextCubes[0:0]
		for _, cube := range cubes {
			nextCubes = append(nextCubes, cube.disjoint(r)...)
		}
		if r.on {
			nextCubes = append(nextCubes, r)
		}
		cubes, nextCubes = nextCubes, cubes
	}

	p1, p2 := 0, 0
	fiftyCube := gridRange{point{-50, -50, -50}, point{50, 50, 50}, true}
	for _, cube := range cubes {
		p1 += cube.intersection(fiftyCube).volume()
		p2 += cube.volume()
	}
	return p1, p2
}

func loadRanges(input string) []gridRange {
	lines := utils.GetLines(input)
	r := make([]gridRange, 0, len(lines))
	for _, line := range lines {
		vals := utils.GetInts(line)
		newRange := gridRange{
			min: point{vals[0], vals[2], vals[4]},
			max: point{vals[1], vals[3], vals[5]},
			on:  line[1] == 'n',
		}
		r = append(r, newRange)
	}
	return r
}

type gridRange struct {
	min, max point
	on       bool
}

func (g gridRange) disjoint(o gridRange) []gridRange {
	intersect := g.intersection(o)
	if intersect == nil {
		return []gridRange{g}
	}
	subCubes := []gridRange{
		{min: point{g.min.x, g.min.y, g.min.z}, max: point{g.max.x, g.max.y, o.min.z - 1}, on: g.on},
		{min: point{g.min.x, g.min.y, o.max.z + 1}, max: point{g.max.x, g.max.y, g.max.z}, on: g.on},

		{min: point{g.min.x, g.min.y, max(g.min.z, o.min.z)}, max: point{g.max.x, o.min.y - 1, min(g.max.z, o.max.z)}, on: g.on},
		{min: point{g.min.x, o.max.y + 1, max(g.min.z, o.min.z)}, max: point{g.max.x, g.max.y, min(g.max.z, o.max.z)}, on: g.on},

		{min: point{g.min.x, max(g.min.y, o.min.y), max(g.min.z, o.min.z)}, max: point{o.min.x - 1, min(g.max.y, o.max.y), min(g.max.z, o.max.z)}, on: g.on},
		{min: point{o.max.x + 1, max(g.min.y, o.min.y), max(g.min.z, o.min.z)}, max: point{g.max.x, min(g.max.y, o.max.y), min(g.max.z, o.max.z)}, on: g.on},
	}

	n := 0
	for i := 0; i < len(subCubes); i++ {
		cube := subCubes[i]
		if cube.max.x < cube.min.x ||
			cube.max.y < cube.min.y ||
			cube.max.z < cube.min.z {
			continue
		}
		subCubes[n] = cube
		n++
	}

	subCubes = subCubes[:n]

	return subCubes
}

func (g gridRange) intersection(o gridRange) *gridRange {
	if g.min.x > o.max.x || g.max.x < o.min.x ||
		g.min.y > o.max.y || g.max.y < o.min.y ||
		g.min.z > o.max.z || g.max.z < o.min.z {
		return nil
	}
	return &gridRange{
		min: point{max(g.min.x, o.min.x), max(g.min.y, o.min.y), max(g.min.z, o.min.z)},
		max: point{min(g.max.x, o.max.x), min(g.max.y, o.max.y), min(g.max.z, o.max.z)},
		on:  g.on,
	}
}

func (g *gridRange) volume() int {
	if g == nil {
		return 0
	}
	return (g.max.x - g.min.x + 1) * (g.max.y - g.min.y + 1) * (g.max.z - g.min.z + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type point struct{ x, y, z int }

var benchmark = false
