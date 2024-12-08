package main

import (
	_ "embed"
	"fmt"
	"slices"
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

const gridCap = 50 * 50

func solve() (int, int) {
	w := 0
	for ; input[w] != '\n'; w++ {
	}
	h := len(input) / (w + 1)
	if gridCap < w*h {
		panic("grid capacity too low")
	}
	type nodeFreq struct {
		pos  point
		freq byte
	}
	nodes := make([]nodeFreq, 0, 250)
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch ch := input[pos]; ch {
		case '\n':
			x = 0
			y++
		case '.':
			x++
		default:
			nodes = append(nodes, nodeFreq{point{x, y}, ch})
			x++
		}
	}
	valid := func(p point) bool {
		return p.x >= 0 && p.x < w && p.y >= 0 && p.y < h
	}
	slices.SortFunc(nodes, func(a, b nodeFreq) int {
		return int(a.freq) - int(b.freq)
	})
	antinodesP1 := make([]bool, gridCap)
	antinodesP2 := make([]bool, gridCap)
	for i := 0; i < len(nodes)-1; i++ {
		n1 := nodes[i]
		for j := i + 1; j < len(nodes); j++ {
			n2 := nodes[j]
			if n2.freq != n1.freq {
				break
			}
			diff := n1.pos.sub(n2.pos)
			antinodesP2[n1.pos.pos(w)] = true
			antinodesP2[n2.pos.pos(w)] = true
			for first, an := true, n1.pos.add(diff); valid(an); an = an.add(diff) {
				if first {
					antinodesP1[an.pos(w)] = true
					first = false
				}
				antinodesP2[an.pos(w)] = true
			}
			for first, an := true, n2.pos.sub(diff); valid(an); an = an.sub(diff) {
				if first {
					antinodesP1[an.pos(w)] = true
					first = false
				}
				antinodesP2[an.pos(w)] = true
			}
		}
	}
	p1, p2 := 0, 0
	for _, an := range antinodesP1 {
		if an {
			p1++
		}
	}
	for _, an := range antinodesP2 {
		if an {
			p2++
		}
	}
	return p1, p2
}

type point struct{ x, y int }

func (p point) pos(w int) int     { return p.x + w*p.y }
func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }
func (p point) sub(o point) point { return point{p.x - o.x, p.y - o.y} }

var benchmark = false
