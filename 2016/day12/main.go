package main

import (
	"fmt"
	"strconv"
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

type instruction struct {
	op  operation
	p1s string
	p1i int
	p2s string
	p2i int
}
type operation int

const (
	opCopyVal operation = iota
	opCopyReg
	opInc
	opDec
	opJNZ
)

func run(ignition int) int {
	registers := map[string]int{"1": 1, "c": ignition}
	lines := utils.ReadInputLines("input.txt")

	instructions := decode(lines)

	for ip := 0; ip >= 0 && ip < len(instructions); {
		inst := instructions[ip]
		switch inst.op {
		case opInc:
			registers[inst.p1s]++
		case opDec:
			registers[inst.p1s]--
		case opCopyReg:
			registers[inst.p2s] = registers[inst.p1s]
		case opCopyVal:
			registers[inst.p2s] = inst.p1i
		case opJNZ:
			if registers[inst.p1s] != 0 {
				ip += inst.p2i
				ip--
			}
		}
		ip++
	}
	return registers["a"]
}

func decode(input []string) []instruction {
	instructions := []instruction{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "cpy":
			val, err := strconv.Atoi(parts[1])
			if err == nil {
				instructions = append(instructions, instruction{
					op:  opCopyVal,
					p1i: val,
					p2s: parts[2],
				})
			} else {
				instructions = append(instructions, instruction{
					op:  opCopyReg,
					p1s: parts[1],
					p2s: parts[2],
				})
			}
		case "inc":
			instructions = append(instructions, instruction{
				op:  opInc,
				p1s: parts[1],
			})
		case "dec":
			instructions = append(instructions, instruction{
				op:  opDec,
				p1s: parts[1],
			})
		case "jnz":
			instructions = append(instructions, instruction{
				op:  opJNZ,
				p1s: parts[1],
				p2i: utils.MustInt(parts[2]),
			})
		}
	}
	return instructions
}

var benchmark = false
