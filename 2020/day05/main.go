package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
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
	part2 := -1
	for try := min; try < max; try++ {
		if _, found := passes[try]; !found {
			part2 = try
			break
		}
	}

	fmt.Printf("Part 1: %d\n", max)
	fmt.Printf("Part 2: %d\n", part2)
}

func load(filename string) boardingPasses {
	passes := boardingPasses{}
	lines := utils.ReadInputLines(filename)
	mapping := map[rune]rune{'B': '1', 'F': '0', 'R': '1', 'L': '0'}
	for _, line := range lines {
		number := strings.Map(func(c rune) rune { return mapping[c] }, line)
		pos, err := strconv.ParseInt(number, 2, 16)
		if err != nil {
			panic(err)
		}
		passes[int(pos)] = true
	}
	return passes
}

type boardingPasses map[int]bool
