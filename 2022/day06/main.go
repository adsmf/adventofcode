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
	p1 := findMarker(len1, 0, signal(input))
	p2 := findMarker(len2, p1-len1, signal(input))
	return p1, p2
}

func findMarker(signalLength, offset int, stream signal) int {
	for i := offset + signalLength; i < len(stream); {
		last := uint32(0)
		for _, ch := range stream[i-signalLength : i] {
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

type signal *[4096]byte

var benchmark = false
