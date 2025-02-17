package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func evolve(secret int) int {
	secret ^= (secret << 6) & ((1 << 24) - 1)
	secret ^= (secret >> 5) & ((1 << 24) - 1)
	secret ^= (secret * 1 << 11) & ((1 << 24) - 1)
	return secret
}

func solve() (int, int) {
	p1, p2 := 0, 0
	counts := [1 << 21]uint16{}
	utils.EachInteger(input, func(_, value int) (done bool) {
		seen := [1 << 20]bool{}
		lastPrice := uint16(value % 10)
		hash := uint32(0)
		for i := 0; i < 2000; i++ {
			value = evolve(value)
			price := uint16(value % 10)
			diff := price - lastPrice
			lastPrice = price
			hash = (hash<<5)&((1<<19)-1) | uint32(diff+10)
			if seen[hash] {
				continue
			}
			seen[hash] = true
			counts[hash] += price
		}
		p1 += value
		return
	})
	for _, count := range counts {
		p2 = max(int(count), p2)
	}
	return p1, p2
}

var benchmark = false
