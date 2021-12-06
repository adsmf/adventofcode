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
	p1, p2 := loadInputArray()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

// Solution using array rather than map or slice
func loadInputArray() (int, int) {
	g1, g2 := arrayGrid{}, arrayGrid{}
	p1, p2 := 0, 0
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
			pos := x*1000 + y
			if incG1 {
				switch g1[pos] {
				case 0:
					g1[pos] = 1
				case 1:
					g1[pos] = 2
					p1++
				}
			}
			switch g2[pos] {
			case 0:
				g2[pos] = 1
			case 1:
				g2[pos] = 2
				p2++
			}
		}
	}

	return p1, p2
}

// Solution using slice rather than map
func loadInputSlice() (int, int) {
	g1, g2 := make(sliceGrid, 1000*1000), make(sliceGrid, 1000*1000)
	p1, p2 := 0, 0
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
			pos := x*1000 + y
			if incG1 {
				switch g1[pos] {
				case 0:
					g1[pos] = 1
				case 1:
					g1[pos] = 2
					p1++
				}
			}
			switch g2[pos] {
			case 0:
				g2[pos] = 1
			case 1:
				g2[pos] = 2
				p2++
			}
		}
	}

	return p1, p2
}

// Initial solution using map
func loadInputMap() (int, int) {
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

type sliceGrid []uint8
type arrayGrid [1000 * 1000]uint8

var benchmark = false
