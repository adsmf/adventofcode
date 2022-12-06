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
	p1, p2 := 0, 0
	const len1, len2 = 4, 14
	for i := len1; i < len(input); i++ {
		last := make(map[byte]int, len1)
		for _, ch := range input[i-len1 : i] {
			last[ch]++
		}
		if len(last) == len1 {
			p1 = i
			break
		}
	}
	for i := len2 - len1 + p1; i < len(input); i++ {
		last := make(map[byte]int, len2)
		for _, ch := range input[i-len2 : i] {
			last[ch]++
		}
		if len(last) == len2 {
			p2 = i
			break
		}
	}
	return p1, p2
}

var benchmark = false
