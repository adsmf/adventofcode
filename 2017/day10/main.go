package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	input := utils.GetInts(string(inputBytes))
	p1 := part1(input)
	p2 := part2([]byte(strings.TrimSpace(string(inputBytes))))
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func part1(input []int) int {
	loop := newRing(256)

	for skip, length := range input {
		loop.twist(length)
		loop.skip(skip)
	}
	return loop.entries[0] * loop.entries[1]
}

func part2(lengths []byte) string {
	loop := newRing(256)

	lengths = append(lengths, 17, 31, 73, 47, 23)
	skip := 0
	for iter := 0; iter < 64; iter++ {
		for _, length := range lengths {
			loop.twist(int(length))
			loop.skip(skip)
			skip++
		}
	}
	hash := ""
	for i := 0; i < 256; i += 16 {
		block := 0
		for j := 0; j < 16; j++ {
			block ^= loop.entries[i+j]
		}
		hash = fmt.Sprintf("%s%02x", hash, block)
	}
	return hash
}

type ring struct {
	entries []int
	offset  int
}

func (r *ring) twist(length int) {
	for i := 0; i < length/2; i++ {
		o1 := r.offset + i
		o2 := r.offset + length - i - 1
		o1 %= len(r.entries)
		o2 %= len(r.entries)
		r.entries[o1], r.entries[o2] = r.entries[o2], r.entries[o1]
	}

	r.skip(length)
}

func (r *ring) skip(n int) {
	r.offset = (r.offset + n) % len(r.entries)
}

func newRing(length int) ring {
	r := ring{
		entries: make([]int, length),
		offset:  0,
	}
	for i := 0; i < length; i++ {
		r.entries[i] = i
	}
	return r
}

var benchmark = false
