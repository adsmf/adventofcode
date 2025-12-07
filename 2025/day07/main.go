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
	width := 0
	for i, ch := range input {
		if ch == '\n' {
			width = i
			break
		}
	}
	beams, next := make([]int, 150), make([]int, 150)
	beams, next = beams[:width], next[:width]
	splits := 0
	x := 0
	for _, ch := range input {
		switch ch {
		case '\n':
			x = 0
			clear(beams)
			next, beams = beams, next
		case '^':
			next[x-1] += beams[x]
			next[x+1] += beams[x]
			if beams[x] > 0 {
				splits++
			}
			x++
		case '.':
			next[x] += beams[x]
			x++
		case 'S':
			next[x] = 1
			x++
		}
	}
	count := 0
	for _, beam := range beams {
		count += beam
	}
	return splits, count
}

var benchmark = false
