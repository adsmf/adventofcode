package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	p1, _ := search(0, 0, 0, 9, -1)
	p2, _ := search(0, 0, 0, 1, 1)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func search(prefix int, pos int, prevZ int, start, step int) (int, bool) {
	if pos == 14 {
		return prefix, prevZ == 0
	}
	max := 1
	for _, div := range zDiv[pos:] {
		max *= div
	}
	if prevZ > max {
		return -1, false
	}
	if addX[pos] < 0 {
		digit := addX[pos] + prevZ%26
		if digit < 1 || digit > 9 {
			return -1, false
		}
		z := calcZ(pos, prevZ, digit)
		val, found := search(prefix*10+digit, pos+1, z, start, step)
		if found {
			return val, true
		}
	} else {
		for i := 0; i < 9; i++ {
			digit := start + i*step
			z := calcZ(pos, prevZ, digit)
			val, found := search(prefix*10+digit, pos+1, z, start, step)
			if found {
				return val, true
			}
		}
	}
	return -1, false
}

var zDiv = [14]int{1, 1, 1, 26, 1, 1, 1, 26, 26, 1, 26, 26, 26, 26}
var addX = [14]int{13, 11, 15, -5, 14, 10, 12, -14, -8, 13, 0, -5, -9, -1}
var addY = [14]int{0, 3, 8, 5, 13, 9, 6, 1, 1, 2, 7, 5, 8, 15}

func calcZ(pos int, z int, w int) int {
	x := addX[pos] + z%26
	z /= zDiv[pos]
	if x != w {
		z *= 26
		z += w + addY[pos]
	}
	return z
}

var benchmark = false
