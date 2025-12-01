package main

import (
	_ "embed"
	"fmt"
	"strconv"

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
	const dialSize = 100
	dialVal := 50
	p1, p2 := 0, 0
	lastCW := true
	utils.EachLine(input, func(index int, line string) (done bool) {
		curCW := line[0] == 'R'
		val, _ := strconv.Atoi(line[1:])
		rotations := val / dialSize
		p2 += rotations
		val -= dialSize * rotations
		if curCW != lastCW {
			if dialVal > 0 {
				dialVal = dialSize - dialVal
			}
			lastCW = curCW
		}
		dialVal += val
		if dialVal >= dialSize {
			dialVal -= dialSize
			p2++
		}
		if dialVal == 0 {
			p1++
		}
		return false
	})
	return p1, p2
}

var benchmark = false
