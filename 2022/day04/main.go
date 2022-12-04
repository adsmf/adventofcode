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
	contains, overlap := 0, 0
	sections := [4]int{}
	accumulator, index := 0, 0
	for _, ch := range input {
		switch {
		case ch >= '0' && ch <= '9':
			accumulator = accumulator * 10
			accumulator += int(ch - '0')
		default:
			sections[index] = accumulator
			accumulator = 0
			if index == 3 {
				if sections[0] <= sections[2] && sections[1] >= sections[3] ||
					sections[0] >= sections[2] && sections[1] <= sections[3] {
					contains++
				}
				if !(sections[1] < sections[2] || sections[0] > sections[3]) {
					overlap++
				}
				index = 0
				continue
			}
			index++
		}
	}
	return contains, overlap
}

var benchmark = false
