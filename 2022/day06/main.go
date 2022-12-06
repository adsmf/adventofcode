package main

import (
	_ "embed"
	"fmt"
	"math/bits"
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
	const len1, len2 = 4, 14
	p1 := findMarker(0, len1)
	p2 := findMarker(p1-len1, len2)
	return p1, p2
}

func findMarker(offset, signalLength int) int {
	for i := offset + signalLength; i < len(input); {
		last := uint32(0)
		for _, ch := range input[i-signalLength : i] {
			last |= 1 << (ch & 0x1f)
		}
		count := bits.OnesCount32(last)
		if count == signalLength {
			return i
		}
		i += signalLength - count
	}
	panic("not found")
}

var benchmark = false
