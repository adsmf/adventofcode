package main

import (
	"fmt"
	"regexp"
	"strconv"

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

func part1() int {
	pairs := loadInput("input.txt")
	return pairs.findBest()
}

func part2() int {
	pairs := loadInput("input.txt")
	pairs["me"] = map[string]int{}
	return pairs.findBest()
}

type pairHappinessMap map[string]map[string]int

func (p pairHappinessMap) findBest() int {
	people := []string{}
	for person := range p {
		people = append(people, person)
	}
	permutations := utils.PermuteStrings(people)
	best := 0
	for _, perm := range permutations {
		happiness := p.sumHappiness(perm)
		// fmt.Printf("%d: %v\n", happiness, perm)
		if happiness > best {
			best = happiness
		}
	}
	return best
}

func (p pairHappinessMap) sumHappiness(order []string) int {
	sum := 0
	for i := 0; i < len(order); i++ {
		p1 := order[i]
		p2 := ""
		if i == 0 {
			p2 = order[len(order)-1]
		} else {
			p2 = order[i-1]
		}
		pairVal1 := p[p1][p2]
		pairVal2 := p[p2][p1]
		sum += pairVal1 + pairVal2
	}
	return sum
}

func loadInput(filename string) pairHappinessMap {
	pairs := pairHappinessMap{}
	lines := utils.ReadInputLines(filename)

	re := regexp.MustCompile(`^(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)

		val, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		if matches[2] == "lose" {
			val *= -1
		}
		p1 := matches[1]
		p2 := matches[4]
		if pairs[p1] == nil {
			pairs[p1] = make(map[string]int)
		}
		pairs[p1][p2] = val
	}
	return pairs
}

var benchmark = false
