package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode2019/utils/intcode"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	return runInput(1)
}

func part2() int {
	return runInput(5)
}

func runInput(input int) int {
	inputCB := func() (int, bool) {
		return input, false
	}
	m := intcode.NewMachine(intcode.M19(inputCB, nil))
	m.LoadProgram(loadInputString())
	m.Run(false)
	return m.Register(intcode.M19RegisterOutput)
}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}
