package main

import (
	"container/ring"
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	intList := utils.GetInts(input)
	ints := make([]*int, 0, len(intList))
	for _, val := range intList {
		copy := val
		ints = append(ints, &copy)
	}
	p1 := part1(ints)
	p2 := part2(ints)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(ints []*int) int {
	r := ring.New(len(ints))
	var zero *ring.Ring
	lookup := map[*int]*ring.Ring{}
	for _, val := range ints {
		r.Value = val
		if *val == 0 {
			zero = r
		}
		lookup[val] = r
		r = r.Next()
	}

	for _, val := range ints {
		r := lookup[val]
		prev := r.Prev()
		moveBy := *val % (len(ints) - 1)
		toMove := prev.Unlink(1)
		moveTo := prev.Move(moveBy)
		moveTo.Link(toMove)
	}

	val1k := zero.Move(1000).Value.(*int)
	val2k := zero.Move(2000).Value.(*int)
	val3k := zero.Move(3000).Value.(*int)
	return *val1k + *val2k + *val3k
}

func part2(ints []*int) int {
	decryptionKey := 811589153

	r := ring.New(len(ints))
	var zero *ring.Ring
	lookup := map[*int]*ring.Ring{}
	for _, val := range ints {
		*val *= decryptionKey
		r.Value = val
		if *val == 0 {
			zero = r
		}
		lookup[val] = r
		r = r.Next()
	}

	for i := 0; i < 10; i++ {
		for _, val := range ints {
			r := lookup[val]
			prev := r.Prev()
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

var benchmark = false
