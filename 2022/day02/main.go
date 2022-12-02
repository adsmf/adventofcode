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
	for pos := 0; pos < len(input); pos += 4 {
		opSym := symbol(input[pos] - 'A')
		mySym := symbol(input[pos+2] - 'X')

		winScore, choiceScore := scoreRound(opSym, mySym)
		p1score += winScore + choiceScore

		p2target := int(mySym * 3)
		if winScore == p2target {
			p2score += winScore + choiceScore
		} else {
			for choice := symbolRock; choice <= symbolScissors; choice++ {
				winScore, choiceScore := scoreRound(opSym, choice)
				if p2target == winScore {
					p2score += winScore + choiceScore
					break
				}
			}
		}
	}
	return p1score, p2score
}

func scoreRound(op, me symbol) (int, int) {
	choiceScore := int(me) + 1
	switch {
	case op == me:
		return 3, choiceScore
	case (op+1)%3 == me:
		return 6, choiceScore
	default:
		return 0, choiceScore
	}
}

type symbol int

const (
	symbolRock symbol = iota
	symbolPaper
	symbolScissors
)

var benchmark = false
