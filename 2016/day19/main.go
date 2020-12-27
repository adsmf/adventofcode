package main

import (
	"fmt"
	"math"
)

const input = 3014387

func main() {
	p1 := part1(input)
	p2 := part2(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(numElves int) int {
	first := &elfInfo{
		id: 1,
	}
	previous := first
	for i := 2; i <= numElves; i++ {
		elf := &elfInfo{
			id: i,
		}
		previous.next = elf
		previous = elf
	}
	previous.next = first
	cur := first
	for cur.next != cur {
		cur.next = cur.next.next
		cur = cur.next
	}
	return cur.id
}

func part2(numElves int) int {
	pow3ind := math.Floor(math.Log(float64(numElves)) / math.Log(3))
	pow3 := int(math.Pow(3, pow3ind))
	if pow3 == numElves {
		return numElves
	}
	rem := numElves - pow3
	if rem <= pow3 {
		return rem
	}
	return pow3 + 2*(rem-pow3)
}

type elfInfo struct {
	id   int
	next *elfInfo
}

var benchmark = false
