package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	regex := strings.TrimSpace(string(inputBytes))
	b := newBase(regex)
	p1 := part1(b)
	p2 := part2(b)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(b base) int {
	max := 0
	for _, dist := range b.grid {
		if dist > max {
			max = dist
		}
	}
	return max
}

func part2(b base) int {
	paths := 0

	for _, dist := range b.grid {
		if dist >= 1000 {
			paths++
		}
	}
	return paths
}

type base struct {
	regex string
	grid  map[point]int
}

func (b *base) walk() {
	stack := []point{}
	pos := point{}
	for _, ch := range b.regex {
		switch ch {
		case '(':
			stack = append(stack, pos)
		case ')':
			pos, stack = stack[len(stack)-1], stack[:len(stack)-1]
		case '|':
			pos = stack[len(stack)-1]
		default:
			newDist := b.grid[pos] + 1
			dir := dirs[ch]
			pos = point{pos.x + dir.x, pos.y + dir.y}
			if prevDist, found := b.grid[pos]; found {
				if prevDist < newDist {
					newDist = prevDist
				}
			}
			b.grid[pos] = newDist
		}
	}
}

var dirs = map[rune]point{'N': {0, 1}, 'S': {0, -1}, 'E': {1, 0}, 'W': {-1, 0}}

type point struct {
	x, y int
}

func newBase(routes string) base {
	routes = strings.TrimPrefix(routes, "^")
	routes = strings.TrimSuffix(routes, "$")
	b := base{
		regex: routes,
		grid:  map[point]int{},
	}
	b.walk()
	return b
}

var benchmark = false
