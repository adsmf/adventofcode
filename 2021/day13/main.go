package main

import (
	_ "embed"
	"fmt"
	"strings"
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
	g := make(grid, len(in)/8)
	i := 0
	acc := 0
	x := 0
	pointsParsed := false
	for ; !pointsParsed; i++ {
		ch := in[i]
		switch ch {
		case '\n':
			if acc == 0 {
				pointsParsed = true
				break
			}
			g[makePoint(x, acc)] = true
			acc = 0
			x = 0
		case ',':
			x, acc = acc, 0
		default:
			acc = acc*10 + int(ch-'0')
		}
	}
	i += 11
	var firstFold, subsequentFolds foldOperation
	firstFold = noFold
	subsequentFolds = noFold
	first := true
	for ; i < len(input); i++ {
		horizontal := false
		if input[i] == 'y' {
			horizontal = true
		}
		i += 2
		axis := 0
		for ; input[i] != '\n'; i++ {
			axis = axis*10 + int(input[i]-'0')
		}

		if first {
			first = false
			firstFold = addFoldOperation(noFold, horizontal, axis)
		} else {
			subsequentFolds = addFoldOperation(subsequentFolds, horizontal, axis)
		}
		i += 11
	}
	return g, firstFold, subsequentFolds
}

type foldOperation func(point) point

func addFoldOperation(previousFolds foldOperation, horizontal bool, axis int) foldOperation {
	if horizontal {
		return func(in point) point {
			in = previousFolds(in)
			if in.y() > axis {
				in = in.withY(axis - (in.y() - axis))
			}
			return in
		}
	}
	return func(in point) point {
		in = previousFolds(in)
		if in.x() > axis {
			in = in.withX(axis - (in.x() - axis))
		}
		return in
	}
}

func noFold(in point) point { return in }

type grid map[point]bool

func (g grid) String() string {
	maxX, maxY := 0, 0
	for pos := range g {
		if pos.x() > maxX {
			maxX = pos.x()
		}
		if pos.y() > maxY {
			maxY = pos.y()
		}
	}
	sb := strings.Builder{}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if g[makePoint(x, y)] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func makePoint(x, y int) point {
	return point(x | (y << 16))
}

type point uint32

func (p point) x() int {
	return int(p & 0xffff)
}
func (p point) y() int {
	return int(p >> 16)
}
func (p point) withY(y int) point {
	return point(p&0xffff | point(y<<16))
}
func (p point) withX(x int) point {
	return point(p&0xffff0000 | point(x))
}

var benchmark = false
