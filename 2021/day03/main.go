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
	values, bitlen := parseInputInts()
	p1 := part1(values, bitlen)
	p2 := part2(values, bitlen)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(values []int, bitlen int) int {
	numNalues := len(values)
	gamma, epsilon := 0, 0
	for i := 0; i < bitlen; i++ {
		count := 0
		for _, val := range values {
			if val&(1<<i) > 0 {
				count++
			}
		}
		if count > numNalues/2 {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	return gamma * epsilon
}

func part2(values []int, bitlen int) int {
	o2, co2 := 0, 0
	generator := append([]int{}, values...)
	for pos := bitlen - 1; pos >= 0; pos-- {
		generator = filter(generator, pos, true)
		if len(generator) == 1 {
			o2 = generator[0]
		}
	}
	scrubber := append([]int{}, values...)
	for pos := bitlen - 1; pos >= 0; pos-- {
		scrubber = filter(scrubber, pos, false)
		if len(scrubber) == 1 {
			co2 = scrubber[0]
		}
	}
	return o2 * co2
}

func filter(values []int, pos int, mostCommon bool) []int {
	count := 0
	searchMask := 1 << pos
	for _, val := range values {
		if val&searchMask > 0 {
			count++
		}
	}
	filterTo := 1 << pos
	if mostCommon != (count<<1 >= len(values)) {
		filterTo = 0
	}
	n := 0
	for _, val := range values {
		if val&searchMask == filterTo {
			values[n] = val
			n++
		}
	}
	return values[:n]
}

func parseInputInts() ([]int, int) {
	lines := utils.GetLines(input)
	values := make([]int, 0, len(lines))
	for _, line := range lines {
		val, _ := strconv.ParseUint(line, 2, 16)
		values = append(values, int(val))
	}
	bitlen := len(lines[0])
	return values, bitlen
}

// Initial implementations (kept for benchmark comparison)
func part1initial() int {
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

func part2initial() int {
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
