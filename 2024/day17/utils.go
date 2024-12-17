package main

import (
	"fmt"
	"strconv"

	"github.com/adsmf/adventofcode/utils"
)

func decodeProgram() int {
	vals := utils.GetInts(input)
	inst := vals[3:]
	ip := 0
	decodeCombo := func(combo uint64) string {
		if combo <= 3 {
			return strconv.Itoa(int(combo))
		}
		if combo >= 7 {
			panic(fmt.Sprintf("Bad combo: %d", combo))
		}
		return string(byte(combo - 4 + 'A'))
	}
	for ip < len(inst) {
		op := opCode(inst[ip])
		literal := inst[ip+1]
		ip += 2
		switch op {
		case opADV:
			fmt.Printf("A >>= %s\n", decodeCombo(uint64(literal)))
		case opBXL:
			fmt.Printf("B ^= %d\n", literal)
		case opBST:
			fmt.Printf("B = %s & 7\n", decodeCombo(uint64(literal)))
		case opJNZ:
			fmt.Printf("JNZ %d\n", literal)
		case opBXC:
			fmt.Println("B ^= C")
		case opOUT:
			fmt.Printf("Out %s & 7\n", decodeCombo(uint64(literal)))
		case opBDV:
			fmt.Printf("B = A >> %s\n", decodeCombo(uint64(literal)))
		case opCDV:
			fmt.Printf("C = A >> %s\n", decodeCombo(uint64(literal)))
		}
	}
	return -1
}

type opCode int

const (
	opADV opCode = iota
	opBXL
	opBST
	opJNZ
	opBXC
	opOUT
	opBDV
	opCDV
)
