package main

import (
	_ "embed"
	"fmt"
	"strings"

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
	top3 := utils.NewTopN[int](3)
	for _, elf := range strings.Split(input, "\n\n") {
		top3.Add(utils.SumInts(elf))
	}
	totalTop := 0
	vals := top3.Values()
	for _, elf := range vals {
		totalTop += elf
	}
	return (vals)[0], totalTop
}

var benchmark = false