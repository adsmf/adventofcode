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
	memory := map[int]int{}
	mask := ""

	for _, line := range utils.ReadInputLines("input.txt") {
		if strings.HasPrefix(line, "mask") {
			mask = strings.TrimPrefix(line, "mask = ")
			continue
		}
		ints := utils.GetInts(line)
		addr, value := ints[0], ints[1]

		memory[addr] = applyMask(mask, value)
	}

	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
}

func applyMask(mask string, value int) int {
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

func part2() int {
	memory := map[int]int{}
	mask := ""

	for _, line := range utils.ReadInputLines("input.txt") {
		if strings.HasPrefix(line, "mask") {
			mask = strings.TrimPrefix(line, "mask = ")
			continue
		}
		ints := utils.GetInts(line)
		addr, value := ints[0], ints[1]
		writeAddresses(memory, mask, addr, value, 0)
	}

	sum := 0
	for _, value := range memory {
		sum += value
	}
	return sum
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

var benchmark = false
