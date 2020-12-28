package main

import (
	"fmt"
	"strconv"
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
	return run("input.txt", map[string]int{"a": 7})
}

func part2() int {
	// input but modified to use newfangled `mul` instruction
	return run("input-mod.txt", map[string]int{"a": 12})
}

type instruction struct {
	op   string
	args []string
}

func run(filename string, registers map[string]int) int {
	lines := utils.ReadInputLines(filename)

	instructions := decode(lines)

	getVal := func(input string) int {
		val, err := strconv.Atoi(input)
		if err == nil {
			return val
		}
		return registers[input]
	}

	for ip := 0; ip >= 0 && ip < len(instructions); {
		inst := instructions[ip]
		switch inst.op {
		case "cpy":
			_, err := strconv.Atoi(inst.args[1])
			if err != nil {
				registers[inst.args[1]] = getVal(inst.args[0])
			}
		case "dec":
			_, err := strconv.Atoi(inst.args[0])
			if err != nil {
				registers[inst.args[0]]--
			}
		case "mul":
			registers[inst.args[2]] = getVal(inst.args[0]) * getVal(inst.args[1])
		case "inc":
			_, err := strconv.Atoi(inst.args[0])
			if err != nil {
				registers[inst.args[0]]++
			}
		case "jnz":
			if getVal(inst.args[0]) != 0 {
				ip += getVal(inst.args[1])
				ip--
			}
		case "tgl":
			target := ip + getVal(inst.args[0])
			if target > 0 && target < len(instructions) {
				switch instructions[target].op {
				case "inc":
					instructions[target].op = "dec"
				case "dec", "tgl":
					instructions[target].op = "inc"
				case "jnz":
					instructions[target].op = "cpy"
				case "cpy": //...
					instructions[target].op = "jnz"
				default:
					panic(fmt.Errorf("Don't know how to toggle: %s", instructions[target].op))
				}
			}
		case "nop":
		default:
			panic(fmt.Errorf("Uhandled instruction: %v", inst.op))
		}
		ip++
	}
	return registers["a"]
}

func decode(input []string) []instruction {
	instructions := []instruction{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		instructions = append(instructions, instruction{
			op:   parts[0],
			args: parts[1:],
		})
	}
	return instructions
}

var benchmark = false
