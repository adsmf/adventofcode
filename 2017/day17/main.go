package main

import (
	"fmt"
)

const (
	input        = 356
	exampleInput = 3
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
	buf := buffer{1, zero}
	cur := buf
	for i := 1; i <= 2017; i++ {
		cur = cur.move(input)
		cur = cur.insertAndSelect(i)
	}
	return cur.next().value()
}

func part2() int {
	zero := &entry{0, nil}
	zero.next = zero
	buf := buffer{1, zero}
	cur := buf
	for i := 1; i <= 50000000; i++ {
		if i%500000 == 0 {
			fmt.Printf("%3d%%", i/500000)
		}
		cur = cur.move(input)
		cur = cur.insertAndSelect(i)
	}
	return zero.next.value
}

type buffer struct {
	size    int
	current *entry
}

func (b buffer) insertAndSelect(ordinal int) buffer {
	newEntry := &entry{
		value: ordinal,
		next:  b.current.next,
	}
	b.current.next = newEntry
	b.current = newEntry
	b.size++
	return b
}

func (b buffer) move(num int) buffer {
	return buffer{
		size:    b.size,
		current: b.current.move(num % b.size),
	}
}
func (b buffer) next() buffer {
	return buffer{
		size:    b.size,
		current: b.current.next,
	}
}

func (b buffer) value() int {
	return b.current.value
}

type entry struct {
	value int
	next  *entry
}

func (e *entry) move(num int) *entry {
	cur := e
	for i := 0; i < num; i++ {
		cur = cur.next
	}
	return cur
}

var benchmark = false
