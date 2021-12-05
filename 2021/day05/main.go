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
	p1, p2 := loadInput()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func loadInput() (int, int) {
	g1, g2 := grid{}, grid{}
	p1, p2 := grid{}, grid{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		ints := utils.GetInts(line)
		x1, y1 := ints[0], ints[1]
		x2, y2 := ints[2], ints[3]
		dX, dY := x2-x1, y2-y1

		count := max(abs(dX), abs(dY))
		dX /= count
		dY /= count
		incG1 := dX == 0 || dY == 0

		for i, x, y := 0, x1, y1; i <= count; i, x, y = i+1, x+dX, y+dY {
			pos := pointHash((x&0xffff)<<16 | (y & 0xffff))
			if incG1 {
				if g1[pos] {
					p1[pos] = true
				} else {
					g1[pos] = true
				}
			}
			if g2[pos] {
				p2[pos] = true
			} else {
				g2[pos] = true
			}
		}
	}

	return len(p1), len(p2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type grid map[pointHash]bool
type pointHash uint32

var benchmark = false
