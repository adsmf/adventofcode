package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var benchmark = false

func main() {
	part1, part2 := loadAndCalculate("input.txt")
	if !benchmark {
		fmt.Printf("Part 1: %d\n", part1)
		fmt.Printf("Part 2: %d\n", part2)
	}
}

func loadAndCalculate(filename string) (int, int) {
	inputBytes, _ := ioutil.ReadFile(filename)
	blocks := strings.Split(string(inputBytes), "\n\n")
	any, all := 0, 0
	for _, block := range blocks {
		groups := groupAnswers{}
		lines := strings.Split(strings.TrimSpace(block), "\n")
		for _, line := range lines {
			for _, char := range line {
				if char >= 'a' && char <= 'z' {
					groups[string(char)]++
				}
			}
		}
		any += len(groups)

		for _, count := range groups {
			if count == len(lines) {
				all++
			}

		}
	}
	return any, all
}

type groupAnswers map[string]int
