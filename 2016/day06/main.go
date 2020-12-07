package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	input := load("input.txt")
	p1, p2 := calculate(input)
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func calculate(data []frequencyData) (string, string) {
	codeMin := make([]rune, len(data))
	codeMax := make([]rune, len(data))
	for i := 0; i < len(data); i++ {
		min := 9999999
		max := 0
		maxChar := '?'
		minChar := '?'
		for char, count := range data[i] {
			if count > max {
				max = count
				maxChar = char
			}
			if count < min {
				min = count
				minChar = char
			}
		}
		codeMax[i] = maxChar
		codeMin[i] = minChar
	}
	return string(codeMax), string(codeMin)
}

func load(filename string) []frequencyData {
	frequencies := []frequencyData{}
	for _, line := range utils.ReadInputLines(filename) {
		for i := 0; i < len(line); i++ {
			if len(frequencies) < i+1 {
				frequencies = append(frequencies, frequencyData{})
			}
			char := line[i]
			frequencies[i][rune(char)]++
		}
	}
	return frequencies
}

type frequencyData map[rune]int
