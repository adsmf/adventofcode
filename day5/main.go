package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Day 1: %d\n", day1())
	fmt.Printf("Day 2: %d\n", day2())
}

func day1() int {
	return runInput(1)
}

func day2() int {
	return 0
}

func runInput(input int) int {
	inputString := loadInputString()
	tape := newMachine(inputString, input)
	tape.run()
	return tape.lastOutput
}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}

type machine struct {
	headPos    int
	values     map[int]int
	input      int
	lastOutput int
}

func newMachine(initial string, input int) machine {
	initialValueStrings := strings.Split(strings.TrimSpace(initial), ",")
	initialValues := map[int]int{}
	for pos, valString := range initialValueStrings {
		val, err := strconv.Atoi(valString)
		if err != nil {
			panic(err)
		}
		initialValues[pos] = val
	}
	mach := machine{
		values:  initialValues,
		headPos: 0,
		input:   input,
	}
	return mach
}

func (t *machine) run() {
	for {
		done := t.step()
		if done {
			return
		}
	}
}

func (t *machine) step() bool {
	oper := t.values[t.headPos]
	paramModes := int(oper / 100)
	// fmt.Printf("Inst: %d (%d)\n", oper, paramModes)
	oper = oper % 100
	switch oper {
	case 1:
		// Add
		p1 := t.getVal(t.headPos + 1)
		p2 := t.getVal(t.headPos + 2)
		p3 := t.getVal(t.headPos + 3)
		if t.paramMode(paramModes, 0) == 0 {
			p1 = t.getVal(p1)
		}
		if t.paramMode(paramModes, 1) == 0 {
			p2 = t.getVal(p2)
		}
		t.values[p3] = p1 + p2
		t.headPos = t.headPos + 4
	case 2:
		// Mult
		p1 := t.getVal(t.headPos + 1)
		p2 := t.getVal(t.headPos + 2)
		p3 := t.getVal(t.headPos + 3)

		if t.paramMode(paramModes, 0) == 0 {
			p1 = t.getVal(p1)
		}
		if t.paramMode(paramModes, 1) == 0 {
			p2 = t.getVal(p2)
		}
		t.values[p3] = p1 * p2
		t.headPos = t.headPos + 4
	case 3:
		// Input
		p := t.getVal(t.headPos + 1)
		fmt.Printf("Storing input in %d\n", p)
		t.values[p] = t.input
		t.headPos = t.headPos + 2
	case 4:
		// Output
		p := t.getVal(t.headPos + 1)
		t.lastOutput = t.values[p]
		fmt.Printf("Value at %d is %d\n", p, t.values[p])
		t.headPos = t.headPos + 2
	case 99:
		return true
	default:
		panic(fmt.Errorf("Invalid opcode %d at position %d: %#v", oper, t.headPos, t))
	}
	return false
}

func (t *machine) paramMode(modes, pos int) int {
	mask := int(math.Pow(10, float64(pos)))
	return (modes / mask) % 10
}

func (t *machine) getVal(pos int) int {
	if pos >= len(t.values) {
		return 0
	}
	return t.values[pos]
}

func (t *machine) String() string {
	valueStrings := []string{}
	keys := []int{}
	for key := range t.values {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		valueStrings = append(valueStrings, strconv.Itoa(t.values[key]))
	}
	return strings.Join(valueStrings, ",")
}
