package main

import (
	"fmt"
	"strings"

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
	minVal := run(0, true, false)
	return minVal
}

func part2() int {
	maxVal := run(0, false, false)
	return maxVal
}

const maxSteps = 1 << 32

func run(reg0 int, part1 bool, debug bool) int {
	registers := [6]int{0: reg0}
	instructions, ipReg := load("input.txt")
	step := 0
	previousReg3 := map[int]bool{}
	lastSeen := 0
	for registers[ipReg] = 0; registers[ipReg] >= 0 && registers[ipReg] < len(instructions); registers[ipReg]++ {
		step++
		if step > maxSteps {
			break
		}
		if registers[ipReg] == 28 {
			reg3 := registers[3]
			if part1 {
				return reg3
			} else {
				if _, seen := previousReg3[reg3]; seen {
					return lastSeen
				}
				lastSeen = reg3
				previousReg3[reg3] = true
			}
			if debug {
				fmt.Printf("Reg 3: %d\n", registers[3])
			}
		}
		instruction := instructions[registers[ipReg]]
		instructions[registers[ipReg]].timesExecuted++
		a, b, c := instruction.args[0], instruction.args[1], instruction.args[2]
		switch instruction.op {
		case "addi":
			registers[c] = registers[a] + b
		case "addr":
			registers[c] = registers[a] + registers[b]
		case "bani":
			registers[c] = registers[a] & b
		case "bori":
			registers[c] = registers[a] | b
		case "eqri":
			if registers[a] == b {
				registers[c] = 1
			} else {
				registers[c] = 0
			}
		case "eqrr":
			if registers[a] == registers[b] {
				registers[c] = 1
			} else {
				registers[c] = 0
			}
		case "gtir":
			if a > registers[b] {
				registers[c] = 1
			} else {
				registers[c] = 0
			}
		case "gtrr":
			if registers[a] > registers[b] {
				registers[c] = 1
			} else {
				registers[c] = 0
			}
		case "muli":
			registers[c] = registers[a] * b
		case "mulr":
			registers[c] = registers[a] * registers[b]
		case "seti":
			registers[c] = a
		case "setr":
			registers[c] = registers[a]
		default:
			panic(fmt.Errorf("Unhandled instruction: %s", instruction.op))
		}
	}
	if debug {
		if step <= maxSteps {
			fmt.Println("Halted properly")
		}
		for idx, inst := range instructions {
			fmt.Printf("%2d: %10d - %s %v\n", idx, inst.timesExecuted, inst.op, inst.args)
		}
	}
	return step
}

type instructionSpec struct {
	op            string
	args          []int
	timesExecuted int
}

func load(filename string) ([]instructionSpec, int) {
	lines := utils.ReadInputLines(filename)
	instructions := make([]instructionSpec, 0, len(lines)-1)
	opReg := int(lines[0][4] - '0')

	for _, line := range lines[1:] {
		parts := strings.Split(line, " ")
		args := make([]int, len(parts)-1)
		for idx, arg := range parts[1:] {
			args[idx] = utils.MustInt(arg)
		}
		instructions = append(instructions, instructionSpec{
			op:   parts[0],
			args: args,
		})
	}

	return instructions, opReg
}

var benchmark = false
var test = false
