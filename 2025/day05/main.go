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
	freshRanges := [][2]int{}
	freshCount := 0
	utils.EachSectionMB(input, "\n\n", func(secIdx int, section string) (done bool) {
		if secIdx == 0 {
			freshInts := utils.GetInts(section)
			for i := 0; i < len(freshInts); i += 2 {
				freshRanges = append(freshRanges, [2]int{freshInts[i], freshInts[i+1]})
			}
			return
		}
		utils.EachInteger(section, func(index, value int) (done bool) {
			for _, freshRange := range freshRanges {
				if value >= freshRange[0] && value <= freshRange[1] {
					freshCount++
					return
				}
			}
			return
		})
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
