package main

import (
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode/utils"
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
	c := loadCuboids("input.txt")
	return c.paperAll()
}

func part2() int {
	c := loadCuboids("input.txt")
	return c.ribbonAll()
}

func loadCuboids(filename string) cuboidList {
	cuboids := cuboidList{}
	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		c := newCuboid(line)
		cuboids = append(cuboids, c)
	}
	return cuboids
}

type cuboidList []cuboid

func (c cuboidList) paperAll() int {
	total := 0

	for _, cube := range c {
		total += cube.paper()
	}
	return total
}

func (c cuboidList) ribbonAll() int {
	total := 0

	for _, cube := range c {
		total += cube.ribbon()
	}
	return total
}

type cuboid struct {
	w, l, h int
}

func (c cuboid) paper() int {
	areas := []int{c.w * c.l, c.w * c.h, c.h * c.l}
	sort.Ints(areas)
	return 3*areas[0] + 2*areas[1] + 2*areas[2]
}

func (c cuboid) ribbon() int {
	dims := []int{c.w, c.l, c.h}
	sort.Ints(dims)
	return 2*dims[0] + 2*dims[1] + c.w*c.h*c.l
}

func newCuboid(def string) cuboid {
	dims := utils.GetInts(def)
	if len(dims) != 3 {
		panic(fmt.Errorf("Wrong number of dimensions (%d) in '%s'", len(dims), def))
	}
	return cuboid{
		dims[0],
		dims[1],
		dims[2],
	}
}

var benchmark = false
