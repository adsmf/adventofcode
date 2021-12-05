package main

import (
	"fmt"
)

var benchmark = false

func main() {
	row, col := 2947, 3029
	p1 := part1(20151125, col, row)
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(seed, col, row int) int {
	n := seqNum(col, row)
	val := seed
	for i := 0; i < n-1; i++ {
		val *= 252533
		val %= 33554393
	}

	return val
}

func part2() int {
	return -1
}

func seqNum(col, row int) int {
	if row == 1 {
		return col*(col+1)/2 + row - 1
	}
	return seqNum(col+row-1, 1) - (row - 1)
}
