package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
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
	gm := load()
	usedRows := map[int][]*galaxy{}
	usedCols := map[int][]*galaxy{}
	for _, gal := range gm.galaxies {
		usedRows[gal.y] = append(usedRows[gal.y], gal)
		usedCols[gal.x] = append(usedCols[gal.x], gal)
	}
	for space, y := 0, 0; y <= gm.max.y; y++ {
		if len(usedRows[y]) == 0 {
			space++
			continue
		}
		for _, gal := range usedRows[y] {
			gal.eY = space
		}
	}
	for space, x := 0, 0; x <= gm.max.x; x++ {
		if len(usedCols[x]) == 0 {
			space++
			continue
		}
		for _, gal := range usedCols[x] {
			gal.eX = space
		}
	}

	p1dist, p2dist := 0, 0
	for i := 0; i < len(gm.galaxies)-1; i++ {
		for j := i + 1; j < len(gm.galaxies); j++ {
			p1dist += gm.galaxies[i].manhattan(*gm.galaxies[j], 1)
			p2dist += gm.galaxies[i].manhattan(*gm.galaxies[j], 1000000-1)
		}
	}

	return p1dist, p2dist
}

func load() galacticMap {
	gm := galacticMap{}
	min, max := galaxy{x: utils.MaxInt, y: utils.MaxInt}, galaxy{x: utils.MinInt, y: utils.MinInt}
	utils.EachLine(input, func(y int, line string) (done bool) {
		for x, ch := range line {
			if ch != '#' {
				continue
			}
			pos := galaxy{x: x, y: y}
			min = min.min(pos)
			max = max.max(pos)
			gm.galaxies = append(gm.galaxies, &pos)
		}
		return false
	})
	gm.min = min
	gm.max = max
	return gm
}

type galacticMap struct {
	galaxies []*galaxy
	min, max galaxy
}

type galaxy struct {
	x, y   int
	eX, eY int
}

func (p galaxy) manhattan(oth galaxy, expFactor int) int {
	xDist := p.x + p.eX*expFactor - (oth.x + oth.eX*expFactor)
	if xDist < 0 {
		xDist *= -1
	}
	yDist := p.y + p.eY*expFactor - (oth.y + oth.eY*expFactor)
	if yDist < 0 {
		yDist *= -1
	}
	return xDist + yDist
}

func (p galaxy) min(oth galaxy) galaxy {
	if oth.x < p.x {
		p.x = oth.x
	}
	if oth.y < p.y {
		p.y = oth.y
	}
	return p
}
func (p galaxy) max(oth galaxy) galaxy {
	if oth.x > p.x {
		p.x = oth.x
	}
	if oth.y > p.y {
		p.y = oth.y
	}
	return p
}

var benchmark = false
