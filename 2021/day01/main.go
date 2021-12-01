package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	depths := getDepths()
	p1 := findIncreases(1, depths)
	p2 := findIncreases(3, depths)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func getDepths() []int {
	depths := []int{}
	last := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			depths = append(depths, utils.MustInt(input[last:i]))
			last = i + 1
		}
	}
	return depths
}

func findIncreases(windowSize int, depths []int) int {
	lastSum := 0
	increases := 0
	for i := 0; i < len(depths); i++ {
		depth := depths[i]
		if i < windowSize {
			lastSum += depth
			continue
		}
		sum := lastSum + depth - depths[i-windowSize]
		if sum > lastSum {
			increases++
		}
		lastSum = sum
	}
	return increases
}

var benchmark = false
