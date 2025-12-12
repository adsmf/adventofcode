package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	packable := 0
	regions := input[regionsStart():]

	utils.EachLine(regions, func(lineIdx int, line string) (done bool) {
		vals := utils.GetInts(line)
		width, height := vals[0], vals[1]
		vals = vals[2:]
		areaAvail := width * height
		maxSpace := 0
		for _, count := range vals {
			maxSpace += 9 * count
		}
		if maxSpace > areaAvail {
			return
		}
		packable++
		return
	})
	return packable
}

func regionsStart() int {
	lastNewline := 0
	for idx, ch := range input {
		switch ch {
		case '\n':
			lastNewline = idx
		case 'x':
			return lastNewline + 1
		}
	}
	return -1
}

var benchmark = false
