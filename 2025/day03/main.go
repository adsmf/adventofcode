package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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

func solve() (int, int) {
	p1, p2 := 0, 0
	utils.EachLine(input, func(index int, line []byte) (done bool) {
		best2 := make(bestJolt, 2)
		best12 := make(bestJolt, 12)
		for i, j := range line {
			jolt := j - '0'
			best2.update(len(line)-i, int(jolt))
			best12.update(len(line)-i, int(jolt))
		}
		p1 += best2.value()
		p2 += best12.value()
		return
	})
	return p1, p2
}

type bestJolt []int

func (b bestJolt) value() int {
	jolts := 0
	for i := range len(b) {
		jolts = jolts*10 + b[i]
	}
	return jolts
}

func (b bestJolt) update(rem, jolt int) {
	for bat := range len(b) {
		if jolt > b[bat] && (len(b)-bat) < (rem+1) {
			b[bat] = jolt
			copy(b[bat+1:], empty[:])
			return
		}
	}
}

var empty = [12]int{}

var benchmark = false
