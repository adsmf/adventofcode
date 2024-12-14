package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

const (
	width   = 101
	height  = 103
	gridMax = width * height
)

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	vals := [4]int{}
	robots := make(robotList, 0, 500)
	utils.EachInteger(input, func(index, value int) (done bool) {
		idx := index % 4
		vals[idx] = value
		if idx < 3 {
			return false
		}
		r := robot{point{vals[0], vals[1]}, point{vals[2], vals[3]}}
		robots = append(robots, r)
		return false
	})
	move(robots, 100)
	p1 := safetyFactor(robots)
	for i := 101; ; i++ {
		move(robots, 1)
		if unique(robots) {
			return p1, i
		}
	}
}

func unique(robots robotList) bool {
	seen := [gridMax]bool{}
	for _, r := range robots {
		pos := r.p.x + r.p.y*width
		if seen[pos] {
			return false
		}
		seen[pos] = true
	}
	return true
}

func safetyFactor(robots robotList) int {
	midH, midW := height/2, width/2
	quads := [4]int{}
	for _, rob := range robots {
		if rob.p.x == midW || rob.p.y == midH {
			continue
		}
		quad := 0
		if rob.p.x < midW {
			quad++
		}
		if rob.p.y > midH {
			quad += 2
		}
		quads[quad]++
	}
	return quads[0] * quads[1] * quads[2] * quads[3]
}

func move(robots robotList, steps int) {
	for i, r := range robots {
		rX := mod(r.p.x+r.v.x*steps, width)
		rY := mod(r.p.y+r.v.y*steps, height)
		robots[i].p = point{rX, rY}
	}
}

func mod(a, b int) int { return (a%b + b) % b }

type robotList []robot
type robot struct {
	p point
	v point
}

type point struct{ x, y int }

var benchmark = false
