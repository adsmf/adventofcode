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
	// p2 = part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1, p2 := 0, 0
	readings := make([]int, 0, 10)
	utils.EachLine(input, func(index int, line string) (done bool) {
		readings = readings[0:0]
		utils.EachInteger(line, func(index, value int) (done bool) {
			readings = append(readings, value)
			return false
		})
		invalidSkipped := false
		validP1, validP2 := true, true
		stepLow, stepHigh := 1, 3
		if readings[len(readings)-1] < readings[0] {
			stepLow, stepHigh = -3, -1
		}
		safeStep := func(i, j int) bool {
			return readings[j] >= readings[i]+stepLow && readings[j] <= readings[i]+stepHigh
		}
		for i := 1; i < len(readings); i++ {
			if !safeStep(i-1, i) {
				validP1 = false
				if invalidSkipped {
					validP2 = false
					break
				}
				invalidSkipped = true
				if i == len(readings)-1 {
					break
				}
				i++
				if i == 2 && safeStep(i-1, i) {
					continue
				}
				if !safeStep(i-2, i) {
					validP2 = false
					break
				}
			}
		}
		if validP1 {
			p1++
		}
		if validP2 {
			p2++
		}
		return false
	})
	return p1, p2
}

var benchmark = false
