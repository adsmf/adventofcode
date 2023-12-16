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
	inputString := loadInputString()
	outputs := gatherOutputs(inputString, 1)
	return outputs[len(outputs)-1]
}

func part2() int {
	inputString := loadInputString()
	outputs := gatherOutputs(inputString, 2)
	return outputs[0]
}

func gatherOutputs(program string, in int) []int {
	outputs := []int{}
	inputCB := func() (int, bool) {
		return in, false
	}
	outputCB := func(out int) {
		outputs = append(outputs, out)
	}

	m := intcode.NewMachine(intcode.M19(inputCB, outputCB))
	m.LoadProgram(program)
	m.Run(false)

	return outputs
}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}

var benchmark = false
