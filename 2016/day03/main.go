package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	rt, ct := load("input.txt")
	p1 := countValid(rt)
	p2 := countValid(ct)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func countValid(triangles []triangle) int {
	count := 0
	for _, triangle := range triangles {
		if triangle.valid() {
			count++
		}
	}
	return count
}

func load(filename string) ([]triangle, []triangle) {
	rowTriangles := []triangle{}
	colTriangles := []triangle{}
	bytes, _ := ioutil.ReadFile(filename)
	allInts := utils.GetInts(string(bytes))
	for i := 0; i < len(allInts)-8; i += 9 {
		for j := 0; j < 3; j++ {
			rInts := append([]int{}, allInts[i+j*3:i+j*3+3]...)
			cInts := []int{allInts[i+j], allInts[i+j+3], allInts[i+j+6]}
			sort.Ints(rInts)
			sort.Ints(cInts)
			rowTriangles = append(rowTriangles, triangle{rInts[0], rInts[1], rInts[2]})
			colTriangles = append(colTriangles, triangle{cInts[0], cInts[1], cInts[2]})
		}
	}
	return rowTriangles, colTriangles
}

type triangle struct{ a, b, c int }

func (t triangle) valid() bool {
	return t.c < t.a+t.b
}
