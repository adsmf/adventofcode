package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	vm := load("input.txt")
	p1 := part1(vm)
	p2 := part2(vm)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(vm *virtualMachine) int {
	vm.run()
	return vm.accumulator
}

func part2(vm *virtualMachine) int {
	for i := 0; i < len(vm.program); i++ {
		vm.reset()
		original := vm.program[i].commandCode
		switch vm.program[i].commandCode {
		case "nop":
			vm.program[i].commandCode = "jmp"
		case "jmp":
			vm.program[i].commandCode = "nop"
		default:
			continue
		}
		vm.run()
		if vm.ip >= len(vm.program) {
			return vm.accumulator
		}
		vm.program[i].commandCode = original
	}
	return -1
}

func load(filename string) *virtualMachine {
	vm := &virtualMachine{}
	for _, line := range utils.ReadInputLines(filename) {
		parts := strings.Split(line, " ")
		argument, _ := strconv.Atoi(parts[1])
		inst := instruction{
			commandCode: parts[0],
			argument:    argument,
		}
		vm.program = append(vm.program, inst)
	}
	return vm
}

type virtualMachine struct {
	ip          int
	accumulator int
	program     programListing
}

func (v *virtualMachine) run() {
	alreadyRun := map[int]bool{}
	for {
		if v.ip < 0 || v.ip >= len(v.program) {
			return
		}
		if _, found := alreadyRun[v.ip]; found {
			return
		}
		alreadyRun[v.ip] = true
		v.step()
	}
}
func (v *virtualMachine) step() {
	next := v.program[v.ip]
	switch next.commandCode {
	case "nop":
		v.ip++
	case "acc":
		v.accumulator += next.argument
		v.ip++
	case "jmp":
		v.ip += next.argument
	}
}
func (v *virtualMachine) reset() {
	v.ip = 0
	v.accumulator = 0
}

type programListing []instruction
type instruction struct {
	commandCode string
	argument    int
}
