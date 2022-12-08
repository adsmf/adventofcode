package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type treeGrid struct {
	height, width int
}

type tree uint8

func loadGrid() treeGrid {
	g := treeGrid{}
	for ; input[g.width] != '\n'; g.width++ {
		input[g.width] &= 0xf
	}
	g.height = len(input)/g.width - 1
	for i := g.width + 1; i < len(input); i++ {
		input[i] &= 0xf
	}
	return g
}

func (g treeGrid) tree(x, y int) tree { return tree(input[x+y*(g.width+1)]) }

func solve() (int, int) {
	g := loadGrid()
	edgeVisible := 0
	bestScore := 0
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.width; y++ {
			th := g.tree(x, y)
			sL, eL := g.lookHorizontal(th, x, y, -1)
			sR, eR := g.lookHorizontal(th, x, y, 1)
			sU, eU := g.lookVertical(th, x, y, -1)
			sD, eD := g.lookVertical(th, x, y, 1)
			if eL || eR || eU || eD {
				edgeVisible++
			}
			score := sL * sR * sU * sD
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return edgeVisible, bestScore
}

func (g treeGrid) lookHorizontal(stopHeight tree, x, y, dx int) (int, bool) {
	countLE, seesEdge := 0, true
	bound := -1
	if dx > 0 {
		bound = g.height
	}
	for x += dx; x != bound; x += +dx {
		countLE++
		th := g.tree(x, y)
		if th >= tree(stopHeight) {
			seesEdge = false
			break
		}
	}
	return countLE, seesEdge
}

func (g treeGrid) lookVertical(stopHeight tree, x, y, dy int) (int, bool) {
	countLE, seesEdge := 0, true
	bound := -1
	if dy > 0 {
		bound = g.height
	}
	for y += dy; y != bound; y += +dy {
		countLE++
		th := g.tree(x, y)
		if th >= tree(stopHeight) {
			seesEdge = false
			break
		}
	}
	return countLE, seesEdge
}

var benchmark = false
