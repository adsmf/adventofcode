package main

import (
	"fmt"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	part1()
	part2()
}

func part1() {
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
	fmt.Printf("Part 1 - Checksum: %d * %d = %d\n", len(twoTimes), len(threeTimes), checksum)
}

func part2() {
	lines := utils.ReadInputLines("input.txt")
	for i, a := range lines {
		for j, b := range lines {
			if j <= i {
				continue
			}
			dist, common := distance(a, b)
			if dist == 1 {
				fmt.Printf("Distance %s %s = %d (common: %s)\n", a, b, dist, common)
			}
		}
	}
}

func distance(a, b string) (int, string) {
	differences := []int{}
	common := ""
	for col, char := range a {
		if a[col] != b[col] {
			differences = append(differences, col)
			continue
		}
		common = fmt.Sprintf("%s%c", common, char)
	}

	return len(differences), common
}
