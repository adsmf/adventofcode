package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	contains := 0
	overlap := 0
	for _, line := range utils.GetLines(input) {
		sections := utils.GetInts(line)
		if sections[0] <= sections[2] && sections[1] >= sections[3] ||
			sections[0] >= sections[2] && sections[1] <= sections[3] {
			contains++
		}
		if sections[1] < sections[2] || sections[0] > sections[3] {
			continue
		}
		overlap++
	}
	return contains, overlap
}

var benchmark = false
