package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
	}
}

var chars = map[int]byte{-2: '=', -1: '-', 0: '0', 1: '1', 2: '2'}
var vals = [...]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}

func part1() snafuDigits {
	sum := snafuDigits{}
	accumulator := snafuDigits{}
	accPos := 0
	for pos := 0; pos < len(input); pos++ {
		ch := input[pos]
		if ch == '\n' {
			accPos--
			len := accPos
			for accPos >= 0 {
				sum[len-accPos] += accumulator[accPos]
				accPos--
			}

			accPos = 0
			accumulator = snafuDigits{}
			continue
		}
		accumulator[accPos] = vals[ch]
		accPos++
	}
	for pos := 0; pos < maxDigits-1; pos++ {
		for sum[pos] > 2 {
			sum[pos] -= 5
			sum[pos+1]++
		}
		for sum[pos] < -2 {
			sum[pos] += 5
			sum[pos+1]--
		}
	}
	return sum
}

const maxDigits = 25

type snafuDigits [maxDigits]int

func (s snafuDigits) String() string {
	sb := strings.Builder{}
	leadingZero := true
	for pos := maxDigits - 1; pos >= 0; pos-- {
		digit := s[pos]
		if leadingZero && digit == 0 {
			continue
		}
		leadingZero = false
		sb.WriteByte(chars[digit])
	}
	return sb.String()
}

var benchmark = false
