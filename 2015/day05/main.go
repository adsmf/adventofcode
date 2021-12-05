package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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
	return countNice("input.txt")
}

func part2() int {
	return countNice2("input.txt")
}

func countNice(filename string) int {
	lines := utils.ReadInputLines(filename)
	nice := 0
	for _, line := range lines {
		if niceString(line) {
			nice++
		}
	}
	return nice
}

func niceString(input string) bool {
	vowels := 0
	var lastChar rune
	double := false
	for _, char := range input {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'b', 'd', 'q', 'y':
			if lastChar == char-rune(1) {
				return false
			}
		}
		if lastChar == char {
			double = true
		}
		lastChar = char
	}
	return double && vowels >= 3
}

func countNice2(filename string) int {
	lines := utils.ReadInputLines(filename)
	nice := 0
	for _, line := range lines {
		if niceString2(line) {
			nice++
		}
	}
	return nice
}

func niceString2(input string) bool {
	pairs := map[string]int{}
	// r2string := ""
	rule2 := false
	for i := 1; i < len(input); i++ {
		pair := input[i-1 : i+1]
		pairs[pair]++
		var p1, p2, p3 byte
		cur := input[i]
		p1 = input[i-1]
		if i > 1 {
			p2 = input[i-2]
			if i > 2 {
				p3 = input[i-3]
			}
		}
		if cur == p2 {
			rule2 = true
		}
		if cur == p1 && cur == p2 {
			if p3 != cur {
				pairs[pair]--
			}
		}
	}
	rule1 := false
	for _, count := range pairs {
		if count > 1 {
			rule1 = true
			break
		}
	}
	return rule1 && rule2
}

var benchmark = false
