package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	p1score, p2score := 0, 0
	p1lookup, p2lookup := calcLookups()
	for pos := 0; pos < len(input); pos += 4 {
		state := stateRep(symbol(input[pos]-'A'), symbol(input[pos+2]-'X'))
		p1score += p1lookup[state]
		p2score += p2lookup[state]
	}
	return p1score, p2score
}

func calcLookups() (scoreLookup, scoreLookup) {
	p1 := make(scoreLookup, 1<<4)
	p2 := make(scoreLookup, 1<<4)
	for op := symbolRock; op <= symbolScissors; op++ {
		for me := symbolRock; me <= symbolScissors; me++ {
			winstate := (4 + me - op) % 3
			choiceScore := me + 1
			p1[stateRep(op, me)] = winstate*3 + choiceScore
			p2[stateRep(op, winstate)] = me + 1 + winstate*3
		}
	}
	return p1, p2
}

type scoreLookup []int

func stateRep(op symbol, me symbol) byte {
	return byte(op)<<2 + byte(me)
}

type symbol = int

const (
	symbolRock symbol = iota
	symbolPaper
	symbolScissors
)

var benchmark = false
