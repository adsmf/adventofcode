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
	p1, p2 := 0, 0
	utils.EachLine(input, func(_ int, line string) (done bool) {
		vals := utils.GetInts(line)
		canP1, canP2 := canMake(vals[0], vals[1:]...)
		if canP1 {
			p1 += vals[0]
		}
		if canP2 {
			p2 += vals[0]
		}
		return false
	})
	return p1, p2
}

func canMake(target int, operands ...int) (bool, bool) {
	type search struct {
		pos    int
		total  int
		concat bool
	}
	validP2 := false
	open, next := []search{{1, operands[0], false}}, []search{}
	for len(open) > 0 {
		for _, cur := range open {
			if cur.pos == len(operands) {
				if cur.total == target {
					if cur.concat {
						validP2 = true
						continue
					}
					return true, true
				}
				continue
			}
			if cur.total > target {
				continue
			}
			next = append(next, search{cur.pos + 1, cur.total * operands[cur.pos], cur.concat})
			next = append(next, search{cur.pos + 1, cur.total + operands[cur.pos], cur.concat})
			concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", cur.total, operands[cur.pos]))
			next = append(next, search{cur.pos + 1, concat, true})
		}
		open, next = next, open[0:0]
	}
	return false, validP2
}

var benchmark = false
