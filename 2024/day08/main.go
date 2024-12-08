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
	nodes := map[byte][]point{}
	x, y := 0, 0
	w := 0
	for pos := 0; pos < len(input); pos++ {
		switch ch := input[pos]; ch {
		case '\n':
			w = x
			x = 0
			y++
		case '.':
			x++
		default:
			nodes[ch] = append(nodes[ch], point{x, y})
			x++
		}
	}
	h := y
	valid := func(p point) bool {
		return p.x >= 0 && p.x < w && p.y >= 0 && p.y < h
	}

	antinodesP1 := map[point]bool{}
	antinodesP2 := map[point]bool{}
	for _, freqNodes := range nodes {
		for i := 0; i < len(freqNodes)-1; i++ {
			n1 := freqNodes[i]
			antinodesP2[n1] = true
			for j := i + 1; j < len(freqNodes); j++ {
				n2 := freqNodes[j]
				diff := n1.sub(n2)
				antinodesP2[n2] = true
				for first, an := true, n1.add(diff); valid(an); an = an.add(diff) {
					if first {
						antinodesP1[an] = true
						first = false
					}
					antinodesP2[an] = true
				}
				for first, an := true, n2.sub(diff); valid(an); an = an.sub(diff) {
					if first {
						antinodesP1[an] = true
						first = false
					}
					antinodesP2[an] = true
				}
			}
		}
	}
	return len(antinodesP1), len(antinodesP2)
}

type point struct{ x, y int }

func (p point) add(o point) point { return point{p.x + o.x, p.y + o.y} }
func (p point) sub(o point) point { return point{p.x - o.x, p.y - o.y} }

var benchmark = false
