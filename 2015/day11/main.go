package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func part1() string {
	start := utils.ReadInputLines("input.txt")[0]

	return nextValid(string(start))
}

func part2() string {
	start := utils.ReadInputLines("input.txt")[0]

	return nextValid(nextValid(string(start)))
}

func nextValid(start string) string {
	for {
		start = inc(start)
		if valid(start) {
			break
		}
	}
	return start
}
func inc(start string) string {
	startBytes := []byte(start)

	for i := len(startBytes) - 1; i >= 0; i-- {
		startBytes[i]++
		if startBytes[i] <= 'z' {
			break
		}
		startBytes[i] = 'a'
	}

	nextString := string(startBytes)

	return nextString
}

func valid(pass string) bool {
	hasStraight := false
	pairs := map[byte]bool{}
	for i := 0; i < len(pass); i++ {
		if pass[i] == 'i' ||
			pass[i] == 'l' ||
			pass[i] == 'o' {
			return false
		}
		if i >= 2 {
			if pass[i-2] == pass[i-1]-1 &&
				pass[i-1] == pass[i]-1 {
				hasStraight = true
			}
		}
		if i >= 1 {
			if pass[i-1] == pass[i] {
				pairs[pass[i]] = true
			}
		}
	}
	return hasStraight && len(pairs) >= 2
}

var benchmark = false
