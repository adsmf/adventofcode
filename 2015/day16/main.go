package main

import (
	"fmt"
	"regexp"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

var scanData = sueInfo{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func part1() int {
	s := loadInput("input.txt")
	return s.find(scanData)
}

func part2() int {
	s := loadInput("input.txt")
	return s.findFuzzy(scanData)
}

type sues map[int]sueInfo

func (s sues) find(targetInfo sueInfo) int {
	for num, info := range s {
		match := true
		for req, reqNum := range targetInfo {
			if has, found := info[req]; found {
				if has != reqNum {
					match = false
					break
				}
			}
		}
		if match {
			return num
		}
	}
	return -1
}

func (s sues) findFuzzy(targetInfo sueInfo) int {
	for num, info := range s {
		match := true
		for req, reqNum := range targetInfo {
			if has, found := info[req]; found {
				switch req {
				case "cats", "trees":
					if has <= reqNum {
						match = false
					}
				case "goldfish", "pomeranians":
					if has >= reqNum {
						match = false
					}
				default:
					if has != reqNum {
						match = false
					}
				}
				if !match {
					break
				}
			}
		}
		if match {
			return num
		}
	}
	return -1
}

type sueInfo map[string]int

func loadInput(filename string) sues {
	s := sues{}

	re := regexp.MustCompile(`^Sue \d+: (\w+): \d+, (\w+): \d+, (\w+): \d+$`)
	for _, line := range utils.ReadInputLines(filename) {
		matches := re.FindStringSubmatch(line)
		ints := utils.GetInts(line)
		newSue := sueInfo{
			matches[1]: ints[1],
			matches[2]: ints[2],
			matches[3]: ints[3],
		}
		s[ints[0]] = newSue
	}

	return s
}

var benchmark = false
