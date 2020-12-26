package main

import (
	"fmt"
)

const base = "ljoxqyyw"

func main() {
	// p1 := part1()
	p1, p2 := countRegions(base)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func countRegions(seed string) (int, int) {
	pointCount := 0
	regions := []region{}
	for y := 0; y < 128; y++ {
		row := hash(fmt.Sprintf("%s-%d", seed, y))
		for x, c := range row {
			if c == '1' {
				pointCount++
				adjacentRegions := []int{}
				for regID, region := range regions {
					for _, dir := range dirs {
						neighbourPoint := point{x + dir.x, y + dir.y}
						if region[neighbourPoint] {
							adjacentRegions = append(adjacentRegions, regID)
						}
					}
				}
				pos := point{x, y}
				switch len(adjacentRegions) {
				case 0:
					regions = append(regions, region{pos: true})
				case 1:
					regions[adjacentRegions[0]][pos] = true
				default:
					mergedRegion := region{pos: true}
					skipRegions := map[int]bool{}
					for _, id := range adjacentRegions {
						skipRegions[id] = true
						for regPoint := range regions[id] {
							mergedRegion[regPoint] = true
						}
					}
					newRegions := make([]region, 0, len(regions)-len(adjacentRegions)+1)
					for regID, region := range regions {
						if !skipRegions[regID] {
							newRegions = append(newRegions, region)
						}
					}
					newRegions = append(newRegions, mergedRegion)
					regions = newRegions
				}
			}
		}
	}

	return pointCount, len(regions)
}

type point struct{ x, y int }

var dirs = []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type region map[point]bool

func hash(input string) string {
	lengths := []byte(input)
	lengths = append(lengths, 17, 31, 73, 47, 23)

	loop := newRing(256)

	skip := 0
	for iter := 0; iter < 64; iter++ {
		for _, length := range lengths {
			loop.twist(int(length))
			loop.skip(skip)
			skip++
		}
	}
	hash := ""
	for i := 0; i < 256; i += 16 {
		block := 0
		for j := 0; j < 16; j++ {
			block ^= loop.entries[i+j]
		}
		hash = fmt.Sprintf("%s%08b", hash, block)
	}
	return hash
}

func newRing(length int) ring {
	r := ring{
		entries: make([]int, length),
		offset:  0,
	}
	for i := 0; i < length; i++ {
		r.entries[i] = i
	}
	return r
}

type ring struct {
	entries []int
	offset  int
}

func (r *ring) twist(length int) {
	for i := 0; i < length/2; i++ {
		o1 := r.offset + i
		o2 := r.offset + length - i - 1
		o1 %= len(r.entries)
		o2 %= len(r.entries)
		r.entries[o1], r.entries[o2] = r.entries[o2], r.entries[o1]
	}

	r.skip(length)
}

func (r *ring) skip(n int) {
	r.offset = (r.offset + n) % len(r.entries)
}

var benchmark = false
