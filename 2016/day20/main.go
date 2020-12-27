package main

import (
	"fmt"
	"sort"

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
	ranges := load("input.txt")
	sort.Sort(ranges)
	lastEnd := 0
	for _, blocked := range ranges {
		if blocked[0] > lastEnd+1 {
			return lastEnd + 1
		}
		lastEnd = blocked[1]
	}
	return -1
}

func part2() int {
	ranges := load("input.txt")
	sort.Sort(ranges)
	lastEnd := 0
	count := 0
	for _, blocked := range ranges {
		if blocked[0] > lastEnd+1 {
			count += blocked[0] - lastEnd - 1
		}
		if blocked[1] > lastEnd {
			lastEnd = blocked[1]
		}
	}
	count += (1<<32 - 1) - lastEnd
	return count
}

type ipRangeList []ipRange

func (irl ipRangeList) Len() int           { return len(irl) }
func (irl ipRangeList) Less(a, b int) bool { return irl[a][0] < irl[b][0] }
func (irl ipRangeList) Swap(a, b int)      { irl[a], irl[b] = irl[b], irl[a] }

type ipRange [2]int

func load(filename string) ipRangeList {
	lines := utils.ReadInputLines(filename)
	ranges := make([]ipRange, len(lines))
	for i, line := range lines {
		ints := utils.GetInts(line)
		ranges[i] = ipRange{ints[0], ints[1]}
	}
	return ranges
}

var benchmark = false
