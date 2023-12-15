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
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	totalHash := 0
	boxes := [256][]lens{}
	utils.EachSection(input, ',', func(index int, instruction string) (done bool) {
		totalHash += hashString(instruction)
		label := ""
		oper := '?'
		value := 0
		for pos, ch := range instruction {
			switch ch {
			case '-', '=':
				oper = ch
				label = instruction[0:pos]
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				value = value*10 + int(ch-'0')
			}
		}
		box := hashString(string(label))
		if boxes[box] == nil {
			boxes[box] = make([]lens, 0, 2)
		}
		switch oper {
		case '-':
			n := 0
			for i := 0; i < len(boxes[box]); i++ {
				if boxes[box][i].label == string(label) {
					continue
				}
				boxes[box][n] = boxes[box][i]
				n++
			}
			boxes[box] = boxes[box][:n]
		case '=':
			found := false
			for i := 0; i < len(boxes[box]); i++ {
				if boxes[box][i].label == string(label) {
					found = true
					boxes[box][i].value = value
					break
				}
			}
			if !found {
				boxes[box] = append(boxes[box], lens{string(label), value})
			}
		}
		return false
	})
	focusPower := 0
	for box, lenses := range boxes {
		for slot, lens := range lenses {
			focusPower += (box + 1) * (slot + 1) * (lens.value)
		}
	}
	return totalHash, focusPower
}

type lens struct {
	label string
	value int
}

func hashString(in string) int {
	hash := 0
	for _, ch := range in {
		if ch == '\n' {
			continue
		}
		hash += int(ch)
		hash *= 17
		hash &= 0xff
	}
	return hash
}

var benchmark = false
