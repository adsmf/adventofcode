package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := calibrate(false)
	p2 := calibrate(true)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func calibrate(allowWords bool) int {
	totalCalibration := 0
	for _, line := range utils.GetLines(input) {
		first, last := unset, unset
		for start, ch := range line {

			digit := unset
			if ch >= '0' && ch <= '9' {
				digit = int(ch - '0')
			} else if allowWords {
				for word, value := range wordValues {
					end := start + len(word)
					if len(line) < end {
						continue
					}
					if line[start:end] == word {
						digit = value
					}
				}
			}

			if digit == unset {
				continue
			}

			if first == unset {
				first = digit
			}
			last = digit
		}
		totalCalibration += first*10 + last
	}
	return totalCalibration
}

const unset = -1

var wordValues = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var benchmark = false
