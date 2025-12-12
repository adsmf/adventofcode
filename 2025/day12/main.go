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
		areaAvail := 0
		width := 0
		maxSpace := 0
		utils.EachInteger(line, func(intIdx, value int) (done bool) {
			switch intIdx {
			case 0:
				width = value
			case 1:
				areaAvail = width * value
			default:
				maxSpace += 9 * value
			}
			return
		})
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
