package main

import (
	"fmt"
)

var (
	passMin = 278384
	passMax = 824795
)

func main() {
	fmt.Printf("Day 1: %d\n", day1())
	fmt.Printf("Day 2: %d\n", day2())
}

func day1() int {
	return countValid(false)
}

func day2() int {
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
