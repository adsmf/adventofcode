package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := findIncreasesSinglePass()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func findIncreasesSinglePass() (int, int) {
	inc1, inc2 := 0, 0
	accumulator := 0
	w1, w2, w3 := 99999, 99999, 99999
	for _, ch := range input {
		switch {
		case ch >= '0' && ch <= '9':
			accumulator = accumulator*10 + int(ch-'0')
		case ch == '\n':
			if accumulator > w1 {
				inc1++
			}
			if accumulator > w3 {
				inc2++
			}
			w1, w2, w3 = accumulator, w1, w2
			accumulator = 0
		}
	}
	return inc1, inc2
}

// Obsolete implementation follows - present for benchmark comparison

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
