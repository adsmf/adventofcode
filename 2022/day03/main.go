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
	line := make([]byte, 0, 60)
	groupItems := map[byte]byte{}
	groupSize := 0
	compItems := make(map[byte]nothing, 52)
	for _, ch := range input {
		switch ch {
		case '\n':
			groupSize++
			compSize := len(line) / 2
			elfID := byte(1 << (groupSize - 1))
			for i := 0; i < compSize; i++ {
				compItems[line[i]] = nothing{}
				groupItems[line[i]] |= elfID
			}
			var commonItem byte
			for i := compSize; i < compSize*2; i++ {
				if commonItem == 0 {
					if _, found := compItems[line[i]]; found {
						commonItem = line[i]
					}
				}
				groupItems[line[i]] |= elfID
			}
			p1 += itemPriority(commonItem)
			if groupSize == 3 {
				for item, elfBits := range groupItems {
					if elfBits == 7 {
						p2 += itemPriority(item)
					}
					delete(groupItems, item)
				}
				groupSize = 0
			}
			for item := range compItems {
				delete(compItems, item)
			}
			line = line[0:0]
		default:
			line = append(line, ch)
		}
	}
	return p1, p2
}

func itemPriority(item byte) int {
	if item >= 'a' {
		return int(item - 'a' + 1)
	}
	return int(item - 'A' + 27)
}

type nothing struct{}

var benchmark = false
