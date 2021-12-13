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
	p1, p2 := foldFunctional(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2:\n%s\n", p2)
	}
}

func foldFunctional(in string) (int, string) {
	g, firstFold, subsequentFolds := loadFunctional(in)
	p1Grid := make(grid, len(g))
	p2Grid := make(grid, len(g))
	for pos := range g {
		p1pos := firstFold(pos)
		p1Grid[p1pos] = true
		p2pos := subsequentFolds(p1pos)
		p2Grid[p2pos] = true
	}
	return len(p1Grid), p2Grid.String()
}

func loadFunctional(in string) (grid, foldOperation, foldOperation) {
	blocks := strings.Split(in, "\n\n")
	pointLines := strings.Split(blocks[0], "\n")
	g := make(grid, len(pointLines))
	for _, line := range pointLines {
		coords := utils.GetInts(line)
		g[point{coords[0], coords[1]}] = true
	}
	foldLines := strings.Split(blocks[1], "\n")
	var firstFold, subsequentFolds foldOperation
	subsequentFolds = noFold
	for i, line := range foldLines {
		if line == "" {
			continue
		}
		horizontal := false
		if line[11] == 'y' {
			horizontal = true
		}
		axis := utils.MustInt(line[13:])
		if i == 0 {
			firstFold = addFoldOperation(noFold, horizontal, axis)

			continue
		}
		subsequentFolds = addFoldOperation(subsequentFolds, horizontal, axis)
	}

	return g, firstFold, subsequentFolds
}

type foldOperation func(point) point

func addFoldOperation(previousFolds foldOperation, horizontal bool, axis int) foldOperation {
	if horizontal {
		return func(in point) point {
			in = previousFolds(in)
			if in.y > axis {
				in.y = axis - (in.y - axis)
			}
			return in
		}
	}
	return func(in point) point {
		in = previousFolds(in)
		if in.x > axis {
			in.x = axis - (in.x - axis)
		}
		return in
	}
}

func noFold(in point) point { return in }

type grid map[point]bool

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

var benchmark = false
