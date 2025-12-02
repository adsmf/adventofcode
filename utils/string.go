package utils

import (
	"strings"
)

func GetInts[I string | []byte](input I) []int {
	ints := make([]int, 0, len(input)/2)

	EachInteger(input, func(index, value int) (done bool) {
		ints = append(ints, value)
		return false
	})
	return ints
}

func GetLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

type Callback[T any] func(index int, value T) (done bool)

func EachLine(input string, callback Callback[string]) {
	EachSection(input, '\n', callback)
}

func EachSection(input string, separator byte, callback Callback[string]) {
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

func EachSectionMB(input string, separator string, callback Callback[string]) {
	index := 0
	pos := 0
	start := 0
	for ; pos < len(input); pos++ {
		if strings.HasPrefix(input[pos:], separator) {
			done := callback(index, input[start:pos])
			if done {
				return
			}
			start = pos + len(separator)
			index++
		}
	}
	if start != len(input) {
		callback(index, input[start:pos])
	}
}

func EachInteger[I string | []byte](input I, callback Callback[int]) {
	accumulator := 0
	negative := false
	started := false
	index := 0
	send := func() (done bool) {
		if started {
			if negative {
				accumulator *= -1
			}
			done = callback(index, accumulator)
			index++
			accumulator = 0
			started = false
		}
		negative = false
		return
	}
	for _, char := range []byte(input) {
		switch {
		case !started && char == '-':
			negative = true
		case char >= '0' && char <= '9':
			accumulator = accumulator*10 + int(char-'0')
			started = true
		default:
			if send() {
				return
			}
		}
	}
	send()
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
