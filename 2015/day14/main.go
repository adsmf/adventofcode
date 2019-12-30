package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	r := loadInput("input.txt")
	return r.distanceAfter(2503)
}

func part2() int {
	r := loadInput("input.txt")
	return r.scoreAfter(2503)
}

type race map[string]reindeerStats

func (r race) scoreAfter(seconds int) int {
	scores := map[string]int{}
	for i := 1; i <= seconds; i++ {
		distances := r.runFor(i)
		best := 0
		for _, dist := range distances {
			if dist > best {
				best = dist
			}
		}
		for name, dist := range distances {
			if dist == best {
				scores[name]++
			}
		}
	}
	best := 0
	for _, score := range scores {
		if score > best {
			best = score
		}
	}
	return best
}

func (r race) distanceAfter(seconds int) int {
	distances := r.runFor(seconds)
	best := 0
	for _, dist := range distances {
		if dist > best {
			best = dist
		}
	}
	return best
}

func (r race) runFor(seconds int) map[string]int {
	distances := map[string]int{}

	for name, stats := range r {
		iterTime := stats.flightTime + stats.restTime
		iters := seconds / iterTime
		distances[name] = iters * stats.speed * stats.flightTime
		remaining := seconds - iters*iterTime
		if remaining > stats.flightTime {
			remaining = stats.flightTime
		}
		distances[name] += remaining * stats.speed
	}

	return distances
}

type reindeerStats struct {
	flightTime int
	restTime   int
	speed      int
}

func loadInput(filename string) race {
	r := race{}

	for _, line := range utils.ReadInputLines(filename) {
		parts := strings.SplitN(line, " ", 2)
		name := parts[0]
		vals := utils.GetInts(parts[1])
		r[name] = reindeerStats{
			speed:      vals[0],
			flightTime: vals[1],
			restTime:   vals[2],
		}
	}
	return r
}
