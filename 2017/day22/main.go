package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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
	g := load("input.txt")
	count := 0
	var infection bool
	for i := 0; i < 10000; i++ {
		infection = g.step(false)
		if infection {
			count++
		}
	}
	return count
}

func part2() int {
	g := load("input.txt")
	count := 0
	var infection bool
	for i := 0; i < 10000000; i++ {
		infection = g.step(true)
		if infection {
			count++
		}
	}
	return count
}

type virusState byte

const (
	virusClean virusState = iota
	virusWeakened
	virusInfected
	virusFlagged
	virusMAX
)

type virusMap struct {
	grid map[vector]virusState
	cPos vector
	cDir vector
}

func (vm *virusMap) step(p2 bool) bool {
	startState := vm.grid[vm.cPos]
	var nextState virusState
	if p2 {
		nextState = (startState + 1) % virusMAX
		switch startState {
		case virusClean:
			vm.cDir = vector{vm.cDir.y, -vm.cDir.x} //Left
		case virusWeakened:
		case virusInfected:
			vm.cDir = vector{-vm.cDir.y, vm.cDir.x} //Right
		case virusFlagged:
			vm.cDir = vector{-vm.cDir.x, -vm.cDir.y} //Reverse
		}
	} else {
		if startState == virusInfected {
			vm.cDir = vector{-vm.cDir.y, vm.cDir.x} //Right
			nextState = virusClean
		} else {
			vm.cDir = vector{vm.cDir.y, -vm.cDir.x} //Left
			nextState = virusInfected
		}
	}
	vm.grid[vm.cPos] = nextState
	vm.cPos = vm.cPos.add(vm.cDir)
	causedInfection := nextState == virusInfected
	return causedInfection
}

type vector struct{ x, y int }

func (v vector) add(a vector) vector {
	return vector{v.x + a.x, v.y + a.y}
}

func load(filename string) virusMap {
	vm := virusMap{
		grid: map[vector]virusState{},
		cPos: vector{0, 0},
		cDir: vector{0, -1},
	}
	lines := utils.ReadInputLines(filename)
	for i, line := range lines {
		y := i - (len(lines)-1)/2
		for j, char := range line {
			x := j - (len(line)-1)/2
			pos := vector{x, y}
			if char == '#' {
				vm.grid[pos] = virusInfected
			}
		}
	}
	return vm
}

var benchmark = false
