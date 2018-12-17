package main

import (
	"fmt"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	lines := utils.ReadInputLines("input.txt")
	twoTimes := []string{}
	threeTimes := []string{}
	for _, line := range lines {
		chars := make(map[rune]int)
		for _, char := range line {
			chars[char]++
		}
		isTwoTimes := false
		isThreeTimes := false
		for _, count := range chars {
			if count == 2 {
				isTwoTimes = true
			} else if count == 3 {
				isThreeTimes = true
			}
		}
		if isTwoTimes {
			twoTimes = append(twoTimes, line)
		}
		if isThreeTimes {
			threeTimes = append(threeTimes, line)
		}
	}
	checksum := len(twoTimes) * len(threeTimes)
	fmt.Printf("Checksum: %d * %d = %d\n", len(twoTimes), len(threeTimes), checksum)
}
