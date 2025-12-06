package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	width := 0
	for i, ch := range input {
		if ch == '\n' {
			width = i
			break
		}
	}
	lines := len(input) / width
	p1 := part1(width, lines)
	p2 := part2(width, lines)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(width, lines int) int {
	total := 0
	probTotal := 0
	colStart := 0
	for colStart < width {
		op := input[(lines-1)*(width+1)+colStart]
		colEnd := -1
		for col := colStart + 1; col <= width; col++ {
			nextOp := input[(lines-1)*(width+1)+col]
			if nextOp != ' ' || nextOp == '\n' {
				colEnd = col - 1
				break
			}
		}
		if colEnd == -1 {
			break
		}
		switch op {
		case '+':
			probTotal = 0
		case '*':
			probTotal = 1
		}
		for lineIdx := range lines - 1 {
			acc := 0
			for col := colStart; col <= colEnd; col++ {
				v := input[lineIdx*(width+1)+col]
				if v == ' ' {
					continue
				}
				acc *= 10
				acc += int(v) - '0'
			}
			switch op {
			case '+':
				probTotal += acc
			case '*':
				probTotal *= acc
			}
		}
		total += probTotal
		colStart = colEnd + 1
	}
	return total
}

func part2(width, lines int) int {
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
