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

type StringIterator func(index int, section string) (done bool)

func EachLine(input string, callback StringIterator) {
	EachSection(input, '\n', callback)
}

func EachSection(input string, separator byte, callback StringIterator) {
	index := 0
	pos := 0
	start := 0
	for ; pos < len(input); pos++ {
		if input[pos] == separator {
			done := callback(index, input[start:pos])
			if done {
				return
			}
			start = pos + 1
			index++
		}
	}
	if start != len(input) {
		callback(index, input[start:pos])
	}
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
