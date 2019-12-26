package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	c := loadCuboids("input.txt")
	return c.paperAll()
}

func part2() int {
	return 0
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

type cuboid struct {
	w, l, h int
}

func (c cuboid) paper() int {
	f1 := c.w * c.l
	f2 := c.w * c.h
	f3 := c.h * c.l
	min := f1
	if f2 < min {
		min = f2
	}
	if f3 < min {
		min = f3
	}

	return 2*f1 + 2*f2 + 2*f3 + min

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
