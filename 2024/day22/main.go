package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func evolve(secret int, steps int) int {
	for i := 0; i < steps; i++ {
		secret ^= (secret << 6) & ((1 << 24) - 1)
		secret ^= (secret >> 5) & ((1 << 24) - 1)
		secret ^= (secret * 1 << 11) & ((1 << 24) - 1)
	}
	return secret
}

func part1() int {
	p1 := 0
	utils.EachInteger(input, func(_, value int) (done bool) {
		evolved := evolve(value, 2000)
		p1 += evolved
		return
	})
	return p1
}

func part2() int {
	window := [4]int{}
	counts := map[int]int{}
	utils.EachInteger(input, func(_, value int) (done bool) {
		seen := map[int]bool{}
		lastPrice := value % 10
		window[3] = lastPrice
		for i := 1; i < 2000; i++ {
			value = evolve(value, 1)
			price := value % 10
			diff := price - lastPrice
			lastPrice = price

			window[0], window[1], window[2], window[3] = window[1], window[2], window[3], diff
			hash := (window[0] + 20) |
				(window[1]+20)<<8 |
				(window[2]+20)<<16 |
				(window[3]+20)<<24
			if !seen[hash] {
				counts[hash] += price
			}
			seen[hash] = true
		}
		return
	})
	p2 := 0
	for _, count := range counts {
		p2 = max(count, p2)
	}
	return p2
}

var benchmark = false
