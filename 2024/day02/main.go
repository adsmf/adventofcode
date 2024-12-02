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
	utils.EachLine(input, func(index int, line string) (done bool) {
		readings := utils.GetInts(line)
		invalid := 0
		validP1, validP2 := true, true
		stepLow, stepHigh := 1, 3
		if readings[len(readings)-1] < readings[0] {
			stepLow, stepHigh = -3, -1
		}
		for i := 1; i < len(readings); i++ {
			low, high := readings[i-1]+stepLow, readings[i-1]+stepHigh
			if readings[i] < low ||
				readings[i] > high {
				validP1 = false
				invalid++
				if invalid > 1 {
					validP2 = false
					break
				}
				if i == len(readings)-1 {
					break
				}
				i++
				if i == 2 && (readings[i] >= readings[i-1]+stepLow &&
					readings[i] <= readings[i-1]+stepHigh) {
					continue
				}
				if readings[i] < low ||
					readings[i] > high {
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
