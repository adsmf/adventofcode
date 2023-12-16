package main

import (
	"fmt"
	"io/ioutil"

	"github.com/adsmf/adventofcode/utils/intcode"
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

var benchmark = false
