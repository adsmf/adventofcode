package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	part1, part2 := loadBitwise("input.txt")

	if !benchmark {
		fmt.Printf("Part 1: %d\n", part1)
		fmt.Printf("Part 2: %d\n", part2)
	}
}

func loadBitwise(filename string) (int, int) {
	lines := utils.ReadInputLines(filename)
	min, max, total := 1<<17, 0, 0
	for _, line := range lines {
		pass := 0
		for i := 0; i <= 9; i++ {
			pass |= (int(line[i]&(4)) >> 2) << (9 - i)
		}
		pass ^= (1<<10 - 1)
		total += pass
		if pass < min {
			min = pass
		}
		if pass > max {
			max = pass
		}
	}
	total -= (max * (max + 1) / 2) - (min * (min - 1) / 2)
	return max, total * -1
}
