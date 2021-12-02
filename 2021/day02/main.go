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
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	pos := position{}
	for _, line := range utils.GetLines(input) {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "forward":
			pos.x += utils.MustInt(parts[1])
		case "down":
			pos.depth += utils.MustInt(parts[1])
		case "up":
			pos.depth -= utils.MustInt(parts[1])
		}
	}
	return pos.x * pos.depth
}

func part2() int {
	pos := position{}
	for _, line := range utils.GetLines(input) {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "forward":
			pos.x += utils.MustInt(parts[1])
			pos.depth += utils.MustInt(parts[1]) * pos.aim
		case "down":
			pos.aim += utils.MustInt(parts[1])
		case "up":
			pos.aim -= utils.MustInt(parts[1])
		}
	}
	return pos.x * pos.depth
}

type position struct {
	x, aim, depth int
}

var benchmark = false
