package main

import (
	"fmt"
)

const d23input = "952316487"

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() string {
	game := newGame(d23input, len(d23input))
	for i := 0; i < 100; i++ {
		game.iterate()
	}
	return game.after1()
}

func part2() int {
	game := newGame(d23input, 1000000)
	for i := 0; i < 10000000; i++ {
		game.iterate()
	}
	n1 := game.cups[1]
	n2 := game.cups[n1]
	return n1 * n2
}

type gameData struct {
	cups []int
	cur  int
	max  int
}

func newGame(start string, numCups int) gameData {
	first := int(start[0] - '0')
	g := gameData{
		cups: make([]int, numCups+1),
		cur:  first,
		max:  numCups,
	}

	prev := first
	for i := 1; i < len(start); i++ {
		cur := int(start[i] - '0')
		g.cups[prev] = cur
		prev = cur
	}
	if numCups >= 10 {
		for i := 10; i <= numCups; i++ {
			g.cups[prev] = i
			prev = i
		}
		prev = numCups
	}
	g.cups[prev] = first
	return g
}

func (g *gameData) after1() string {
	res := make([]byte, len(g.cups)-1)

	lookup := 1
	for i := 0; i < len(g.cups)-1; i++ {
		next := g.cups[lookup]
		res[i] = byte(next + '0')
		lookup = next
	}
	return string(res[:len(res)-1])
}

func (g *gameData) iterate() {
	// Cut
	n1 := g.cups[g.cur]
	n2 := g.cups[n1]
	n3 := g.cups[n2]
	g.cups[g.cur] = g.cups[n3]

	// Search
	target := g.cur - 1
	if target < 1 {
		target += g.max
	}
	for target == n1 || target == n2 || target == n3 {
		target--
		if target < 1 {
			target += g.max
		}
	}

	// Insert
	targetNext := g.cups[target]
	g.cups[target] = n1
	g.cups[n3] = targetNext

	// Move on
	g.cur = g.cups[g.cur]
}

var benchmark = false
