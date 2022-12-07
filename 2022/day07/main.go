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

func solve() (fsSize, fsSize) {
	nodePool := [180]fsSize{}
	poolAllocated := 0
	p1 := fsSize(0)
	p2 := fsSize(70000000)

	stack := [10]*fsSize{}
	stackPos := -1

	var curNode *fsSize

	completeDir := func() {
		stackPos--
		*(stack[stackPos]) += *curNode
		curNode = stack[stackPos]
	}

	for pos := 0; pos < len(input); {
		if input[pos] == '$' {
			if input[pos+2] == 'l' {
				pos += 5
				continue
			}
			if input[pos+5] == '.' {
				completeDir()
				pos += 8
				continue
			}
			stackPos++
			stack[stackPos] = &nodePool[poolAllocated]
			poolAllocated++
			curNode = stack[stackPos]
			pos = nextLine(input, pos)
			continue
		}
		size := 0
		size, pos = getInt(input, pos)
		*curNode += fsSize(size)
		pos = nextLine(input, pos)
	}
	for stackPos > 0 {
		completeDir()
	}
	spaceNeeded := fsSize(nodePool[0] + 30000000 - 70000000)
	for n := 0; n < poolAllocated; n++ {
		curNodeSize := nodePool[n]
		if curNodeSize <= 100000 {
			p1 += curNodeSize
		}
		if curNodeSize >= spaceNeeded && curNodeSize < p2 {
			p2 = curNodeSize
		}
	}

	return p1, p2
}

func nextLine(in []byte, pos int) int {
	for ; in[pos] != '\n'; pos++ {
	}
	return pos + 1
}

type fsSize uint32

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; pos < len(in) && in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

var benchmark = false
