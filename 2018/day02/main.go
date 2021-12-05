package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func part1() int {
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
	return checksum
}

func part2() string {
	lines := utils.ReadInputLines("input.txt")
	for i, a := range lines {
		for j, b := range lines {
			if j <= i {
				continue
			}
			dist, common := distance(a, b)
			if dist == 1 {
				return common
			}
		}
	}
	return "????"
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

var benchmark = false
