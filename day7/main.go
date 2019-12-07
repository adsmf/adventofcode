package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2019/utils"
)

// var debug = debugPrintf
var debug = noOut

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	inputString := loadInputString()
	_, thrust := findBestSequence(inputString)
	return thrust
}

func part2() int {
	return 0
}

func findBestSequence(program string) ([]int, int) {
	phases := []int{0, 1, 2, 3, 4}
	bestSeq := []int{}
	bestThrust := 0
	perms := utils.PermuteInts(phases)
	for _, phase := range perms {
		thrust := testPhaseSequence(phase, program)
		if thrust > bestThrust {
			bestThrust = thrust
			bestSeq = phase
		}
	}
	return bestSeq, bestThrust
}

func testPhaseSequence(phases []int, program string) int {
	signal := 0
	for _, phase := range phases {
		inputs := []int{phase, signal}
		tape := newMachine(program, inputs)
		tape.run()
		signal = tape.lastOutput
	}
	return signal
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
	inputs     []int
	lastOutput int
}

func newMachine(initial string, inputs []int) machine {
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
		inputs:  inputs,
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
	initialHead := t.headPos
	oper := t.values[initialHead]
	paramModes := int(oper / 100)
	oper = oper % 100
	debug("\t%04d\tInst: %d #%d\n", initialHead, oper, paramModes)
	switch oper {
	case 1:
		// Add
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		debug("\t\t\tAdd: %d + %d => %d\n", p1, p2, p3)
		t.values[p3] = p1 + p2
	case 2:
		// Mult
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		debug("\t\t\tMul: %d * %d => %d\n", p1, p2, p3)
		t.values[p3] = p1 * p2
	case 3:
		// Input
		params := t.getParams(paramModes, 1, true)
		p := params[0]

		nextInput := t.inputs[0]
		t.inputs = t.inputs[1:]
		debug("\t\t\tStore: %d => %d\n", nextInput, p)
		t.values[p] = nextInput
	case 4:
		// Output
		params := t.getParams(paramModes, 1, false)
		p := params[0]
		t.lastOutput = p
		debug("\t\t\tRead: %d is %d\n", p, t.lastOutput)
	case 5:
		// Opcode 5 is jump-if-true:
		//   if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter.
		//   Otherwise, it does nothing.

		params := t.getParams(paramModes, 2, false)
		p1 := params[0]
		p2 := params[1]

		debug("\t\t\tJump if %d != 0\n", p1)
		if p1 != 0 {
			debug("\t\t\tJumping to %d\n", p2)
			t.headPos = p2
		}
	case 6:
		// Opcode 6 is jump-if-false:
		//   if the first parameter is zero, it sets the instruction pointer to the value from the second parameter.
		//   Otherwise, it does nothing.
		params := t.getParams(paramModes, 2, false)
		p1 := params[0]
		p2 := params[1]

		debug("\t\t\tJump if %d == 0\n", p1)
		if p1 == 0 {
			debug("\t\t\tJumping to %d\n", p2)
			t.headPos = p2
		}
	case 7:
		// Opcode 7 is less than:
		//   if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter.
		//   Otherwise, it stores 0.
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		debug("\t\t\tset %d < %d => %d\n", p1, p2, p3)
		if p1 < p2 {
			t.values[p3] = 1
		} else {
			t.values[p3] = 0
		}
	case 8:
		// Opcode 8 is equals:
		//   if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter.
		//   Otherwise, it stores 0.
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		debug("\t\t\tset %d == %d => %d\n", p1, p2, p3)
		if p1 == p2 {
			t.values[p3] = 1
		} else {
			t.values[p3] = 0
		}
	case 99:
		return true
	default:
		panic(fmt.Errorf("Invalid opcode %d at position %d: %#v", oper, t.headPos, t))
	}
	return false
}

func (t *machine) getParams(paramModes, numParams int, hasOutput bool) []int {
	params := []int{}
	for param := 0; param < numParams; param++ {
		lastParam := (param == numParams-1)
		p := t.getVal(t.headPos + param + 1)
		if !hasOutput || !lastParam {
			if t.paramMode(paramModes, param) == 0 {
				p = t.getVal(p)
			}
		}
		params = append(params, p)
	}

	t.headPos = t.headPos + numParams + 1
	return params
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

func debugPrintf(format string, params ...interface{}) {
	fmt.Printf(format, params...)
}
func noOut(format string, params ...interface{}) {}
