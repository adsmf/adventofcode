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
	sev := 0
	for _, line := range utils.ReadInputLines("input.txt") {
		parts := utils.GetInts(line)
		dist := parts[0]
		depth := parts[1]
		mod := (depth - 1) * 2
		if mod == 0 || (dist+1)%mod == 1 {
			sev += dist * depth
		}
	}
	return sev
}

func part2() int {
	firewalls := map[int]int{}
	for _, line := range utils.ReadInputLines("input.txt") {
		parts := utils.GetInts(line)
		dist := parts[0]
		depth := parts[1]
		mod := (depth - 1) * 2
		firewalls[dist+1] = mod
	}
	var time int
	for time = 1; ; time++ {
		hit := false
		for dist, mod := range firewalls {
			if (dist+time)%mod == 1 {
				hit = true
				break
			}
		}
		if !hit {
			break
		}
	}
	return time
}

var benchmark = false
