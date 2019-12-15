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
	return runInput(12, 2)
}

func part2() int {
	target := 19690720
	for input1 := 0; input1 <= 99; input1++ {
		for input2 := 0; input2 <= 99; input2++ {
			result := runInput(input1, input2)
			if result == target {
				return input1*100 + input2
			}
		}
	}
	return 0
}

func runInput(input1, input2 int) int {
	m := intcode.NewMachine(intcode.M19(nil, nil))
	m.LoadProgram(loadInputString())
	m.WriteRAM(1, input1)
	m.WriteRAM(2, input2)
	m.Run(false)
	return m.ReadRAM(0)
}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}
