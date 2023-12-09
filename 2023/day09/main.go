package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1, p2 := 0, 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		initial := utils.GetInts(line)

		orderDiffs := append([]int{}, initial...)
		diffStack := [][]int{initial}

		for !allZero(orderDiffs) {
			lastOrder := diffStack[len(diffStack)-1]
			orderDiffs = []int{}
			for i := 0; i < len(lastOrder)-1; i++ {
				diff := lastOrder[i+1] - lastOrder[i]
				orderDiffs = append(orderDiffs, diff)
			}
			diffStack = append(diffStack, orderDiffs)
		}
		next := 0
		lastDiff := 0
		for i := len(diffStack) - 1; i >= 0; i-- {
			next += diffStack[i][len(diffStack[i])-1]
			if i == 0 {
				break
			}
			lastDiff = diffStack[i-1][0] - lastDiff
		}

		p1 += next
		p2 += lastDiff
		return false
	})
	return p1, p2
}

func allZero(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}
	return true
}

var benchmark = false
