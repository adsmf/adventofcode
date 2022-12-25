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
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
	}
}

func part1() string {
	sum := 0
	for _, line := range utils.GetLines(input) {
		sum += deSnafu(line)
	}
	return snafu(sum)
}

func deSnafu(in string) int {
	val := 1
	total := 0
	for pos := len(in) - 1; pos >= 0; pos-- {
		ch := in[pos]
		switch ch {
		case '0', '1', '2':
			total += int(ch-'0') * val
		case '-':
			total += -val
		case '=':
			total += -2 * val
		}
		val *= 5
	}
	return total
}

func snafu(in int) string {
	val := 1
	for val < in {
		val *= 5
	}
	digits := []int{}
	for {
		val /= 5
		digit := in / val
		in -= digit * val
		digits = append(digits, int(digit))
		if val == 1 {
			break
		}
	}
	for pos := len(digits) - 1; pos >= 0; pos-- {
		if digits[pos] > 2 {
			digits[pos] -= 5
			digits[pos-1]++
		}
	}
	sb := strings.Builder{}
	for _, digit := range digits {
		sb.WriteByte(chars[digit])
	}
	return sb.String()
}

var chars = map[int]byte{-2: '=', -1: '-', 0: '0', 1: '1', 2: '2'}

var benchmark = false
