package main

import (
	"fmt"
)

const (
	multA = uint64(16807)
	multB = uint64(48271)
)
const (
	seedA = uint64(703)
	seedB = uint64(516)
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	genA, genB := seedA, seedB
	mask := (uint64(1) << 16) - 1
	mod := uint64(2147483647)
	count := 0
	for i := 0; i < 40000000; i++ {
		genA *= multA
		genA %= mod
		genB *= multB
		genB %= mod
		if genA&mask == genB&mask {
			count++
		}
	}
	return count
}

func part2() int {
	genA, genB := seedA, seedB
	mask := (uint64(1) << 16) - 1
	mod := uint64(2147483647)
	count := 0
	for i := 0; i < 5000000; i++ {
		for a := 0; a == 0 || genA%4 != 0; a++ {
			genA *= multA
			genA %= mod
		}
		for b := 0; b == 0 || genB%8 != 0; b++ {
			genB *= multB
			genB %= mod
		}
		if genA&mask == genB&mask {
			count++
		}
	}
	return count
}

var benchmark = false
