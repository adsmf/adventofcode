package main

import (
	_ "embed"
	"fmt"
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
	depths := make([]int, 0, len(input)/4)
	accumulator := 0
	for _, ch := range input {
		if ch == '\n' {
			depths = append(depths, accumulator)
			accumulator = 0
			continue
		}
		accumulator *= 10
		accumulator += int(ch - '0')
	}
	return depths
}

func findIncreases(windowSize int, depths []int) int {
	increases := 0
	for i := windowSize; i < len(depths); i++ {
		if depths[i] > depths[i-windowSize] {
			increases++
		}
	}
	return increases
}

var benchmark = false
