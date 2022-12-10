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
		fmt.Printf("Part 2:\n%s\n", p2)
	}
}
func solve() (int, screenBuffer) {
	regX, signal, cycle := 1, 0, 0
	countCycle := 20
	screen := screenBuffer{40: '\n', 81: '\n', 122: '\n', 163: '\n', 204: '\n', 245: '\n'}
	screenPos, hPos := 0, 0
	execCycle := func() {
		cycle++
		if cycle == countCycle {
			signal += cycle * regX
			countCycle += 40
		}
		hPos++
		if hPos-regX >= 0 && hPos-regX <= 2 {
			screen[screenPos] = '#'
		} else {
			screen[screenPos] = '.'
		}
		screenPos++
		if hPos == 40 {
			hPos = 0
			screenPos++
		}
	}
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case 'n':
			i += 4
			execCycle()
		case 'a':
			i += 5
			var val int
			val, i = getInt(input, i)
			execCycle()
			execCycle()
			regX += val
		}
	}
	return signal, screen
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	negative := false
	if in[pos] == '-' {
		negative = true
		pos++
	}
	for ; in[pos] != '\n'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	if negative {
		accumulator = -accumulator
	}
	return accumulator, pos
}

type screenBuffer [246]byte

var benchmark = false
