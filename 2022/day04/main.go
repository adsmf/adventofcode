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
	sec := [4]int{}
	accumulator, index := 0, 0
	for _, ch := range input {
		if ch&0xf0 == 0x30 {
			accumulator <<= 4
			accumulator |= int(ch & 0xf)
			continue
		}
		sec[index] = accumulator
		accumulator = 0
		index++
		if index < 4 {
			continue
		}
		index = 0
		if !(sec[1] < sec[2] || sec[0] > sec[3]) {
			overlap++
			if sec[0] <= sec[2] && sec[1] >= sec[3] ||
				sec[0] >= sec[2] && sec[1] <= sec[3] {
				contains++
			}
		}
	}
	return contains, overlap
}

var benchmark = false
