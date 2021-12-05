package main

import (
	"fmt"
)

var input = 33100000

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	return equalDeliveries()
}

func part2() int {
	return equalDeliveries2()
}

func equalDeliveries() int {
	deliveries := make(map[int]int, input)

	for elf := 1; elf <= input/10; elf++ {
		for house := elf; house <= input/10; house += elf {
			deliveries[house] += elf * 10
		}
	}

	for house := 1; house <= input; house++ {
		if deliveries[house] >= input {
			return house
		}
	}
	return -1
}

func equalDeliveries2() int {
	deliveries := make(map[int]int, input)

	for elf := 1; elf <= input/10; elf++ {
		for house := elf; house <= input/10 && house <= elf*50; house += elf {
			deliveries[house] += elf * 11
		}
	}

	for house := 1; house <= input; house++ {
		if deliveries[house] >= input {
			return house
		}
	}
	return -1
}

var benchmark = false
