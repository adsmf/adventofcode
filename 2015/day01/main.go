package main

import (
	"fmt"
	"io/ioutil"
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
	finalFloor, _ := getFloor("input.txt")
	return finalFloor
}

func part2() int {
	_, basement := getFloor("input.txt")
	return basement
}

func getFloor(filename string) (int, int) {
	inp, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	floor := 0
	basementAt := 0
	step := 0
	for _, char := range inp {
		step++
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor < 0 && basementAt == 0 {
			basementAt = step
		}
	}
	return floor, basementAt
}

var benchmark = false
