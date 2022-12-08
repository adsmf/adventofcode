package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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
	for pos := range g.trees {
		score := g.score(pos)
		if score > bestScore {
			bestScore = score
		}
	}
	return edgeVisible, bestScore
}

type treeGrid struct {
	trees         map[point]tree
	height, width int
}

type point struct {
	x, y int
}
type tree int

func loadGrid() treeGrid {
	lines := utils.GetLines(input)
	g := treeGrid{
		height: len(lines),
		width:  len(lines[0]),
	}
	g.trees = make(map[point]tree, g.height*g.width)
	for y, line := range lines {
		for x, ch := range line {
			g.trees[point{x, y}] = tree(ch - '0')
		}
	}
	return g
}

func (g treeGrid) getVisible() int {
	visible := map[point]bool{}
	// From top
	for x := 0; x < g.width; x++ {
		maxHeight := tree(-1)
		for y := 0; y < g.height; y++ {
			pos := point{x, y}
			treeHeight := g.trees[pos]
			if treeHeight > maxHeight {
				visible[pos] = true
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
			pos := point{x, y}
			treeHeight := g.trees[pos]
			if treeHeight > maxHeight {
				visible[pos] = true
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
			pos := point{x, y}
			treeHeight := g.trees[pos]
			if treeHeight > maxHeight {
				visible[pos] = true
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
			pos := point{x, y}
			treeHeight := g.trees[pos]
			if treeHeight > maxHeight {
				visible[pos] = true
			}
			if treeHeight > maxHeight {
				maxHeight = treeHeight
			}
			if treeHeight == 9 {
				break
			}
		}
	}
	return len(visible)
}

func (g treeGrid) score(pos point) int {
	sU, sD, sL, sR := 0, 0, 0, 0
	th := g.trees[pos]
	for y := pos.y - 1; y >= 0; y-- {
		check := point{pos.x, y}
		checkHeight := g.trees[check]
		sU++
		if checkHeight >= th {
			break
		}
	}
	for y := pos.y + 1; y < g.height; y++ {
		check := point{pos.x, y}
		checkHeight := g.trees[check]
		sD++
		if checkHeight >= th {
			break
		}
	}
	for x := pos.x + 1; x < g.width; x++ {
		check := point{x, pos.y}
		checkHeight := g.trees[check]
		sR++
		if checkHeight >= th {
			break
		}
	}
	for x := pos.x - 1; x >= 0; x-- {
		check := point{x, pos.y}
		checkHeight := g.trees[check]
		sL++
		if checkHeight >= th {
			break
		}
	}
	return sU * sD * sL * sR
}

var benchmark = false
