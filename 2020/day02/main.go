package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	samples := load("input.txt")
	validCount := 0
	for _, sample := range samples {
		if validP1(sample) {
			validCount++
		}
	}
	return validCount
}

func part2() int {
	samples := load("input.txt")
	validCount := 0
	for _, sample := range samples {
		if validP2(sample) {
			validCount++
		}
	}
	return validCount
}

func validP1(sample policySample) bool {
	count := 0
	for _, char := range sample.pass {
		if char == sample.char {
			count++
			if count > sample.upper {
				return false
			}
		}
	}
	return count >= sample.lower
}

func validP2(sample policySample) bool {
	char1 := sample.pass[sample.lower-1]
	char2 := sample.pass[sample.upper-1]
	p1v := rune(char1) == sample.char
	p2v := rune(char2) == sample.char

	return p1v != p2v
}

func load(filename string) []policySample {
	samples := []policySample{}
	lines := utils.ReadInputLines(filename)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		rangeParts := utils.GetInts(parts[0])
		lineSample := policySample{
			lower: rangeParts[0],
			upper: rangeParts[1],
			char:  rune(parts[1][0]),
			pass:  parts[2],
		}
		samples = append(samples, lineSample)
	}
	return samples
}

type policySample struct {
	lower, upper int
	char         rune
	pass         string
}
