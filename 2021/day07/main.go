package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := calcCostsTargeted()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

// Test zone
func calcCostsTargeted() (int, int) {
	positions := utils.GetInts(input)
	occupied := make(map[int]int, len(positions))
	totalPos := 0
	for _, pos := range positions {
		totalPos += pos
		occupied[pos]++
	}
	minPos := math.MaxInt32
	maxPos := 0
	for pos := range occupied {
		if minPos > pos {
			minPos = pos
		}
		if maxPos < pos {
			maxPos = pos
		}
	}
	costP1 := 0
	costP2 := 0

	sort.Ints(positions)
	medianPos := positions[len(positions)/2]
	for from, count := range occupied {
		dist := int(math.Abs(float64(medianPos - from)))
		costP1 += dist * count
	}

	mean := totalPos / len(positions)
	for c2, count := range occupied {
		dist := int(math.Abs(float64(mean - c2)))
		p2cost := distCost(dist)
		costP2 += p2cost * count
	}

	return costP1, costP2
}

func distCost(dist int) int {
	return dist * (dist + 1) / 2
}

var benchmark = false
