package main

import (
	"fmt"
	"github.com/adsmf/adventofcode/utils"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	c := loadInput("input.txt")
	c.run()
	return int(c.registers[regB])
}

func part2() int {
	c := loadInput("input.txt")
	c.registers[regA] = 1
	c.run()
	return int(c.registers[regB])
}

type computer struct {
	ip           int
	instructions []instruction
	registers    map[parameter]uint
}

func (c *computer) run() {
	for {
		if c.ip < 0 || c.ip >= len(c.instructions) {
			return
		}
		curInst := c.instructions[c.ip]
		switch curInst.op {
		case opHalf:
			c.registers[curInst.p1] /= 2
			c.ip++
		case opTriple:
			c.registers[curInst.p1] *= 3
			c.ip++
		case opIncrement:
			c.registers[curInst.p1]++
			c.ip++
		case opJump:
			c.ip += int(curInst.p1)
		case opJumpIfEven:
			if c.registers[curInst.p1]%2 == 0 {
				c.ip += int(curInst.p2)
			} else {
				c.ip++
			}
		case opJumpIfOne:
			if c.registers[curInst.p1] == 1 {
				c.ip += int(curInst.p2)
			} else {
				c.ip++
			}
		}
	}
}

type instruction struct {
	op     operationCode
	p1, p2 parameter
}

type operationCode int

const (
	opUndefined operationCode = iota
	opHalf
	opTriple
	opIncrement
	opJump
	opJumpIfEven
	opJumpIfOne
)

type parameter int

const (
	regUndefined parameter = iota
	regA
	regB
)

func loadInput(filename string) computer {
	c := computer{
		instructions: []instruction{},
		registers:    map[parameter]uint{},
	}

	for _, line := range utils.ReadInputLines(filename) {
		inst := instruction{}
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "hlf":
			inst.op = opHalf
		case "tpl":
			inst.op = opTriple
		case "inc":
			inst.op = opIncrement
		case "jmp":
			inst.op = opJump
		case "jie":
			inst.op = opJumpIfEven
		case "jio":
			inst.op = opJumpIfOne
		}
		inst.p1 = decodeParam(parts[1])
		if len(parts) > 2 {
			inst.p2 = decodeParam(parts[2])
		}
		c.instructions = append(c.instructions, inst)
	}
	return c
}

func decodeParam(param string) parameter {
	param = strings.Trim(param, ",")
	if param == "a" {
		return regA
	}
	if param == "b" {
		return regB
	}
	paramVal, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	return parameter(paramVal)
}
