package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := fold(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2:\n%s\n", p2)
	}
}

func fold(in string) (int, string) {
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

type grid map[point]bool

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

func (g grid) String() string {
	maxX, maxY := 0, 0
	for pos := range g {
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	sb := strings.Builder{}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if g[point{x, y}] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

type point struct{ x, y int }
type foldInstruction struct {
	horizontal bool
	axis       int
}

var benchmark = false
