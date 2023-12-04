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
	cardCopies := map[int]int{}
	maxCard := 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		values := utils.GetInts(line)
		cardNum, winningNumbers, haveNumbers := values[0], values[1:numWinningNumbers+1], values[numWinningNumbers+1:]

		if cardNum > maxCard {
			maxCard = cardNum
		}
		cardCopies[cardNum]++

		winMap := make(map[int]bool, len(winningNumbers))
		for _, val := range winningNumbers {
			winMap[val] = true
		}

		numMatches := 0
		for _, have := range haveNumbers {
			if !winMap[have] {
				continue
			}
			numMatches++
		}
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
