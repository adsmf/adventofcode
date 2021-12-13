package main

import (
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

// Alternative implementation
func foldIterative(in string) (int, string) {
	g, folds := load(in)
	p1 := 0
	for i, fold := range folds {
		g = g.fold(fold)
		if i == 0 {
			p1 = len(g)
		}
	}
	return p1, g.String()
}

func load(in string) (grid, []foldInstruction) {
	blocks := strings.Split(in, "\n\n")
	pointLines := strings.Split(blocks[0], "\n")
	g := make(grid, len(pointLines))
	for _, line := range pointLines {
		coords := utils.GetInts(line)
		g[point{coords[0], coords[1]}] = true
	}
	foldLines := strings.Split(blocks[1], "\n")
	folds := make([]foldInstruction, 0, len(foldLines))
	for _, line := range foldLines {
		if line == "" {
			continue
		}
		horizontal := false
		if line[11] == 'y' {
			horizontal = true
		}
		axis := utils.MustInt(line[13:])
		folds = append(folds, foldInstruction{horizontal: horizontal, axis: axis})
	}

	return g, folds
}

func (g grid) fold(fold foldInstruction) grid {
	newGrid := make(grid, len(g))
	for pos := range g {
		if fold.horizontal && pos.y > fold.axis {
			pos.y = fold.axis - (pos.y - fold.axis)
		}
		if !fold.horizontal && pos.x > fold.axis {
			pos.x = fold.axis - (pos.x - fold.axis)
		}
		newGrid[pos] = true
	}
	return newGrid
}

type foldInstruction struct {
	horizontal bool
	axis       int
}
