package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := run("input.txt", 0)
	p2 := run("input.txt", 1)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func run(filename string, flag int) int {
	mulTimes := 0
	registers := map[string]int{"a": flag}

	lines := utils.ReadInputLines(filename)
	instructions := make([]instructionDetail, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, " ")
		instructions = append(instructions, instructionDetail{parts[0], parts[1], parts[2], 0})
	}

	getVal := func(input string) int {
		val, err := strconv.Atoi(input)
		if err == nil {
			return val
		}
		return registers[input]
	}

	maxSteps := 10000000

	for ip, step := 0, 0; ip >= 0 && ip < len(lines) && step < maxSteps; ip, step = ip+1, step+1 {
		inst := instructions[ip]
		instructions[ip].timesExecuted++

		if flag > 0 && ip == 8 {
			h := 0
			b := registers["b"]
			c := registers["c"]
			step := -1 * getVal(instructions[30].arg2)
			for ; b <= c; b += step {
				if !prime(b) {
					h++
				}
			}
			return h
		}
		switch inst.op {
		case "set":
			registers[inst.arg1] = getVal(inst.arg2)
		case "jnz":
			if getVal(inst.arg1) != 0 {
				ip += getVal(inst.arg2) - 1
			}
		case "sub":
			registers[inst.arg1] -= getVal(inst.arg2)
		case "mul":
			mulTimes++
			registers[inst.arg1] *= getVal(inst.arg2)
		default:
			panic(fmt.Errorf("Unhandled operation: %s", inst.op))
		}
	}

	return mulTimes
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

type instructionDetail struct {
	op            string
	arg1, arg2    string
	timesExecuted int
}

var benchmark = false
