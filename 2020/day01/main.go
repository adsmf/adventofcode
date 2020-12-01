package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	numbers := utils.GetInts(string(inputBytes))
	for fence1, num1 := range numbers {
		for _, num2 := range numbers[fence1+1:] {
			if num1+num2 == 2020 {
				return num1 * num2
			}
		}
	}
	return -1
}

func part2() int {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	numbers := utils.GetInts(string(inputBytes))
	for fence1, num1 := range numbers {
		for fence2, num2 := range numbers[fence1+1:] {
			for _, num3 := range numbers[fence1+fence2+2:] {
				if num1+num2+num3 == 2020 {
					return num1 * num2 * num3
				}
			}
		}
	}
	return -1
}
