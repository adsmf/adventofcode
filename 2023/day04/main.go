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
	cardCopies := make(map[int]int, 200+numWinningNumbers)
	maxCard := 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		var cardNum int
		winMap := [100]int{}

		numMatches := 0
		utils.EachInteger(line, func(index, value int) (done bool) {
			switch {
			case index == 0:
				cardNum = value
			case index <= numWinningNumbers:
				winMap[value] = 1
			default:
				numMatches += winMap[value]
			}
			return false
		})

		if cardNum > maxCard {
			maxCard = cardNum
		}
		cardCopies[cardNum]++

		if numMatches == 0 {
			return false
		}

		totalPoints += (1 << numMatches) >> 1
		copies := cardCopies[cardNum]
		for i := cardNum + 1; i <= cardNum+numMatches; i++ {
			cardCopies[i] += copies
		}

		return false
	})

	totalCards := 0
	for i := 1; i <= maxCard; i++ {
		totalCards += cardCopies[i]
	}
	return totalPoints, totalCards
}

var benchmark = false
