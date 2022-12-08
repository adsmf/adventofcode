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
func (g treeGrid) treeAt(offset int) tree { return tree(input[offset]) }

func solve() (int, int) {
	g := loadGrid()
	edgeVisible := 0
	bestScore := 0
	offset := 0
	maxY := g.height * (g.width + 1)
	for y := 0; y < g.width; y++ {
		for x := 0; x < g.width; x++ {
			th := g.treeAt(offset)
			sL, eL := g.look(th, offset, -1, offset-x-1)
			sR, eR := g.look(th, offset, 1, offset+g.width-x)
			sU, eU := g.look(th, offset, -g.width-1, x-g.width-1)
			sD, eD := g.look(th, offset, g.width+1, x+maxY)
			if eL || eR || eU || eD {
				edgeVisible++
			}
			score := sL * sR * sU * sD
			if score > bestScore {
				bestScore = score
			}
			offset++
		}
		offset++
	}
	return edgeVisible, bestScore
}

func (g treeGrid) look(stopHeight tree, offset, inc, bound int) (int, bool) {
	count := 0
	for offset += inc; offset != bound; offset += inc {
		count++
		th := g.treeAt(offset)
		if th >= tree(stopHeight) {
			return count, false
		}
	}
	return count, true
}

var benchmark = false
