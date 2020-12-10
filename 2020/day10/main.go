package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	adapters := utils.GetInts(string(inputBytes))
	p1 := part1(adapters)
	p2 := part2(adapters)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(adapters []int) int {
	sort.Ints(adapters)
	diffs := map[int]int{}
	last := 0
	for _, adapter := range adapters {
		diffs[adapter-last]++
		last = adapter
	}
	diffs[3]++
	return diffs[1] * diffs[3]
}

func part2(adapters []int) int {
	adapterMap := map[int]bool{}
	max := 0
	for _, adapter := range adapters {
		adapterMap[adapter] = true
		if adapter > max {
			max = adapter
		}
	}
	deviceJolts := max + 3
	adapterMap[deviceJolts] = true

	cache := map[string]int{}

	return search(0, deviceJolts, adapterMap, cache)
}

func search(jolt, deviceJolts int, adapters map[int]bool, cache map[string]int) int {
	if jolt == deviceJolts {
		return 1
	}
	count := 0
	for diff := 1; diff <= 3; diff++ {
		if adapters[jolt+diff] {
			key := fmt.Sprintf("%d,%d", jolt, diff)
			if _, found := cache[key]; !found {
				subCount := search(jolt+diff, deviceJolts, adapters, cache)
				cache[key] = subCount
			}
			count += cache[key]
		}
	}
	return count
}

var benchmark = false
