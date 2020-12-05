package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	passes := load("input.txt")
	max := 0
	for pass := range passes {
		if pass > max {
			max = pass
		}
	}
	return max
}

func part2() int {
	passes := load("input.txt")
	min := 1 << 17
	max := 0
	for pass := range passes {
		if pass > max {
			max = pass
		}
		if pass < min {
			min = pass
		}
	}
	for try := min; try < max; try++ {
		if _, found := passes[try]; !found {
			return try
		}
	}
	return -1
}

func load(filename string) boardingPasses {
	passes := boardingPasses{}
	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "R", "1")
		line = strings.ReplaceAll(line, "L", "0")
		pos, err := strconv.ParseInt(line, 2, 16)
		if err != nil {
			panic(err)
		}
		passes[int(pos)] = true
	}
	return passes
}

type boardingPasses map[int]bool
