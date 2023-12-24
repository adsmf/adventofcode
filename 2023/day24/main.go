package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	testMin, testMax := 200000000000000.0, 400000000000000.0
	hs := []hailstoneInfo{}
	utils.EachLine(input, func(index int, line string) (done bool) {
		v := utils.GetInts(line)
		hs = append(hs, hailstoneInfo{
			point{v[0], v[1], v[2]},
			point{v[3], v[4], v[5]},
		})
		return false
	})
	count := 0
	for i := 0; i < len(hs)-1; i++ {
		for j := i + 1; j < len(hs); j++ {
			h1, h2 := hs[i], hs[j]
			cX, cY := h1.crossing2D(h2)
			f1 := (cX - float64(h1.pos.x)) / float64(h1.vel.x)
			f2 := (cX - float64(h2.pos.x)) / float64(h2.vel.x)
			if (f1 > 0 && f2 > 0) &&
				cX >= testMin && cX <= testMax &&
				cY >= testMin && cY <= testMax {
				count++
			}
		}
	}
	return count
}

func part2() int {
	return -1
}

type hailstoneInfo struct {
	pos point
	vel point
}

func (h hailstoneInfo) posAt(t int) point {
	return point{
		h.pos.x + t*h.vel.x,
		h.pos.y + t*h.vel.y,
		h.pos.z + t*h.vel.z,
	}
}

func (h hailstoneInfo) crossing2D(o hailstoneInfo) (float64, float64) {
	m1, c1 := h.equation2D()
	m2, c2 := o.equation2D()
	crossX := (c1 - c2) / (m2 - m1)
	crossY := (c1*m2 - c2*m1) / (m2 - m1)
	return crossX, crossY
}

func (h hailstoneInfo) equation2D() (float64, float64) {
	m := float64(h.vel.y) / float64(h.vel.x)
	i := float64(h.pos.x) / float64(-h.vel.x)
	c := float64(h.vel.y)*i + float64(h.pos.y)
	return m, c
}

type point struct{ x, y, z int }

func (p point) add(q point) point      { return point{p.x + q.x, p.y + q.y, p.z + q.z} }
func (p point) sub(q point) point      { return point{p.x - q.x, p.y - q.y, p.z - q.z} }
func (p point) times(scalar int) point { return point{scalar * p.x, scalar * p.y, scalar * p.z} }
func (p point) manhattan() int {
	return utils.IntAbs(p.x) + utils.IntAbs(p.y) + utils.IntAbs(p.z)
}

var benchmark = false
