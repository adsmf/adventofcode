package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils"
	"io/ioutil"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	sizes := loadInput("input.txt")
	return combinations(150, sizes)
}

func part2() int {
	sizes := loadInput("input.txt")
	return minCombinations(150, sizes)
}

func combinations(target int, sizes []int) int {
	combs := 0

	for i := 0; i < (1 << len(sizes)); i++ {
		sum := 0

		for idx, val := range sizes {
			if i&(1<<idx) > 0 {
				sum += val
			}
		}

		if sum == target {
			combs++
		}
	}

	return combs
}

func minCombinations(target int, sizes []int) int {
	combs := map[int]int{}
	minContainers := -1

	for i := 0; i < (1 << len(sizes)); i++ {
		sum := 0
		containers := 0

		for idx, val := range sizes {
			if i&(1<<idx) > 0 {
				sum += val
				containers++
			}
		}

		if sum == target {
			combs[containers]++
			if minContainers == -1 || containers < minContainers {
				minContainers = containers
			}
		}
	}

	return combs[minContainers]
}

func loadInput(filename string) []int {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	sizes := utils.GetInts(string(raw))
	sort.Ints(sizes)
	return sizes
}
