package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve(input)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(input string) (int, int) {
	im, enahncement := load(input)
	for i := 0; i < 2; i++ {
		im.enhance(enahncement)
	}
	p1 := im.lit
	for i := 2; i < 50; i++ {
		im.enhance(enahncement)
	}
	return p1, im.lit
}

func load(input string) (image, []bool) {
	blocks := strings.Split(input, "\n\n")

	enhancement := make([]bool, 0, len(blocks[0]))
	for _, ch := range blocks[0] {
		enhancement = append(enhancement, ch == '#')
	}

	lines := strings.Split(blocks[1], "\n")
	g := make(grid, 1<<(bitsize*4))
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				g[makePoint(x, y)] = true
			}
		}
	}
	im := image{
		grid:     g,
		buf:      make(grid, len(g)),
		inverted: false,
		min:      makePoint(0, 0),
		max:      makePoint(len(lines[0]), len(lines)),
	}
	return im, enhancement
}

type image struct {
	grid     grid
	buf      grid
	inverted bool
	min, max point
	lit      int
}

type grid []bool

func (im *image) enhance(enhancement []bool) {
	nextInverted := im.inverted
	if enhancement[0] {
		nextInverted = !im.inverted
	}

	lit := 0
	for y := im.min.y() - 1; y <= im.max.y()+1; y++ {
		for x := im.min.x() - 1; x <= im.max.x()+1; x++ {
			pos := makePoint(x, y)
			lookup := 0
			if im.inverted {
				for _, n := range pos.neighbours() {
					lookup <<= 1
					if !im.grid[n] {
						lookup += 1
					}
				}
			} else {
				for _, n := range pos.neighbours() {
					lookup <<= 1
					if im.grid[n] {
						lookup += 1
					}
				}
			}
			value := enhancement[lookup]
			if nextInverted {
				value = !value
			}
			if value {
				lit++
			}
			im.buf[pos] = value
		}
	}

	im.grid, im.buf = im.buf, im.grid
	im.inverted = nextInverted
	im.lit = lit
	im.min = im.min.dec()
	im.max = im.max.inc()
}

const (
	bitsize = 6
	offset  = 1 << bitsize
	mask    = (1 << (bitsize * 2)) - 1
)

func makePoint(x, y int) point { return point(x+offset)<<(bitsize*2) + point(y+offset) }

type point uint32

func (p point) x() int     { return int(p>>(bitsize*2)) - offset }
func (p point) y() int     { return int(p&mask) - offset }
func (p point) dec() point { return p - 1<<(bitsize*2) - 1 }
func (p point) inc() point { return p + 1<<(bitsize*2) + 1 }

func (p point) neighbours() []point {
	return []point{
		p - 1<<(bitsize*2) - 1,
		p - 1,
		p + 1<<(bitsize*2) - 1,

		p - 1<<(bitsize*2),
		p,
		p + 1<<(bitsize*2),

		p - 1<<(bitsize*2) + 1,
		p + 1,
		p + 1<<(bitsize*2) + 1,
	}
}

var benchmark = false
