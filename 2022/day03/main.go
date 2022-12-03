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
		compItems := uint(0)
		elfID := groupSize - 1
		for i := 0; i < compSize; i++ {
			itemID := uint(1 << itemPriority(input[i+lineStart]))
			compItems |= itemID
			groupItems[elfID] |= itemID
		}
		var commonItem byte
		for i := compSize; i < compSize*2; i++ {
			itemID := uint(1 << itemPriority(input[i+lineStart]))
			if commonItem == 0 {
				if compItems&(itemID) > 0 {
					commonItem = input[i+lineStart]
				}
			}
			groupItems[elfID] |= itemID
		}
		p1 += itemPriority(commonItem)
		if groupSize == 3 {
			common := groupItems[0] & groupItems[1] & groupItems[2]
			p2 += bits.Len(common) - 1
			groupItems[0], groupItems[1], groupItems[2] = 0, 0, 0
			groupSize = 0
		}

		lineStart = lineEnd + 2
	}
	return p1, p2
}

func itemPriority(item byte) int {
	if item >= 'a' {
		return int(item - 'a' + 1)
	}
	return int(item - 'A' + 27)
}

var benchmark = false
