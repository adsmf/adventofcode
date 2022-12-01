package utils

import (
	"strings"
)

func GetInts(input string) []int {
	ints := make([]int, 0, len(input)/2)

	accumulator := 0
	negative := false
	started := false

	for _, char := range append([]byte(input), '\n') {
		switch {
		case !started && char == '-':
			negative = true
		case char >= '0' && char <= '9':
			accumulator = accumulator*10 + int(char-'0')
			started = true
		default:
			if started {
				if negative {
					accumulator *= -1
				}
				ints = append(ints, accumulator)
				accumulator = 0
				started = false
			}
			negative = false
		}
	}
	return ints
}

func GetLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func SumInts(input string) int {
	accumulator := 0
	negative := false
	started := false
	sum := 0

	for _, char := range append([]byte(input), '\n') {
		switch {
		case !started && char == '-':
			negative = true
		case char >= '0' && char <= '9':
			accumulator = accumulator*10 + int(char-'0')
			started = true
		default:
			if started {
				if negative {
					accumulator *= -1
				}
				sum += accumulator
				accumulator = 0
				started = false
			}
			negative = false
		}
	}
	return sum
}
