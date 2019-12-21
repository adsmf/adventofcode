package main

import (
	"fmt"
	"github.com/adsmf/adventofcode2019/utils/intcode"
	"io/ioutil"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	s := springdroid{}
	s.instructions = []string{
		"NOT A J\n",
		"NOT B T\n",
		"OR T J\n",
		"NOT C T\n",
		"OR T J\n",
		"AND D J\n",
		"WALK\n",
	}
	return s.run()
}

func part2() int {
	return 0
}

type springdroid struct {
	curInst      string
	instructions []string
	finalValue   int
}

func (s *springdroid) run() int {
	m := intcode.NewMachine(intcode.M19(s.inputHandler, s.outputHandler))
	prog, _ := ioutil.ReadFile("input.txt")
	m.LoadProgram(string(prog))
	m.Run(false)
	return s.finalValue
}

func (s *springdroid) inputHandler() (int, bool) {
	if s.curInst == "" && len(s.instructions) > 0 {
		s.curInst, s.instructions = s.instructions[0], s.instructions[1:]
	}
	if s.curInst != "" {
		var nextChar byte
		nextChar, s.curInst = s.curInst[0], s.curInst[1:]
		return int(nextChar), false
	}
	return 0, false
}

func (s *springdroid) outputHandler(op int) {
	if op < 255 {
		fmt.Printf("%c", op)
	} else {
		s.finalValue = op
		fmt.Printf("Final value: %d\n", op)
	}
}
