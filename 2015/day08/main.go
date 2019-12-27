package main

import (
	"fmt"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
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
	return -1
}
