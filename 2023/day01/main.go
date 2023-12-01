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
	utils.EachLine(input, func(line string) (done bool) {
		first, last := unset, unset
		for start, ch := range line {
			digit := unset
			if ch >= '0' && ch <= '9' {
				digit = int(ch - '0')
			} else if allowWords {
				for idx, word := range words {
					end := start + len(word)
					if len(line) < end {
						continue
					}
					if line[start:end] == word {
						digit = idx + 1
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
		return false
	})
	return totalCalibration
}

const unset = -1

var words = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var benchmark = false
