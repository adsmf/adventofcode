package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	initial := loadInitial("input.txt")
	p1 := part1(initial)
	p2 := part2(initial)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(dim dimension) int {
	for i := 0; i < 6; i++ {
		dim = dim.next(false)
	}
	return len(dim)
}

func part2(dim dimension) int {
	for i := 0; i < 6; i++ {
		dim = dim.next(true)
	}
	return len(dim)
}

type dimension map[point]bool
type dimensionCount map[point]int

func (d dimension) next(hyper bool) dimension {
	next := dimension{}
	counts := dimensionCount{}
	for pos := range d {
		neighbours := pos.neighbours(hyper)
		for _, neighbour := range neighbours {
			counts[neighbour]++
		}
	}
	for pos, count := range counts {
		if (count == 3) || (d[pos] && count == 2) {
			next[pos] = true
		}
	}
	return next
}

type point struct{ x, y, z, w int }

func (p point) neighbours(hyper bool) []point {
	minW, maxW := 0, 0
	if hyper {
		minW = -1
		maxW = 1
	}
	points := []point{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := minW; w <= maxW; w++ {
					if x == 0 && y == 0 && z == 0 && w == 0 {
						continue
					}
					points = append(points, point{p.x + x, p.y + y, p.z + z, p.w + w})
				}
			}
		}
	}
	return points
}

func loadInitial(filename string) dimension {
	inputBytes, _ := ioutil.ReadFile(filename)
	dim := dimension{}
	x, y := 0, 0
	for _, char := range inputBytes {
		switch char {
		case '#':
			dim[point{x, y, 0, 0}] = true
			x++
		case '.':
			x++
		case '\n':
			x = 0
			y++
		}
	}
	return dim
}

var benchmark = false
