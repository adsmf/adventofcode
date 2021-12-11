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
				g[toPoint(x, y)] = int(ch - '0')
			}
		}
	}
	return g
}

type grid []int

func (g grid) step() int {
	toFlash := make([]point, 0, 100)
	hasFlashed := make([]bool, 100)
	for p := range g {
		g[p]++
		if g[p] > 9 {
			toFlash = append(toFlash, point(p))
			hasFlashed[p] = true
		}
	}
	neighbors := make([]point, 0, 8)
	for len(toFlash) > 0 {
		p := toFlash[len(toFlash)-1]
		toFlash = toFlash[:len(toFlash)-1]
		neighbors = p.neighbours(neighbors)
		for _, n := range neighbors {
			g[n]++
			if !hasFlashed[n] && g[n] > 9 {
				hasFlashed[n] = true
				toFlash = append(toFlash, n)
			}
		}
	}
	numFlashed := 0
	for p, flashed := range hasFlashed {
		if flashed {
			g[p] = 0
			numFlashed++
		}
	}

	return numFlashed
}

type point int

func (p point) neighbours(neighbours []point) []point {
	pX, pY := int(p)%10, int(p)/10
	numNeighbours := 0
	neighbours = neighbours[:cap(neighbours)]
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			nX, nY := pX+x, pY+y
			if nX < 0 || nX > 9 || nY < 0 || nY > 9 {
				continue
			}
			neighbours[numNeighbours] = point(nX + 10*nY)
			numNeighbours++
		}
	}
	neighbours = neighbours[:numNeighbours]
	return neighbours
}
func toPoint(x, y int) point {
	return point(x + 10*y)
}

var benchmark = false
