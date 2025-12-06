package main

import (
	_ "embed"
	"fmt"

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
	lineInts := [][]int{}
	total := 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		if line[0] >= '0' && line[0] <= '9' || line[0] == ' ' {
			lineInts = append(lineInts, utils.GetInts(line))
			return
		}
		idx := 0
		for _, ch := range line {
			switch ch {
			case ' ':
				continue
			case '*':
				mul := 1
				for i := range len(lineInts) {
					mul *= lineInts[i][idx]
				}
				total += mul
			case '+':
				sum := 0
				for i := range len(lineInts) {
					sum += lineInts[i][idx]
				}
				total += sum
			}
			idx++

		}
		return
	})
	return total
}

func part2() int {
	width := 0
	for i, ch := range input {
		if ch == '\n' {
			width = i
			break
		}
	}
	lines := len(input) / width
	operation := byte('?')
	total := 0
	probTotal := 0
	for col := range width {
		acc := 0
		for lineIdx := range lines - 1 {
			v := input[lineIdx*(width+1)+col]
			if v == ' ' {
				continue
			}
			acc *= 10
			acc += int(v - '0')
		}
		if acc == 0 {
			total += probTotal
			probTotal = 0
			continue
		}
		newOp := input[(lines-1)*(width+1)+col]
		if newOp != ' ' {
			switch newOp {
			case '*':
				probTotal = 1
			case '+':
				probTotal = 0
			}
			operation = newOp
		}

		switch operation {
		case '*':
			probTotal *= acc
		case '+':
			probTotal += acc
		}
	}
	total += probTotal
	return total
}

var benchmark = false
