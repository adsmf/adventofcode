package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
	"github.com/adsmf/adventofcode/utils/solvers"
)

//go:embed input.txt
var input string

const (
	width   = 101
	height  = 103
	gridMax = width * height
)

func main() {
	p1, p2 := solveAlt()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solveAlt() (int, int) {
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
	p1 := 0
	h1, v1 := 0, 0
	for i := 1; ; i++ {
		move(robots, 1)
		quads := calcQuads(robots)
		if i == 100 {
			p1 = safetyFactor(quads)
		}
		if h1 == 0 && hSig(quads) {
			h1 = i
		}
		if v1 == 0 && vSig(quads) {
			v1 = i
		}
		if h1 > 0 && v1 > 0 {
			if i < 100 {
				move(robots, 100-i)
				p1 = safetyFactor(calcQuads(robots))
			}
			break
		}
	}
	res, err := solvers.ChineseRemainderTheoremStd([]int{h1, v1}, []int{width, height})
	if err != nil {
		panic(err)
	}
	return p1, int(res)
}

func hSig(quads [4]int) bool {
	return quads[0]+quads[2]-quads[1]-quads[3] > 100
}
func vSig(quads [4]int) bool {
	return quads[2]+quads[3]-quads[0]-quads[1] > 100
}

func calcQuads(robots robotList) [4]int {
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
	return quads
}

func safetyFactor(quads [4]int) int {
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
