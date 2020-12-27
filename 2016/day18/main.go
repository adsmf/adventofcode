package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	traps := load("input.txt")
	p1 := countSafe(traps, 40)
	p2 := countSafe(traps, 400000)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func countSafe(traps rowTraps, rows int) int {
	count := 0
	for i := 0; i < rows; i++ {
		count += traps.countSafe()
		traps = traps.next()
	}
	return count
}

type rowTraps []bool

func (rt rowTraps) countSafe() int {
	count := 0
	for _, trap := range rt {
		if !trap {
			count++
		}
	}
	return count
}

func (rt rowTraps) next() rowTraps {
	next := make(rowTraps, len(rt))
	for i := 0; i < len(rt); i++ {
		l, r := false, false
		if i > 0 {
			l = rt[i-1]
		}
		if i+1 < len(rt) {
			r = rt[i+1]
		}
		next[i] = (l && !r) || (r && !l)
	}
	return next
}

func load(filename string) rowTraps {
	inputBytes, _ := ioutil.ReadFile(filename)
	input := strings.TrimSpace(string(inputBytes))
	traps := make(rowTraps, len(input))
	for idx, char := range input {
		if char == '^' {
			traps[idx] = true
		}
	}
	return traps
}

var benchmark = false
