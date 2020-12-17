package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	dim := loadInitial("input.txt")
	for i := 0; i < 6; i++ {
		dim = dim.next(false)
	}
	return dim.countActive()
}

func part2() int {
	dim := loadInitial("input.txt")
	for i := 0; i < 6; i++ {
		dim = dim.next(true)
	}
	return dim.countActive()
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

type dimension map[point]bool
type dimensionCount map[point]int

func (d dimension) countActive() int {
	count := 0
	for _, val := range d {
		if val {
			count++
		}
	}
	return count
}

func (d dimension) next(hyper bool) dimension {
	next := dimension{}
	min, max := d.bounds()

	counts := dimensionCount{}
	for x := min.x - 1; x <= max.x+1; x++ {
		for y := min.y - 1; y <= max.y+1; y++ {
			for z := min.z - 1; z <= max.z+1; z++ {
				for w := min.w - 1; w <= max.w+1; w++ {
					pos := point{x, y, z, w}
					if d[pos] {
						neighbours := pos.neighbours(hyper)
						for _, neighbour := range neighbours {
							counts[neighbour]++
						}
					}
				}
			}
		}
	}

	for x := min.x - 1; x <= max.x+1; x++ {
		for y := min.y - 1; y <= max.y+1; y++ {
			for z := min.z - 1; z <= max.z+1; z++ {
				for w := min.w - 1; w <= max.w+1; w++ {
					pos := point{x, y, z, w}
					if (counts[pos] == 3) || (d[pos] && counts[pos] == 2) {
						next[pos] = true
					}
				}
			}
		}
	}

	return next
}

func (d dimension) bounds() (point, point) {
	min, max := point{}, point{}
	for pos := range d {
		if pos.x < min.x {
			min.x = pos.x
		}
		if pos.x > max.x {
			max.x = pos.x
		}
		if pos.y < min.y {
			min.y = pos.y
		}
		if pos.y > max.y {
			max.y = pos.y
		}
		if pos.z < min.z {
			min.z = pos.z
		}
		if pos.z > max.z {
			max.z = pos.z
		}
		if pos.w < min.w {
			min.w = pos.w
		}
		if pos.w > max.w {
			max.w = pos.w
		}
	}
	return min, max
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

var benchmark = false
