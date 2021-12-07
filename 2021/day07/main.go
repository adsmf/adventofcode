package main

import (
	_ "embed"
	"fmt"
	"math"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := calcCosts()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func calcCosts() (int, int) {
	positions := utils.GetInts(input)
	minPos := math.MaxInt32
	maxPos := 0
	for _, pos := range positions {
		if minPos > pos {
			minPos = pos
		}
		if maxPos < pos {
			maxPos = pos
		}
	}
	costsP1 := map[int]int{}
	costsP2 := map[int]int{}

	for target := minPos; target <= maxPos; target++ {
		for _, c2 := range positions {
			dist := int(math.Abs(float64(target - c2)))
			p2cost := 0
			for i := 1; i <= dist; i++ {
				p2cost += i
			}
			costsP1[target] += dist
			costsP2[target] += p2cost
		}
	}
	minP1Cost := math.MaxInt32
	for _, cost := range costsP1 {
		if cost < minP1Cost {
			minP1Cost = cost
		}
	}
	minP2Cost := math.MaxInt32
	for _, cost := range costsP2 {
		if cost < minP2Cost {
			minP2Cost = cost
		}
	}
	return minP1Cost, minP2Cost
}

var benchmark = false
