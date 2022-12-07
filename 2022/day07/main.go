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
	p1 := fsSize(0)
	p2 := fsSize(70000000)
	spaceNeeded := fsSize(sumAllInts(input) + 30000000 - 70000000)

	stack := make(nodeList, 10)
	stackPos := -1

	var curNode *node

	completeDir := func() {
		stackPos--
		curNodeSize := curNode.size()
		stack[stackPos].subdirSize += curNodeSize
		if curNodeSize <= 100000 {
			p1 += curNodeSize
		}
		if curNodeSize >= spaceNeeded && curNodeSize < p2 {
			p2 = curNodeSize
		}

		curNode = &stack[stackPos]
	}

	for pos := 0; pos < len(input); {
		switch input[pos] {
		case '$':
			if input[pos+2] != 'c' {
				pos += 5
				continue
			}
			if input[pos+5] == '.' {
				completeDir()
				pos += 8
				continue
			}
			stackPos++
			stack[stackPos] = node{}
			curNode = &stack[stackPos]
			pos = nextLine(input, pos)
		default:
			size := 0
			size, pos = getInt(input, pos)
			curNode.addFile(fsSize(size))
			pos = nextLine(input, pos)
		}
	}
	for stackPos > 0 {
		completeDir()
	}

	return p1, p2
}

func nextLine(in []byte, pos int) int {
	for ; in[pos] != '\n'; pos++ {
	}
	return pos + 1
}

type nodeList []node

type fsSize uint32

type node struct {
	totalFileSize fsSize
	subdirSize    fsSize
}

func (n *node) addFile(size fsSize) { n.totalFileSize += size }
func (n *node) size() fsSize        { return n.totalFileSize + n.subdirSize }

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	for ; pos < len(in) && in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return accumulator, pos
}

func sumAllInts(in []byte) int {
	sum := 0

	accumulator := 0
	for pos := 0; pos < len(in); pos++ {
		if in[pos]&0xf0 != 0x30 {
			sum += accumulator
			accumulator = 0
			continue
		}
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	return sum
}

var benchmark = false
