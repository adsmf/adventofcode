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
	cubes := make(map[cubeBounds]int8, len(ranges))
	type mergeEntry struct {
		cube  cubeBounds
		value int8
	}
	toMerge := make([]mergeEntry, 0, len(cubes)*10)
	for _, r := range ranges {
		toMerge = toMerge[0:0]
		for prev, value := range cubes {
			intersect := prev.intersection(r.cube)
			if intersect != nil {
				toMerge = append(toMerge, mergeEntry{*intersect, value})
			}
		}
		for _, entry := range toMerge {
			value := cubes[entry.cube]
			if value == entry.value {
				delete(cubes, entry.cube)
			} else {
				cubes[entry.cube] -= entry.value
			}
		}
		if r.on {
			cubes[r.cube]++
		}
	}

	p1, p2 := 0, 0
	fiftyCube := cubeBounds{point{-50, -50, -50}, point{50, 50, 50}}
	for cube, value := range cubes {
		p1 += cube.intersection(fiftyCube).volume() * int(value)
		p2 += cube.volume() * int(value)
	}
	return p1, p2
}

func loadRanges(input string) []gridRange {
	lines := utils.GetLines(input)
	r := make([]gridRange, 0, len(lines))
	for _, line := range lines {
		vals := utils.GetInts(line)
		newRange := gridRange{
			cube: cubeBounds{
				min: point{pointVal(vals[0]), pointVal(vals[2]), pointVal(vals[4])},
				max: point{pointVal(vals[1]), pointVal(vals[3]), pointVal(vals[5])},
			},
			on: line[1] == 'n',
		}
		r = append(r, newRange)
	}
	return r
}

type gridRange struct {
	cube cubeBounds
	on   bool
}

type cubeBounds struct {
	min, max point
}

func (c cubeBounds) intersection(o cubeBounds) *cubeBounds {
	if c.min.x > o.max.x || c.max.x < o.min.x ||
		c.min.y > o.max.y || c.max.y < o.min.y ||
		c.min.z > o.max.z || c.max.z < o.min.z {
		return nil
	}
	return &cubeBounds{
		min: point{max(c.min.x, o.min.x), max(c.min.y, o.min.y), max(c.min.z, o.min.z)},
		max: point{min(c.max.x, o.max.x), min(c.max.y, o.max.y), min(c.max.z, o.max.z)},
	}
}

func (c *cubeBounds) volume() int {
	if c == nil {
		return 0
	}
	return int(c.max.x-c.min.x+1) * int(c.max.y-c.min.y+1) * int(c.max.z-c.min.z+1)
}

func max(a, b pointVal) pointVal {
	if a > b {
		return a
	}
	return b
}
func min(a, b pointVal) pointVal {
	if a < b {
		return a
	}
	return b
}

type pointVal int32
type point struct{ x, y, z pointVal }

var benchmark = false
