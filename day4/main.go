package main

import (
	"fmt"
	"strconv"
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
	validPasses := 0
	for test := passMin; test <= passMax; test++ {
		if validatePass(test, false) {
			validPasses = validPasses + 1
		}
	}
	return validPasses
}

func day2() int {
	validPasses := 0
	for test := passMin; test <= passMax; test++ {
		if validatePass(test, true) {
			validPasses = validPasses + 1
		}
	}
	return validPasses
}

func validatePass(testPass int, strict bool) bool {
	passString := strconv.Itoa(testPass)
	if len(passString) != 6 {
		return false
	}
	digits := []int{}
	for _, ch := range passString {
		val := int(ch - '0' + 1)
		digits = append(digits, val)
	}
	lastDigit := -1
	rangeLengths := []int{}
	for _, digit := range digits {
		if digit < lastDigit {
			return false
		}
		if digit == lastDigit {
			rangeLengths[len(rangeLengths)-1] = rangeLengths[len(rangeLengths)-1] + 1
		} else {
			rangeLengths = append(rangeLengths, 1)
		}
		lastDigit = digit
	}

	for _, length := range rangeLengths {
		if length == 2 || !strict && length > 2 {
			return true
		}
	}
	return false
}
