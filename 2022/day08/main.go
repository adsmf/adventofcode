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

type tree uint8

func solve() (int, int) {
	g := loadGrid()
	edgeVisible, bestScore := 0, 0
	maxY := g.height * (g.width + 1)
	for offset := 0; offset < len(input); offset++ {
		lineStart := offset
		for x := 0; x < g.width; x, offset = x+1, offset+1 {
			th := g.treeAt(offset)
			sL, eL := g.look(th, offset, -1, lineStart-1)
			sR, eR := g.look(th, offset, 1, lineStart+g.width)
			sU, eU := g.look(th, offset, -g.width-1, x-g.width-1)
			sD, eD := g.look(th, offset, g.width+1, x+maxY)
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

func loadGrid() treeGrid {
	g := treeGrid{}
	for ; input[g.width] != '\n'; g.width++ {
	}
	g.height = len(input)/g.width - 1
	return g
}

type treeGrid struct {
	height, width int
}

func (g treeGrid) treeAt(offset int) tree { return tree(input[offset]) }
func (g treeGrid) look(stopHeight tree, pos, inc, bound int) (int, bool) {
	count := 0
	for pos += inc; pos != bound; pos += inc {
		count++
		th := g.treeAt(pos)
		if th >= tree(stopHeight) {
			return count, false
		}
	}
	return count, true
}

var benchmark = false
