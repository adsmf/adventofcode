package main

import (
	"fmt"
	"strings"
)

type opcode int

type operation int
type argument int

type opmap map[opcode]operation

type instruction struct {
	op     opcode
	input1 argument
	input2 argument
	output argument
}

// List opcodes in order from instructions until we know the actual mapping
const (
	addr operation = iota
	addi

	mulr
	muli

	banr
	bani

	borr
	bori

	setr
	seti

	gtir
	gtri
	gtrr

	eqir
	eqri
	eqrr

	opcodeEND
	opcodeSTART = addr
)

type registers []int

func processOp(op operation, inst instruction, input registers) registers {
	reg := make([]int, len(input))
	copy(reg, input)
	switch op {
	case addr, addi:
		inst.input1 = fromref(reg, inst.input1)
		if op == addr {
			inst.input2 = fromref(reg, inst.input2)
		}
		newVal := inst.input1 + inst.input2
		return putref(reg, inst.output, newVal)
	case mulr, muli:
		inst.input1 = fromref(reg, inst.input1)
		if op == mulr {
			inst.input2 = fromref(reg, inst.input2)
		}
		newVal := inst.input1 * inst.input2
		return putref(reg, inst.output, newVal)
	case banr, bani:
		inst.input1 = fromref(reg, inst.input1)
		if op == banr {
			inst.input2 = fromref(reg, inst.input2)
		}
		newVal := inst.input1 & inst.input2
		return putref(reg, inst.output, newVal)
	case borr, bori:
		inst.input1 = fromref(reg, inst.input1)
		if op == borr {
			inst.input2 = fromref(reg, inst.input2)
		}
		newVal := inst.input1 | inst.input2
		return putref(reg, inst.output, newVal)
	case setr:
		return putref(reg, inst.output, fromref(reg, inst.input1))
	case seti:
		return putref(reg, inst.output, inst.input1)
	case gtir, gtri, gtrr:
		if op == gtri || op == gtrr {
			inst.input1 = fromref(reg, inst.input1)
		}
		if op == gtir || op == gtrr {
			inst.input2 = fromref(reg, inst.input2)
		}
		var newVal argument
		if inst.input1 > inst.input2 {
			newVal = 1
		} else {
			newVal = 0
		}
		return putref(reg, inst.output, newVal)
	case eqir, eqri, eqrr:
		if op == eqri || op == eqrr {
			inst.input1 = fromref(reg, inst.input1)
		}
		if op == eqir || op == eqrr {
			inst.input2 = fromref(reg, inst.input2)
		}
		var newVal argument
		if inst.input1 == inst.input2 {
			newVal = 1
		} else {
			newVal = 0
		}
		return putref(reg, inst.output, newVal)
	default:
		panic(fmt.Sprintf("Operation not implemented: %d", op))
	}
}

func fromref(reg registers, ref argument) argument {
	debug("%v[%d] => %d", reg, ref, reg[ref])
	return argument(reg[ref])
}

func putref(reg registers, ref, value argument) registers {
	reg[ref] = int(value)
	return reg
}

func (op operation) toString() string {
	switch op {
	case addr:
		return "addr"
	case addi:
		return "addi"
	case mulr:
		return "mulr"
	case muli:
		return "muli"
	case banr:
		return "banr"
	case bani:
		return "bani"
	case borr:
		return "borr"
	case bori:
		return "bori"
	case setr:
		return "setr"
	case seti:
		return "seti"
	case gtir:
		return "gtir"
	case gtri:
		return "gtri"
	case gtrr:
		return "gtrr"
	case eqir:
		return "eqir"
	case eqri:
		return "eqri"
	case eqrr:
		return "eqrr"
	default:
		return "UNDEFINED"
	}
}

func opListToString(ops []operation) string {
	retStrings := []string{}
	for _, op := range ops {
		retStrings = append(retStrings, op.toString())
	}
	return strings.Join(retStrings, ", ")
}
