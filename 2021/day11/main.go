package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	g := load(input)
	p1, p2 := runSim(g)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func runSim(g grid) (int, int) {
	total := 0
	for i := 0; i < 100; i++ {
		flashes := g.step()
		total += flashes
	}
	for i := 101; ; i++ {
		flashes := g.step()
		if flashes == 100 {
			return total, i
		}
	}
}

func load(in string) grid {
	g := make(grid, 100)
	for y, line := range utils.GetLines(in) {
		for x, ch := range []byte(line) {
			if ch >= '0' {
				g[point{x, y}] = int(ch - '0')
			}
		}
	}
	return g
}

type grid map[point]int

func (g grid) step() int {
	toFlash := []point{}
	flashed := map[point]bool{}
	for p := range g {
		g[p]++
		if g[p] > 9 {
			toFlash = append(toFlash, p)
		}
	}
	for len(toFlash) > 0 {
		p := toFlash[len(toFlash)-1]
		toFlash = toFlash[:len(toFlash)-1]
		if !flashed[p] {
			flashed[p] = true
			for _, n := range p.neighbours() {
				if _, valid := g[n]; valid {
					g[n]++
					if g[n] > 9 {
						toFlash = append(toFlash, n)
					}
				}
			}
		}
	}
	for p := range flashed {
		g[p] = 0
	}
	return len(flashed)
}

type point struct{ x, y int }

func (p point) neighbours() []point {
	n := make([]point, 0, 8)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			n = append(n, point{p.x + x, p.y + y})
		}
	}
	return n
}

var benchmark = false
