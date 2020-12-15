package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	numbers := utils.GetInts(string(input))
	p1 := play(numbers, 2020)
	p2 := play(numbers, 30000000)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func play(numbers []int, dinnerTime int) int {
	spoken := make([]int, dinnerTime+len(numbers))

	for turn, num := range numbers {
		spoken[num] = turn + 1
	}
	lastSpoken := numbers[len(numbers)-1]
	for turn := len(numbers) + 1; turn <= dinnerTime; turn++ {
		speak := 0
		if previous := spoken[lastSpoken]; previous != 0 {
			speak = turn - previous - 1
		}
		spoken[lastSpoken] = turn - 1
		lastSpoken = speak
	}
	return lastSpoken
}

var benchmark = false
