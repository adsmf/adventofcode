package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	return -1
}

func part2() int {
	return -1
}

var benchmark = false
