package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	return getFloor("input.txt")
}

func part2() int {
	return 0
}

func getFloor(filename string) int {
	inp, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	floor := 0
	for _, char := range inp {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
	}
	return floor
}
