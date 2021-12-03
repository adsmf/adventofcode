package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	lines := utils.GetLines(input)
	bitlen := len(lines[0])
	counts := make([]int, bitlen)
	for _, line := range lines {
		for i, ch := range line {
			if ch == '1' {
				counts[i]++
			}
		}
	}

	avg := len(lines) / 2
	gamma, epsilon := 0, 0
	for i := 0; i < bitlen; i++ {
		if counts[i] > avg {
			gamma |= (1 << (bitlen - 1)) >> i
		} else {
			epsilon |= (1 << (bitlen - 1)) >> i
		}
	}
	return gamma * epsilon
}

func part2() int {
	lines := utils.GetLines(input)
	bitlen := len(lines[0])
	oxygen := ""
	for i := 0; i < bitlen; i++ {
		common := mostCommonForBit(lines, i)

		n := 0
		for _, line := range lines {
			if line[i] == byte(common) {
				lines[n] = line
				n++
			}
		}
		lines = lines[:n]
		if len(lines) == 1 {
			oxygen = lines[0]
			break
		}
	}
	lines = utils.GetLines(input)
	co2 := ""
	for i := 0; i < bitlen; i++ {
		common := mostCommonForBit(lines, i)

		n := 0
		for _, line := range lines {
			if line[i] != byte(common) {
				lines[n] = line
				n++
			}
		}
		lines = lines[:n]
		if len(lines) == 1 {
			co2 = lines[0]
			break
		}
	}
	o2int, _ := strconv.ParseInt(oxygen, 2, 16)
	co2int, _ := strconv.ParseInt(co2, 2, 16)
	return int(o2int) * int(co2int)
}

func mostCommonForBit(lines []string, pos int) rune {
	count := 0
	for _, line := range lines {
		if line[pos] == '1' {
			count++
		}
	}
	if float64(count) >= (float64(len(lines)) / 2) {
		return '1'
	}
	return '0'
}

var benchmark = false
