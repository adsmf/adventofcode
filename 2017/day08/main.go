package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1, p2 := run()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func run() (int, int) {
	registers := map[string]int{}
	maxReg := 0
	for _, line := range utils.ReadInputLines("input.txt") {
		reg, op, amount, condReg, condOp, condAmount := "", "", 0, "", "", 0
		fmt.Sscanf(line, "%s %s %d if %s %s %d", &reg, &op, &amount, &condReg, &condOp, &condAmount)
		var conditionMatches bool
		condRegVal := registers[condReg]
		switch condOp {
		case "<":
			conditionMatches = condRegVal < condAmount
		case "<=":
			conditionMatches = condRegVal <= condAmount
		case ">":
			conditionMatches = condRegVal > condAmount
		case ">=":
			conditionMatches = condRegVal >= condAmount
		case "==":
			conditionMatches = condRegVal == condAmount
		case "!=":
			conditionMatches = condRegVal != condAmount
		}
		if conditionMatches {
			switch op {
			case "inc":
				registers[reg] += amount
			case "dec":
				registers[reg] -= amount
			}
		}
		if registers[reg] > maxReg {
			maxReg = registers[reg]
		}
	}
	fmt.Println(registers)
	max := 0
	for _, val := range registers {
		if val > max {
			max = val
		}
	}
	return max, maxReg
}

var benchmark = false
