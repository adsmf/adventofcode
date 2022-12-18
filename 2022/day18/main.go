package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

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
	minX, maxX := allCubes.min.x(), allCubes.max.x()
	minY, maxY := allCubes.min.y(), allCubes.max.y()
	minZ, maxZ := allCubes.min.z(), allCubes.max.z()
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				cube := pointAt(x, y, z)
				if !allCubes.vals[cube] {
					continue
				}
				for _, face := range faces {
					if allCubes.vals[cube.add(face)] {
						continue
					}
					area++
				}
			}
		}
	}
	return area
}

func part2(allCubes cubeSet) int {
	area := 0
	min := allCubes.min.add(pointAt(-1, -1, -1))
	max := allCubes.max.add(pointAt(1, 1, 1))
	explored := [1 << 15]bool{}
	explored[min] = true
	const exploreSize = 2500
	toExplore := [exploreSize]point{min}
	nextExplore := [exploreSize]point{}
	numExplore := 1
	for numExplore > 0 {
		numNext := 0
		for i := 0; i < numExplore; i++ {
			exploreFrom := toExplore[i]
			for _, adj := range exploreFrom.adjacent() {
				if allCubes.vals[adj] {
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
				nextExplore[numNext] = adj
				numNext++
			}
		}
		toExplore, nextExplore = nextExplore, toExplore
		numExplore = numNext
	}
	return area
}

type cubeSet struct {
	vals     [1 << 15]bool
	min, max point
}

const pointOffset = 2

func pointAt(x, y, z int) point {
	return point(
		(((x + pointOffset) & 0x1f) << 10) |
			(((y + pointOffset) & 0x1f) << 5) |
			((z + pointOffset) & 0x1f),
	)
}

type point int16

func (p point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x(), p.y(), p.z())
}

func (p point) x() int { return int((p>>10)&0x1f - pointOffset) }
func (p point) y() int { return int((p>>5)&0x1f - pointOffset) }
func (p point) z() int { return int((p)&0x1f - pointOffset) }

func (p point) minBound(o point) point {
	return pointAt(min(p.x(), o.x()), min(p.y(), o.y()), min(p.z(), o.z()))
}
func (p point) maxBound(o point) point {
	return pointAt(max(p.x(), o.x()), max(p.y(), o.y()), max(p.z(), o.z()))
}
func (p point) add(o point) point { return pointAt(p.x()+o.x(), p.y()+o.y(), p.z()+o.z()) }
func (p point) adjacent() [6]point {
	adjacent := [6]point{}
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
	cubes.min = pointAt(255, 255, 255)
	x, y, z := 0, 0, 0
	for pos := 0; pos < len(input); pos++ {
		x, pos = getInt(input, pos)
		y, pos = getInt(input, pos+1)
		z, pos = getInt(input, pos+1)
		cube := pointAt(x, y, z)
		cubes.min = cubes.min.minBound(cube)
		cubes.max = cubes.max.maxBound(cube)
		cubes.vals[cube] = true
	}
	return cubes
}

var faces = [6]point{
	pointAt(1, 0, 0), pointAt(-1, 0, 0),
	pointAt(0, 1, 0), pointAt(0, -1, 0),
	pointAt(0, 0, 1), pointAt(0, 0, -1),
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; in[pos]&0xf0 == 0x30; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

var benchmark = false
