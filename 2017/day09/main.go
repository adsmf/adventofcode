package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	p1, p2 := process(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func process(input []byte) (int, int) {
	score := 0
	garbageCount := 0
	depth := 0
	for pos := 0; pos < len(input); pos++ {
		char := input[pos]
		switch char {
		case '{':
			depth++
			score += depth
		case '}':
			depth--
		case '<':
			offset, count := findCompleteGarbage(input[pos+1:])
			garbageCount += count
			pos += offset
		}
	}
	return score, garbageCount
}

func findCompleteGarbage(stream []byte) (int, int) {

	for offset, count := 0, 0; offset < len(stream); offset, count = offset+1, count+1 {
		switch stream[offset] {
		case '>':
			return offset, count
		case '!':
			count--
			offset++
		}
	}

	panic("Endless garbage")
}

var benchmark = false
