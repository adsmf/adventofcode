package main

import (
	"fmt"
)

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	state := stateA
	maxSteps := 12994925
	tape := map[int]int{}
	cursor := 0
	for i := 0; i < maxSteps; i++ {
		act := stateActions[state][tape[cursor]]
		tape[cursor] = act.write
		cursor += act.move
		state = act.nextState
	}

	checksum := 0
	for _, loc := range tape {
		if loc == 1 {
			checksum++
		}
	}
	return checksum
}

var stateActions = map[int]stateSettings{
	stateA: {
		0: action{1, 1, stateB},
		1: action{0, -1, stateF},
	},
	stateB: {
		0: action{0, 1, stateC},
		1: action{0, 1, stateD},
	},
	stateC: {
		0: action{1, -1, stateD},
		1: action{1, 1, stateE},
	},
	stateD: {
		0: action{0, -1, stateE},
		1: action{0, -1, stateD},
	},
	stateE: {
		0: action{0, 1, stateA},
		1: action{1, 1, stateC},
	},
	stateF: {
		0: action{1, -1, stateA},
		1: action{1, 1, stateA},
	},
}

type stateSettings map[int]action

type action struct {
	write     int
	move      int
	nextState int
}

const (
	stateA = iota
	stateB
	stateC
	stateD
	stateE
	stateF
)

var benchmark = false
