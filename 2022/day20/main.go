package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	ints, numInts := getInts(input)
	p1 := run(ints, numInts, 1, 1)
	p2 := run(ints, numInts, 811589153, 10)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type entry struct {
	value int
	prev  int
	next  int
}

func run(ints intList, numInts, decryptionKey, iterations int) int {
	var zeroPos int
	lookup := ringList{}
	for pos := range ints {
		ints[pos] *= decryptionKey
		lookup[pos].value = ints[pos]
		if ints[pos] == 0 {
			zeroPos = pos
		}
		next := pos + 1
		if next == numInts {
			next = 0
		}
		lookup[pos].next = next
		lookup[next].prev = pos
	}

	move := func(pos, by int) int {
		if by == 0 {
			return pos
		}
		if by < 0 {
			for i := 0; i > by; i-- {
				pos = lookup[pos].prev
			}
			return pos
		}
		for i := 0; i < by; i++ {
			pos = lookup[pos].next
		}
		return pos
	}

	for i := 0; i < iterations; i++ {
		for pos, val := range ints {
			moveBy := val % (numInts - 1)
			lookup[lookup[pos].prev].next, lookup[lookup[pos].next].prev = lookup[pos].next, lookup[pos].prev

			newPrev := move(lookup[pos].prev, moveBy)
			newNext := lookup[newPrev].next
			lookup[newPrev].next = pos
			lookup[newNext].prev = pos
			lookup[pos].prev = newPrev
			lookup[pos].next = newNext
		}
	}

	sum := 0
	sum += lookup[move(zeroPos, 1000)].value
	sum += lookup[move(zeroPos, 2000)].value
	sum += lookup[move(zeroPos, 3000)].value
	return sum
}

type intList [5000]int
type ringList [5000]entry

func getInts(input []byte) (intList, int) {
	ints := intList{}
	idx := 0
	for pos := 0; pos < len(input); pos, idx = pos+1, idx+1 {
		var val int
		val, pos = getInt(input, pos)
		ints[idx] = val
	}

	return ints, idx
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
