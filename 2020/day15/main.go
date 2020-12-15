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
	spoken := spokenData{}

	for turn, num := range numbers {
		spoken[num] = []int{turn + 1}
	}
	prevNumber := numbers[len(numbers)-1]
	for turn := len(numbers) + 1; ; turn++ {
		var nextNum int
		if lastSpoken, found := spoken[prevNumber]; found {
			if len(lastSpoken) >= 2 {
				nextNum = lastSpoken[len(lastSpoken)-1] - lastSpoken[len(lastSpoken)-2]
			} else {
				nextNum = 0
			}
		} else {
			nextNum = 0
		}
		if turn == dinnerTime {
			return nextNum
		}
		spoken[nextNum] = append(spoken[nextNum], turn)
		prevNumber = nextNum
	}
}

type spokenData map[int][]int

var benchmark = false
