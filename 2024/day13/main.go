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
	p1, p2 := 0, 0
	offset := point{10_000_000_000_000, 10_000_000_000_000}
	vals := [6]int{}
	pos := 0
	utils.EachInteger(input, func(_, value int) (done bool) {
		vals[pos] = value
		if pos < 5 {
			pos++
			return true
		}
		pos = 0
		a := point{vals[0], vals[1]}
		b := point{vals[2], vals[3]}
		p := point{vals[4], vals[5]}
		p1 += calcTokens(a, b, p)
		p2 += calcTokens(a, b, p.add(offset))
		return true
	})
	return p1, p2
}

func calcTokens(a, b, p point) int {
	d := a.y*b.x - a.x*b.y
	tA := p.det(b) / d
	tB := a.det(p) / d
	if a.mul(tA).add(b.mul(tB)) == p {
		return 3*tA + tB
	}
	return 0
}

type point struct{ x, y int }

func (p point) add(o point) point    { return point{p.x + o.x, p.y + o.y} }
func (p point) det(o point) int      { return p.y*o.x - p.x*o.y }
func (p point) mul(scalar int) point { return point{p.x * scalar, p.y * scalar} }

var benchmark = false
