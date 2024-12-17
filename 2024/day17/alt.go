package main

import (
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

func altMain() {
	p2 := altSolve()
	i := 0
	for ; p1buf[i] > 0; i++ {
	}
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1buf[:i])
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func altSolve() int {
	regA := 0
	inst := make([]int, 0, 16)
	utils.EachInteger(input, func(idx, value int) (done bool) {
		if idx == 0 {
			regA = value
			return
		}
		if idx < 3 {
			return
		}
		inst = append(inst, value)
		return
	})
	compiled := compile(inst)
	altPart1(regA, compiled)
	return altPart2(inst, compiled)
}

type loopRunner func(regA int) int

func compile(inst []int) loopRunner {
	var out int
	var regs [3]int

	comp := func() int { return out }
	ip := 0
	decodeCombo := func(combo int) int {
		if combo <= 3 {
			return int(combo)
		}
		if combo >= 7 {
			panic(fmt.Sprintf("Bad combo: %d", combo))
		}
		return regs[combo-4]
	}
	for ip < len(inst) {
		op := opCode(inst[ip])
		literal := inst[ip+1]
		ip += 2
		switch op {
		case opADV:
			prevComp := comp
			comp = func() int {
				prevComp()
				regs[0] >>= decodeCombo(literal)
				return out
			}
		case opBXL:
			prevComp := comp
			comp = func() int {
				prevComp()
				regs[1] ^= literal
				return out
			}
		case opBST:
			prevComp := comp
			comp = func() int {
				prevComp()
				regs[1] = decodeCombo(literal) & 7
				return out
			}
		case opJNZ:
			if literal != 0 {
				panic("Unhandled: JNZ target not zero")
			}
		case opBXC:
			prevComp := comp
			comp = func() int {
				prevComp()
				regs[1] ^= regs[2]
				return out
			}
		case opOUT:
			prevComp := comp
			comp = func() int {
				prevComp()
				out = decodeCombo(literal) & 7
				return out
			}
		case opBDV:
			prevComp := comp
			comp = func() int {
				prevComp()
				regs[1] = regs[0] >> decodeCombo(literal)
				return out
			}
		case opCDV:
			prevComp := comp
			comp = func() int {
				prevComp()
				regs[2] = regs[0] >> decodeCombo(literal)
				return out
			}
		}
	}

	return func(regA int) int {
		regs = [3]int{regA, 0, 0}
		return comp()
	}
}

func altPart1(regA int, nextOut loopRunner) {
	first := true
	i := 0
	for {
		if regA == 0 {
			break
		}
		if !first {
			p1buf[i] = ','
			i++
		}
		first = false
		val := nextOut(regA)
		p1buf[i] = '0' + byte(val)
		i++
		regA >>= 3
	}
}

func altPart2(inst []int, nextOut loopRunner) int {
	open := make([]int, 0, 25)
	next := make([]int, 0, 25)
	open = append(open, 0)
	var p2 int
	for i := len(inst) - 1; i >= 0; i-- {
		p2 = utils.MaxInt
		val := inst[i]
		for _, cur := range open {
			for i := range 8 {
				try := cur<<3 + int(i)
				if nextOut(try) == val {
					next = append(next, int(try))
					p2 = min(p2, try)
				}
			}
		}
		open, next = next, open[0:0]
	}
	return p2
}
