package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

var benchmark = false

func main() {
	p1 := part1("input.txt")
	p2 := part2("input.txt")
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

type key struct{ x, y int }
type keypad map[key]rune

func part1(filename string) string {
	pad := keypad{
		{-1, -1}: '1', {0, -1}: '2', {1, -1}: '3',
		{-1, 0}: '4', {0, 0}: '5', {1, 0}: '6',
		{-1, 1}: '7', {0, 1}: '8', {1, 1}: '9',
	}
	return process(filename, pad)
}

func part2(filename string) string {
	pad := keypad{
		{0, 0}:  '5',
		{1, -1}: '2',
		{1, 0}:  '6',
		{1, 1}:  'A',
		{2, -2}: '1',
		{2, -1}: '3',
		{2, 0}:  '7',
		{2, 1}:  'B',
		{2, 2}:  'D',
		{3, -1}: '4',
		{3, 0}:  '8',
		{3, 1}:  'C',
		{4, 0}:  '9',
	}
	return process(filename, pad)
}
func process(filename string, pad keypad) string {
	code := ""
	pos := key{0, 0}
	for _, line := range utils.ReadInputLines(filename) {
		for _, char := range line {
			next := pos
			switch char {
			case 'U':
				next.y--
			case 'D':
				next.y++
			case 'L':
				next.x--
			case 'R':
				next.x++
			}
			if _, found := pad[next]; !found {
				continue
			}
			pos = next
		}
		code = fmt.Sprintf("%s%c", code, pad[pos])
	}
	return code
}
