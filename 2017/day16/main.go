package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	ops := choreograph("input.txt")
	p1 := dance(ops, 1)
	p2 := dance(ops, 1000000000)
	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

func dance(operations []operation, reps int) string {
	numProgs := 16
	progs := make([]byte, numProgs)
	for i := 0; i < numProgs; i++ {
		progs[i] = byte(i) + 'a'
	}
	seen := make(map[string]int, 1000)
	for iter := 0; iter < reps; iter++ {
		for _, op := range operations {
			switch op.op {
			case opSpin:
				size := op.arg1.(int)
				progs = append(progs[numProgs-size:], progs[:numProgs-size]...)
			case opExchange:
				posA, posB := op.arg1.(int), op.arg2.(int)
				progs[posA], progs[posB] = progs[posB], progs[posA]
			case opPartner:
				c1 := op.arg1.(byte)
				c2 := op.arg2.(byte)
				for i := 0; i < len(progs); i++ {
					if progs[i] == c1 {
						progs[i] = c2
					} else if progs[i] == c2 {
						progs[i] = c1
					}
				}
			}
		}
		state := string(progs)
		if lastSaw, found := seen[state]; found {
			offset := (reps - 1) % (iter - lastSaw)
			for state, onTurn := range seen {
				if onTurn == offset {
					return state
				}
			}
		}
		seen[state] = iter
	}
	return string(progs)
}

type operationID int
type operation struct {
	op   operationID
	arg1 interface{}
	arg2 interface{}
}

const (
	opSpin operationID = iota
	opExchange
	opPartner
)

func choreograph(filename string) []operation {
	operations := []operation{}
	inputBytes, _ := ioutil.ReadFile(filename)
	for _, op := range strings.Split(strings.TrimSpace(string(inputBytes)), ",") {
		switch op[0] {
		case 's':
			size := utils.MustInt(op[1:])
			operations = append(operations, operation{
				op:   opSpin,
				arg1: size,
			})
		case 'x':
			positions := utils.GetInts(op[1:])
			posA, posB := positions[0], positions[1]
			operations = append(operations, operation{
				op:   opExchange,
				arg1: posA,
				arg2: posB,
			})
		case 'p':
			c1 := op[1]
			c2 := op[3]
			operations = append(operations, operation{
				op:   opPartner,
				arg1: c1,
				arg2: c2,
			})
		}
	}
	return operations
}

var benchmark = false
