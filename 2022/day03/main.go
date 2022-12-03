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
	p1, p2 := 0, 0
	groupItems := make([]uint, 3)
	groupSize := 0
	for lineStart := 0; lineStart < len(input); {
		lineEnd := lineStart
		for pos := lineStart; pos < len(input); pos++ {
			if input[pos] == '\n' {
				lineEnd = pos - 1
				break
			}
		}

		groupSize++
		compSize := (lineEnd - lineStart + 1) / 2
		comp1Items := uint(0)
		comp2Items := uint(0)
		elfID := groupSize - 1
		for i := 0; i < compSize; i++ {
			comp1Items |= itemID(input[i+lineStart])
		}
		for i := compSize; i < compSize*2; i++ {
			comp2Items |= itemID(input[i+lineStart])
		}
		commonItemID := comp1Items & comp2Items
		groupItems[elfID] = comp1Items | comp2Items
		p1 += idToPriority(commonItemID)
		if groupSize == 3 {
			common := groupItems[0] & groupItems[1] & groupItems[2]
			p2 += idToPriority(common)
			groupItems[0], groupItems[1], groupItems[2] = 0, 0, 0
			groupSize = 0
		}

		lineStart = lineEnd + 2
	}
	return p1, p2
}

func idToPriority(id uint) int {
	return bits.Len(id)
}

func itemID(item byte) uint {
	if item >= 'a' {
		return uint(1 << (item - 'a'))
	}
	return uint(1 << (item - 'A' + 26))
}

var benchmark = false
