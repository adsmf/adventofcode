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
	groupItems := idMax
	groupSize := 0
	for lineStart, lineEnd := 0, 0; lineStart < len(input); {
		for pos := lineStart; input[pos] != '\n'; pos++ {
			lineEnd = pos
		}
		groupSize++
		compSize := (lineEnd - lineStart + 1) / 2
		comp1Items := uint(0)
		comp2Items := uint(0)
		for i := 0; i < compSize; i++ {
			comp1Items |= itemID(input[i+lineStart])
		}
		for i := compSize; i < compSize*2; i++ {
			comp2Items |= itemID(input[i+lineStart])
		}
		commonItemID := comp1Items & comp2Items
		p1 += idToPriority(commonItemID)
		groupItems &= comp1Items | comp2Items
		if groupSize == 3 {
			p2 += idToPriority(groupItems)
			groupItems = idMax
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
	item -= 'A' - 26
	if item >= 52 {
		item -= 32 + 26
	}
	return uint(1 << item)
}

const idMax = uint(1<<52 - 1)

var benchmark = false
