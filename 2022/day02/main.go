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
		opSym := symbol(input[pos] - 'A')
		mySym := symbol(input[pos+2] - 'X')
		p1score += p1lookup[inState(opSym, mySym)]
		p2score += p2lookup[inState(opSym, mySym)]
	}
	return p1score, p2score
}

func calcLookups() (scoreLookup, scoreLookup) {
	p1 := make(scoreLookup, 1<<4)
	p2 := make(scoreLookup, 1<<4)

	for op := symbolRock; op <= symbolScissors; op++ {
		for me := symbolRock; me <= symbolScissors; me++ {
			winScore, choiceScore := scoreRound(op, me)
			p1[inState(op, me)] = winScore + choiceScore
			p2[inState(op, symbol(winScore)/3)] = int(me) + 1 + winScore
		}
	}
	return p1, p2
}

type scoreLookup []int

func inState(op symbol, me symbol) byte { return byte(op)<<2 + byte(me) }

func scoreRound(op, me symbol) (int, int) { return 3 * int((4+me-op)%3), int(me) + 1 }

type symbol int

const (
	symbolRock symbol = iota
	symbolPaper
	symbolScissors
)

var benchmark = false
