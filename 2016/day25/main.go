package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	for i := 0; i < 10000; i++ {
		result := string(run("input.txt", map[string]int{"a": i}))
		if result == "0101010101" {
			return i
		}
	}
	return -1
}

type instruction struct {
	op   string
	args []string
}

func run(filename string, registers map[string]int) []byte {
	lines := utils.ReadInputLines(filename)

	output := make([]byte, 0, 10)
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
		case "out":
			output = append(output, '0'+byte(getVal(inst.args[0])))
			if len(output) == 10 {
				return output
			}
		default:
			panic(fmt.Errorf("Uhandled instruction: %v", inst.op))
		}
		ip++
	}
	return []byte{}
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
