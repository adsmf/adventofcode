package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	s := newSolver(input)
	p1 := s.run(9, -1)
	p2 := s.run(1, 1)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type solver struct {
	zDiv  [14]int
	addX  [14]int
	addY  [14]int
	start int
	step  int
}

func newSolver(input string) solver {
	s := solver{}
	lines := utils.GetLines(input)
	for i := 0; i < 14; i++ {
		s.zDiv[i] = utils.GetInts(lines[i*18+4][6:])[0]
		s.addX[i] = utils.GetInts(lines[i*18+5][6:])[0]
		s.addY[i] = utils.GetInts(lines[i*18+15][6:])[0]
	}
	return s
}

func (s solver) run(start, step int) int {
	s.start = start
	s.step = step
	return s.search(0, 0, 0)
}

func (s solver) search(prefix int, pos int, prevZ int) int {
	if pos == 14 {
		if prevZ != 0 {
			return -1
		}
		return prefix
	}
	max := 1
	for _, div := range s.zDiv[pos:] {
		max *= div
	}
	if prevZ > max {
		return -1
	}
	if s.addX[pos] < 0 {
		digit := s.addX[pos] + prevZ%26
		if digit < 1 || digit > 9 {
			return -1
		}
		z := s.calcZ(pos, prevZ, digit)
		val := s.search(prefix*10+digit, pos+1, z)
		if val > 0 {
			return val
		}
	} else {
		for i, digit := 0, s.start; i < 9; i, digit = i+1, digit+s.step {
			z := s.calcZ(pos, prevZ, digit)
			val := s.search(prefix*10+digit, pos+1, z)
			if val > 0 {
				return val
			}
		}
	}
	return -1
}

func (s solver) calcZ(pos int, z int, w int) int {
	x := s.addX[pos] + z%26
	z /= s.zDiv[pos]
	if x != w {
		z *= 26
		z += w + s.addY[pos]
	}
	return z
}

var benchmark = false
