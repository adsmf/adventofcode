package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	commands := load("input.txt")
	p1 := part1(commands)
	p2 := part2(commands)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(commands []command) int {
	memory := map[int]int{}
	for _, cmd := range commands {
		memory[cmd.address] = applyValueMask(cmd.mask, cmd.value)
	}
	return sum(memory)
}

func applyValueMask(mask string, value int) int {
	setMask := 0
	unsetMask := 0

	for _, char := range mask {
		setMask *= 2
		unsetMask *= 2
		switch char {
		case '1':
			setMask++
		case '0':
			unsetMask++
		}
	}

	value = value&(^unsetMask) | setMask
	return value
}

func part2(commands []command) int {
	memory := make(map[int]int, 100000)

	for _, cmd := range commands {
		writeAddresses(memory, cmd.mask, cmd.address, cmd.value, 0)
	}

	return sum(memory)
}

func writeAddresses(memory map[int]int, mask string, addr int, value, index int) {
	if index == 36 {
		memory[addr] = value
		return
	}
	switch mask[index] {
	case '0':
		writeAddresses(memory, mask, addr, value, index+1)
	case '1':
		addr = addr | 1<<(35-index)
		writeAddresses(memory, mask, addr, value, index+1)
	case 'X':
		addr0 := addr & ^(1 << (35 - index))
		writeAddresses(memory, mask, addr0, value, index+1)
		addr1 := addr | 1<<(35-index)
		writeAddresses(memory, mask, addr1, value, index+1)
	}
}

func sum(memory map[int]int) int {
	sum := 0
	count := 0
	for _, value := range memory {
		sum += value
		count++
	}
	return sum
}

func load(filename string) []command {
	commands := []command{}
	mask := ""
	for _, line := range utils.ReadInputLines("input.txt") {
		if strings.HasPrefix(line, "mask") {
			mask = strings.TrimPrefix(line, "mask = ")
			continue
		}
		ints := utils.GetInts(line)
		commands = append(commands, command{
			mask:    mask,
			address: ints[0],
			value:   ints[1],
		})
	}
	return commands
}

type command struct {
	mask    string
	address int
	value   int
}

var benchmark = false
