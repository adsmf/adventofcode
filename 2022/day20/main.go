package main

import (
	"container/ring"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	ints := getInts(input)
	p1 := run(ints, 1, 1)
	p2 := run(ints, 811589153, 10)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func run(ints intList, decryptionKey, iterations int) int {
	var zero *ring.Ring
	lookup := ringList{}
	for pos, val := range ints {
		*val *= decryptionKey
		lookup[pos].Value = val
		if *val == 0 {
			zero = &lookup[pos]
		}
		next := pos + 1
		if next == len(ints) {
			next = 0
		}
		lookup[pos].Link(&lookup[next])
	}

	for i := 0; i < iterations; i++ {
		for pos, val := range ints {
			prev := lookup[pos].Prev()
			moveBy := *val % (len(ints) - 1)
			toMove := prev.Unlink(1)
			moveTo := prev.Move(moveBy)
			moveTo.Link(toMove)
		}
	}

	val1k := zero.Move(1000).Value.(*int)
	val2k := zero.Move(2000).Value.(*int)
	val3k := zero.Move(3000).Value.(*int)
	return *val1k + *val2k + *val3k
}

type intList [5000]*int
type ringList [5000]ring.Ring

func getInts(input []byte) intList {
	ints := intList{}

	for pos, idx := 0, 0; pos < len(input); pos, idx = pos+1, idx+1 {
		var val int
		val, pos = getInt(input, pos)
		ints[idx] = &val
	}

	return ints
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	negative := false
	if in[pos] == '-' {
		negative = true
		pos++
	}
	for ; in[pos] >= '0' && in[pos] <= '9'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	if negative {
		accumulator = -accumulator
	}
	return accumulator, pos
}

var benchmark = false
