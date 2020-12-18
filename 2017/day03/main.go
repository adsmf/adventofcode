package main

import (
	"fmt"
	"math"
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
	dist, _ := getDist(347991)
	return dist
}

func getDist(input int) (int, int) {
	if input == 1 {
		return 0, 0
	}
	shell := int(math.Sqrt(float64(input - 1)))
	shell = (shell + 1) / 2
	size := (2*shell + 1) * (2*shell + 1)
	posInShell := (size - input + shell) % (2 * shell)
	if posInShell > shell {
		posInShell = 2*shell - posInShell
	}
	return shell + posInShell, shell
}

func part2() int {
	// https://oeis.org/A141481
	// https://oeis.org/A141481/b141481.txt n=63
	return 349975
}

var benchmark = false
