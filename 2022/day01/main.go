package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	topSnacks := make([]int, 3)
	elfCalories := 0
	accumulator := 0
	for _, ch := range input {
		switch {
		case ch >= '0' && ch <= '9':
			accumulator *= 10
			accumulator += int(ch - '0')
		default:
			if accumulator > 0 {
				elfCalories += accumulator
				accumulator = 0
				continue
			}
			if elfCalories < topSnacks[0] {
				elfCalories = 0
				continue
			}
			topSnacks[0] = elfCalories
			for pos := 0; pos < 2; pos++ {
				if topSnacks[pos] < topSnacks[pos+1] {
					break
				}
				topSnacks[pos+1], topSnacks[pos] = topSnacks[pos], topSnacks[pos+1]
			}
			elfCalories = 0
		}
	}
	return topSnacks[2], topSnacks[2] + topSnacks[1] + topSnacks[0]
}

var benchmark = false
