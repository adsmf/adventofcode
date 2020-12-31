package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := run(0)
	p2 := run(1)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func prime(num int) bool {
	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func run(reg0 int) int {
	registers := [6]int{0: reg0}
	instructions, ipReg := load("input.txt")
	maxSteps := 10000000
	step := 0
	for registers[ipReg] = 0; registers[ipReg] >= 0 && registers[ipReg] < len(instructions); registers[ipReg]++ {
		if registers[ipReg] == 1 {
			target := registers[4]
			sum := 0
			sqrt := int(math.Sqrt(float64(target)))
			for i := 1; i <= sqrt; i++ {
				if !prime(i) {
					continue
				}
				if target%i == 0 {
					sum += i + target/i
				}
			}
			return sum
		}
		step++
		if step > maxSteps {
			fmt.Println("Halting after max steps reached")
			break
		}
		instruction := instructions[registers[ipReg]]
		instructions[registers[ipReg]].timesExecuted++
		a, b, c := instruction.args[0], instruction.args[1], instruction.args[2]
		switch instruction.op {
		case "seti":
			registers[c] = a
		case "setr":
			registers[c] = registers[a]
		case "addi":
			registers[c] = registers[a] + b
		case "addr":
			registers[c] = registers[a] + registers[b]
		case "muli":
			registers[c] = registers[a] * b
		case "mulr":
			registers[c] = registers[a] * registers[b]
		case "eqrr":
			if registers[a] == registers[b] {
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
		default:
			panic(fmt.Errorf("Unhandled instruction: %s", instruction.op))
		}
	}
	if step <= maxSteps {
		fmt.Println("Halted properly")
	}
	for idx, inst := range instructions {
		fmt.Printf("%2d: %10d - %s %v\n", idx, inst.timesExecuted, inst.op, inst.args)
	}
	return registers[0]
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
