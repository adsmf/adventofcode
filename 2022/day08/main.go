package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	g := loadGrid()
	edgeVisible := g.getVisible()
	bestScore := 0
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.width; y++ {
			score := g.score(x, y)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return edgeVisible, bestScore
}

type treeGrid struct {
	height, width int
}

type point struct {
	x, y int
}
type tree int

func loadGrid() treeGrid {
	g := treeGrid{}
	for ; input[g.width] != '\n'; g.width++ {
	}
	g.height = len(input)/g.width - 1
	return g
}

func (g treeGrid) tree(x, y int) tree {
	return tree(input[x+y*(g.width+1)] & 0xf)
}

func (g treeGrid) getVisible() int {
	visible := [100][100]bool{}
	// From top
	for x := 0; x < g.width; x++ {
		maxHeight := tree(-1)
		for y := 0; y < g.height; y++ {
			treeHeight := g.tree(x, y)
			if treeHeight > maxHeight {
				visible[x][y] = true
			}
			if treeHeight > maxHeight {
				maxHeight = treeHeight
			}
			if treeHeight == 9 {
				break
			}
		}
	}
	// From bottom
	for x := 0; x < g.width; x++ {
		maxHeight := tree(-1)
		for y := g.height - 1; y >= 0; y-- {
			treeHeight := g.tree(x, y)
			if treeHeight > maxHeight {
				visible[x][y] = true
			}
			if treeHeight > maxHeight {
				maxHeight = treeHeight
			}
			if treeHeight == 9 {
				break
			}
		}
	}
	// Left
	for y := 0; y < g.height; y++ {
		maxHeight := tree(-1)
		for x := 0; x < g.width; x++ {
			treeHeight := g.tree(x, y)
			if treeHeight > maxHeight {
				visible[x][y] = true
			}
			if treeHeight > maxHeight {
				maxHeight = treeHeight
			}
			if treeHeight == 9 {
				break
			}
		}
	}
	// Right
	for y := 0; y < g.height; y++ {
		maxHeight := tree(-1)
		for x := g.width - 1; x >= 0; x-- {
			treeHeight := g.tree(x, y)
			if treeHeight > maxHeight {
				visible[x][y] = true
			}
			if treeHeight > maxHeight {
				maxHeight = treeHeight
			}
			if treeHeight == 9 {
				break
			}
		}
	}
	count := 0
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			if visible[x][y] {
				count++
			}
		}
	}
	return count
}

func (g treeGrid) score(x, y int) int {
	th := g.tree(x, y)
	score := 1

	dirScore := 0
	for cy := y - 1; cy >= 0; cy-- {
		dirScore++
		if g.tree(x, cy) >= th {
			break
		}
	}
	if dirScore == 0 {
		return 0
	}
	score *= dirScore

	dirScore = 0
	for cy := y + 1; cy < g.height; cy++ {
		dirScore++
		if g.tree(x, cy) >= th {
			break
		}
	}
	if dirScore == 0 {
		return 0
	}
	score *= dirScore

	dirScore = 0
	for cx := x - 1; cx >= 0; cx-- {
		dirScore++
		if g.tree(cx, y) >= th {
			break
		}
	}
	if dirScore == 0 {
		return 0
	}
	score *= dirScore

	dirScore = 0
	for cx := x + 1; cx < g.width; cx++ {
		dirScore++
		if g.tree(cx, y) >= th {
			break
		}
	}
	score *= dirScore
	return score
}

var benchmark = false
