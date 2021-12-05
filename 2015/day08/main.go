package main

import (
	"fmt"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	lines := utils.ReadInputLines("input.txt")
	diff := 0
	for _, line := range lines {
		diff += len(line)
		line, _ = strconv.Unquote(line)
		diff -= len(line)
	}

	return diff
}

func part2() int {
	lines := utils.ReadInputLines("input.txt")
	diff := 0
	for _, line := range lines {
		diff -= len(line)
		line = fmt.Sprintf("%+q", line)
		diff += len(line)
	}

	return diff
}

var benchmark = false
