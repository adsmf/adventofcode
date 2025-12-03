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
	total := 0

	utils.EachLine(input, func(index int, line string) (done bool) {
		bestH, bestL := 0, 0
		for i, j := range line {
			jolt := int(j - '0')
			if jolt > bestH && i < len(line)-1 {
				bestH = jolt
				bestL = -1
				continue
			}
			bestL = max(bestL, jolt)
		}
		total += bestH*10 + bestL
		return
	})
	return total
}

func part2() int {
	total := 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		best := [12]int{}
		for i, j := range line {
			jolt := int(j - '0')
			for bat := range len(best) {
				if jolt > best[bat] && (len(best)-bat) < (len(line)-i+1) {
					best[bat] = jolt
					for reset := bat + 1; reset < len(best); reset++ {
						best[reset] = -1
					}
					break
				}
			}
		}
		jolts := 0
		for i := range len(best) {
			jolts = jolts*10 + best[i]
		}
		total += jolts
		return
	})
	return total
}

var benchmark = false
