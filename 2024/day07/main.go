package main

import (
	_ "embed"
	"fmt"
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
	vals := make([]int, 0, 15)
	acc := 0
	for i := 0; i < len(input); i++ {
		ch := input[i]
		switch {
		case ch == '\n':
			vals = append(vals, acc)
			canP1, canP2 := canMake(vals[0], vals[1:]...)
			if canP1 {
				p1 += vals[0]
			}
			if canP2 {
				p2 += vals[0]
			}
			vals = vals[0:0]
			acc = 0
		case ch == ':':
			i++
			fallthrough
		case ch == ' ':
			vals = append(vals, acc)
			acc = 0
		default:
			acc *= 10
			acc += int(ch - '0')
		}
	}
	return p1, p2
}

const searchBuffer = 60_000

var open = make([]search, 0, searchBuffer)
var next = make([]search, 0, searchBuffer)

type search struct {
	pos    int
	total  int
	concat bool
}

func canMake(target int, operands ...int) (bool, bool) {
	validP2 := false
	open = open[0:0]
	next = next[0:0]
	open = append(open, search{1, operands[0], false})
	for len(open) > 0 {
		for _, cur := range open {
			mult := cur.total * operands[cur.pos]
			add := cur.total + operands[cur.pos]
			concat := concatInts(cur.total, operands[cur.pos])
			if cur.pos < len(operands)-1 {
				if mult <= target {
					next = append(next, search{cur.pos + 1, mult, cur.concat})
				}
				if add <= target {
					next = append(next, search{cur.pos + 1, add, cur.concat})
				}
				if concat < target {
					next = append(next, search{cur.pos + 1, concat, true})
				}
				continue
			}
			if mult == target || add == target {

				if cur.concat {
					validP2 = true
					continue
				}
				return true, true
			}
			if concat == target {
				validP2 = true
			}
		}
		open, next = next, open[0:0]
	}
	return false, validP2
}

func concatInts(a, b int) int {
	bOff := 10
	a *= 10
	for ; bOff <= b; bOff *= 10 {
		a *= 10
	}
	return a + b
}

var benchmark = false
