package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var benchmark = false

func main() {
	if !benchmark {
		fmt.Printf("Part 1: %d\n", part1())
		fmt.Printf("Part 2: %d\n", part2())
	}
}

func part1() int {
	groups := load("input.txt")
	count := 0
	for _, group := range groups {
		count += len(group)
	}
	return count
}

func part2() int {
	groups := load2("input.txt")
	count := 0
	for _, group := range groups {
		count += len(group)
	}
	return count
}

func load(filename string) []groupAnswers {
	inputBytes, _ := ioutil.ReadFile(filename)
	blocks := strings.Split(string(inputBytes), "\n\n")
	groups := []groupAnswers{}
	for _, block := range blocks {
		group := groupAnswers{}
		for _, char := range block {
			switch {
			case char >= 'a' && char <= 'z':
				group[string(char)] = true
			default:
			}
		}
		groups = append(groups, group)
	}
	return groups
}

func load2(filename string) []groupAnswers {
	inputBytes, _ := ioutil.ReadFile(filename)
	blocks := strings.Split(string(inputBytes), "\n\n")
	groups := []groupAnswers{}
	for _, block := range blocks {
		group := groupAnswers{}
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
				default:
				}
			}
			people = append(people, person)
		}
		for answer := range people[0] {
			allAnswered := true
			for _, person := range people {
				if _, found := person[answer]; !found {
					allAnswered = false
					break
				}
			}
			if allAnswered {
				group[answer] = true
			}
		}
		groups = append(groups, group)
	}
	return groups
}

type groupAnswers map[string]bool
