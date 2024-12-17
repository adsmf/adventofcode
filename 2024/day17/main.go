package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

var p1buf = [50]byte{}

func main() {
	p2 := solve()
	i := 0
	for ; p1buf[i] > 0; i++ {
	}
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1buf[:i])
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() int {
	regA := 0
	inst := make([]int, 0, 16)
	utils.EachInteger(input, func(idx, value int) (done bool) {
		if idx == 0 {
			regA = value
			return
		}
		if idx < 3 {
			return
		}
		inst = append(inst, value)
		return
	})
	part1(regA)
	return part2(inst)
}

func part1(regA int) {
	first := true
	i := 0
	for {
		if regA == 0 {
			break
		}
		if !first {
			p1buf[i] = ','
			i++
		}
		first = false
		val := nextOut(regA)
		p1buf[i] = '0' + byte(val)
		i++
		regA >>= 3
	}
}

func part2(inst []int) int {
	open := make([]int, 0, 25)
	next := make([]int, 0, 25)
	open = append(open, 0)
	var p2 int
	for i := len(inst) - 1; i >= 0; i-- {
		p2 = utils.MaxInt
		val := inst[i]
		for _, cur := range open {
			for i := range 8 {
				try := cur<<3 + int(i)
				if nextOut((try)) == (val) {
					next = append(next, int(try))
					p2 = min(p2, try)
				}
			}
		}
		open, next = next, open[0:0]
	}
	return p2
}

func nextOut(A int) int {
	/*
		B = A & 7
		B ^= 7
		C = A >> B
		B ^= 7
		B ^= C
		A >>= 3
		Out B & 7
		JNZ 0
	*/
	return ((A & 7) ^ (A >> ((A & 7) ^ 7))) & 7
}

var benchmark = false
