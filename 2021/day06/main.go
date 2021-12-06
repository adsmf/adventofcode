package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := runSim()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func runSim() (int, int) {
	fish := utils.GetInts(input)
	fishCounts := make([]int, 9, 256+9)
	for _, f := range fish {
		fishCounts[f]++
	}
	p1 := 0
	for day := 0; day < 256; day++ {
		if day == 80 {
			p1 = sum(fishCounts)
		}
		nextDay := append(fishCounts[1:], fishCounts[0])
		nextDay[6] += fishCounts[0]
		fishCounts = nextDay
	}
	return p1, sum(fishCounts)
}

func sum(counts []int) int {
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

var benchmark = false
