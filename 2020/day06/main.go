package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var benchmark = false

func main() {
	part1, part2 := loadAndCalculate("input.txt")
	if !benchmark {
		fmt.Printf("Part 1: %d\n", part1)
		fmt.Printf("Part 2: %d\n", part2)
	}
}

func loadAndCalculate(filename string) (int, int) {
	inputBytes, _ := ioutil.ReadFile(filename)
	blocks := strings.Split(string(inputBytes), "\n\n")
	any, all := 0, 0
	for _, block := range blocks {
		groupAny := groupAnswers{}
		people := []groupAnswers{}
		for _, line := range strings.Split(block, "\n") {
			if line == "" {
				continue
			}
			person := groupAnswers{}
			for _, char := range line {
				switch {
				case char >= 'a' && char <= 'z':
					person[string(char)] = true
					groupAny[string(char)] = true
				default:
				}
			}
			people = append(people, person)
		}
		any += len(groupAny)

		groupAll := groupAnswers{}
		for answer := range people[0] {
			allAnswered := true
			for _, person := range people {
				if _, found := person[answer]; !found {
					allAnswered = false
					break
				}
			}
			if allAnswered {
				groupAll[answer] = true
			}
		}
		all += len(groupAll)
	}
	return any, all
}

type groupAnswers map[string]bool
