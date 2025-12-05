package main

import (
	_ "embed"
	"fmt"
	"slices"

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
	freshRanges := make([][2]int, 0, 200)
	freshCount := 0
	sectionStart := 0
	for i := range len(input) - 1 {
		for ; input[i] != '\n'; i++ {
		}
		if input[i+1] == '\n' {
			sectionStart = i
			break
		}
	}
	start := 0
	utils.EachInteger(input[:sectionStart], func(index, value int) (done bool) {
		if index&1 == 0 {
			start = value
			return
		}
		freshRanges = append(freshRanges, [2]int{start, value})
		return
	})
	utils.EachInteger(input[sectionStart+2:], func(index, value int) (done bool) {
		for _, freshRange := range freshRanges {
			if value >= freshRange[0] && value <= freshRange[1] {
				freshCount++
				return
			}
		}
		return
	})
	slices.SortFunc(freshRanges, func(a, b [2]int) int {
		return a[0] - b[0]
	})
	n := 0
	for i := 0; i < len(freshRanges); i++ {
		if n > 0 && freshRanges[n-1][1] >= freshRanges[i][0] {
			freshRanges[n-1][1] = max(freshRanges[n-1][1], freshRanges[i][1])
			continue
		}
		freshRanges[n] = freshRanges[i]
		n++
	}
	freshRanges = freshRanges[:n]
	totalFresh := 0
	for i := range len(freshRanges) {
		totalFresh += freshRanges[i][1] - freshRanges[i][0] + 1
	}
	return freshCount, totalFresh
}

var benchmark = false
