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
	startCol := 0
	for i, ch := range input {
		if ch == '\n' {
			width = i
			break
		}
		if ch == 'S' {
			startCol = i
		}
	}
	lines := len(input) / width
	beams, next := make([]int, 150), make([]int, 150)
	beams, next = beams[:width], next[:width]
	beams[startCol] = 1
	splits := 0
	for row := 2; row < lines-1; row += 2 {
		for x := 0; x < width; x++ {
			switch input[row*(width+1)+x] {
			case '^':
				if beams[x] > 0 {
					next[x-1] += beams[x]
					next[x+1] += beams[x]
					splits++
				}
			case '.':
				next[x] += beams[x]
			}
		}
		clear(beams)
		next, beams = beams, next
	}
	count := 0
	for _, beam := range beams {
		count += beam
	}
	return splits, count
}

var benchmark = false
