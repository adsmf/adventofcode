package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

const numWinningNumbers = 10

func main() {
	p1, p2 := analyse()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func analyse() (int, int) {
	totalPoints := 0
	totalCards := 0
	extraCopies := [numWinningNumbers + 1]int{}

	shiftExtraCopies := func() int {
		first := extraCopies[0]
		for i := 0; i < numWinningNumbers; i++ {
			extraCopies[i] = extraCopies[i+1]
		}
		extraCopies[numWinningNumbers] = 0
		return first
	}

	utils.EachLine(input, func(index int, line string) (done bool) {
		winMap := [100]int{}

		numMatches := 0
		utils.EachInteger(line, func(index, value int) (done bool) {
			switch {
			case index == 0:
				// Ignore card number
			case index <= numWinningNumbers:
				winMap[value] = 1
			default:
				numMatches += winMap[value]
			}
			return false
		})

		copies := 1 + shiftExtraCopies()
		totalCards += copies

		if numMatches == 0 {
			return false
		}

		totalPoints += (1 << numMatches) >> 1
		for i := 0; i < numMatches; i++ {
			extraCopies[i] += copies
		}

		return false
	})

	return totalPoints, totalCards
}

var benchmark = false
