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
	p1, p2 := calcCostsNoSort()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func calcCostsNoSort() (int, int) {
	positions := utils.GetInts(input)
	occupied := make(map[int]int, len(positions))
	totalPos := 0
	for _, pos := range positions {
		totalPos += pos
		occupied[pos]++
	}
	costP1 := 0
	costP2 := 0

	median := 0
	halfway := len(positions) / 2
	for i, medianPos := 0, 0; medianPos < halfway; i++ {
		medianPos += occupied[i]
		median = i
	}
	mean := totalPos / len(positions)
	for from, count := range occupied {
		costP1 += int(math.Abs(float64(median-from))) * count
		costP2 += distCost(int(math.Abs(float64(mean-from)))) * count
	}

	return costP1, costP2
}

func distCost(dist int) int {
	return dist * (dist + 1) / 2
}

var benchmark = false
