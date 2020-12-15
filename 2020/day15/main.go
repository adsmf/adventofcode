package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	p1 := play(string(input), 2020)
	p2 := play(string(input), 30000000)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func play(input string, dinnerTime int) int {
	numbers := utils.GetInts(input)
	spoken := map[int]int{}

	for turn, num := range numbers {
		spoken[num] = turn + 1
	}
	lastSpoken := numbers[len(numbers)-1]
	for turn := len(numbers) + 1; turn <= dinnerTime; turn++ {
		speak := 0
		if previous, found := spoken[lastSpoken]; found {
			speak = turn - previous - 1
		}
		spoken[lastSpoken] = turn - 1
		lastSpoken = speak
	}
	return lastSpoken
}

var benchmark = false
