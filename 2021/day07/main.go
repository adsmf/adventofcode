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

// Alternative implementations
func calcCostsSlice() (int, int) {
	positions := utils.GetInts(input)
	occupied := make(map[int]int, len(positions))
	for _, pos := range positions {
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
	costsP1 := make([]int, maxPos+1)
	costsP2 := make([]int, maxPos+1)

	for target := minPos; target <= maxPos; target++ {
		for c2, count := range occupied {
			dist := int(math.Abs(float64(target - c2)))
			p2cost := distCost(dist)
			costsP1[target] += dist * count
			costsP2[target] += p2cost * count
		}
	}
	minP1Cost := math.MaxInt32
	for target := minPos; target <= maxPos; target++ {
		cost := costsP1[target]
		if cost < minP1Cost {
			minP1Cost = cost
		}
	}
	minP2Cost := math.MaxInt32
	for target := minPos; target <= maxPos; target++ {
		cost := costsP2[target]
		if cost < minP2Cost {
			minP2Cost = cost
		}
	}
	return minP1Cost, minP2Cost
}

func calcCostsDedup() (int, int) {
	positions := utils.GetInts(input)
	occupied := make(map[int]int, len(positions))
	for _, pos := range positions {
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
	costsP1 := map[int]int{}
	costsP2 := map[int]int{}

	for target := minPos; target <= maxPos; target++ {
		for c2, count := range occupied {
			dist := int(math.Abs(float64(target - c2)))
			p2cost := distCost(dist)
			costsP1[target] += dist * count
			costsP2[target] += p2cost * count
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

func calcCostsInitial() (int, int) {
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
			p2cost := distCost(dist)
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
