package main

import (
	"fmt"
)

const (
	input = 356
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
	zero := &entry{0, nil}
	zero.next = zero
	cur := zero
	for i := 1; i <= 2017; i++ {
		moves := input % i
		for m := 0; m < moves; m++ {
			cur = cur.next
		}
		new := &entry{
			value: i,
			next:  cur.next,
		}
		cur.next = new
		cur = new
	}
	return cur.next.value
}

func part2() int {
	after0 := 1
	pos := 0
	for i := 1; i <= 50000000; i++ {
		pos += input
		pos %= i
		if pos == 0 {
			after0 = i
		}
		pos++
	}
	return after0
}

type entry struct {
	value int
	next  *entry
}

var benchmark = false
