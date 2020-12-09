package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	p1 := part1()
	p2 := part2(p1)
	p2alt := part2alt(p1)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
		fmt.Printf("Part 2 alt: %d\n", p2alt)
	}
}

func part1() int {
	n := 25
	input, _ := ioutil.ReadFile("input.txt")
	lastN := map[int]int{}
	numbers := utils.GetInts(string(input))
	for index, number := range numbers {
		if index < n {
			lastN[number]++
			continue
		}
		valid := false
		for opt := range lastN {
			if _, found := lastN[number-opt]; found {
				valid = true
				break
			}
		}
		if !valid {
			return number
		}

		lastN[number]++
		oldIndex := numbers[index-n]
		lastN[oldIndex]--
		if lastN[oldIndex] <= 0 {
			delete(lastN, oldIndex)
		}
	}
	return -1
}

func part2(target int) int {
	// target := part1()

	input, _ := ioutil.ReadFile("input.txt")
	numbers := utils.GetInts(string(input))

	for checkLen := 2; checkLen < len(numbers); checkLen++ {
		for i := 0; i < len(numbers)-checkLen; i++ {
			checkRange := numbers[i : i+checkLen]
			sum := 0
			min, max := utils.MaxInt, 0
			for _, num := range checkRange {
				sum += num
				if min > num {
					min = num
				}
				if max < num {
					max = num
				}
			}
			if sum == target {
				return min + max
			}
		}
	}
	return -1
}

func part2alt(target int) int {
	input, _ := ioutil.ReadFile("input.txt")
	numbers := utils.GetInts(string(input))

	maxSearch := 18
	minSearch := 2
	sums := map[int]int{}
	for index, number := range numbers {
		for rangeLen, rangeSum := range sums {
			if rangeSum == target {
				numRange := numbers[index-rangeLen : index]
				min, max := utils.MaxInt, 0
				for _, num := range numRange {
					if min > num {
						min = num
					}
					if max < num {
						max = num
					}
				}
				return min + max
			}
		}
		for i := maxSearch; i >= minSearch; i-- {
			sums[i] += number
			if index-i >= 0 {
				sums[i] -= numbers[index-i]
			}
		}
	}
	return -1
}
