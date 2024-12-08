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
	w := 0
	for ; input[w] != '\n'; w++ {
	}
	h := len(input) / (w + 1)
	nodes := make(map[byte][]point, 100)
	x, y := 0, 0
	for pos := 0; pos < len(input); pos++ {
		switch ch := input[pos]; ch {
		case '\n':
			x = 0
			y++
		case '.':
			x++
		default:
			nodes[ch] = append(nodes[ch], point{x, y})
			x++
		}
	}
	valid := func(p point) bool {
		return p.x >= 0 && p.x < w && p.y >= 0 && p.y < h
	}

	antinodesP1 := make([]bool, w*h)
	antinodesP2 := make([]bool, w*h)
	for _, freqNodes := range nodes {
		for i := 0; i < len(freqNodes)-1; i++ {
			n1 := freqNodes[i]
			antinodesP2[n1.pos(w)] = true
			for j := i + 1; j < len(freqNodes); j++ {
				n2 := freqNodes[j]
				diff := n1.sub(n2)
				antinodesP2[n2.pos(w)] = true
				for first, an := true, n1.add(diff); valid(an); an = an.add(diff) {
					if first {
						antinodesP1[an.pos(w)] = true
						first = false
					}
					antinodesP2[an.pos(w)] = true
				}
				for first, an := true, n2.sub(diff); valid(an); an = an.sub(diff) {
					if first {
						antinodesP1[an.pos(w)] = true
						first = false
					}
					antinodesP2[an.pos(w)] = true
				}
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
