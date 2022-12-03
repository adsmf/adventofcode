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
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	sum := 0
	for _, line := range utils.GetLines(input) {
		comp1, comp2 := line[0:len(line)/2], line[len(line)/2:]
		chars1 := map[rune]bool{}
		for _, ch := range comp1 {
			chars1[ch] = true
		}
		var item rune
		for _, ch := range comp2 {
			if chars1[ch] {
				item = ch
			}
		}
		sum += itemPriority(item)
	}
	return sum
}

func part2() int {
	sum := 0
	lines := utils.GetLines(input)
	for i := 0; i < len(lines); i += 3 {
		group := lines[i : i+3]
		items := map[rune]int{}
		for _, elf := range group {
			elfItems := map[rune]int{}
			for _, ch := range elf {
				elfItems[ch]++
			}
			for ch := range elfItems {
				items[ch]++
			}
		}
		for item, count := range items {
			if count == 3 {
				sum += itemPriority(item)
				break
			}
		}
	}
	return sum
}

func itemPriority(item rune) int {
	if item >= 'a' {
		return int(item - 'a' + 1)
	}
	return int(item - 'A' + 27)
}

var benchmark = false
