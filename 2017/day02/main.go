package main

import (
	"fmt"

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
	checksum := 0
	for _, line := range utils.ReadInputLines("input.txt") {
		min, max := utils.MaxInt, 0
		for _, val := range utils.GetInts(line) {
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}
		checksum += max - min
	}
	return checksum
}

func part2() int {
	checksum := 0
	for _, line := range utils.ReadInputLines("input.txt") {
		ints := utils.GetInts(line)
		for i := 0; i < len(ints)-1; i++ {
			found := false
			for j := i + 1; j < len(ints); j++ {
				int1, int2 := ints[i], ints[j]
				if int1 > int2 {
					int1, int2 = int2, int1
				}
				if int2%int1 == 0 {
					checksum += int2 / int1
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}
	return checksum
}

var benchmark = false
