package main

import (
	"fmt"
)

var (
	passMin = 278384
	passMax = 824795
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	return countValid(false)
}

func part2() int {
	return countValid(true)
}

func countValid(strict bool) int {
	validPasses := 0
	for test := passMin; test <= passMax; test++ {
		if validatePass(test, strict) {
			validPasses = validPasses + 1
		}
	}
	return validPasses
}

func validatePass(testPass int, strict bool) bool {
	if testPass < 100000 || testPass > 999999 {
		return false
	}
	lastDigit := 10
	rangeLengths := []int{}
	for testPass > 0 {
		digit := testPass % 10
		if digit > lastDigit {
			return false
		}
		if digit == lastDigit {
			rangeLengths[len(rangeLengths)-1] = rangeLengths[len(rangeLengths)-1] + 1
		} else {
			rangeLengths = append(rangeLengths, 1)
		}
		lastDigit = digit
		testPass = int(testPass / 10)
	}
	for _, length := range rangeLengths {
		if length == 2 || !strict && length > 2 {
			return true
		}
	}
	return false
}

var benchmark = false
