package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"

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
	_, thrust := findBestSequence(inputString, false)
	return thrust
}

func part2() int {
	inputString := loadInputString()
	_, thrust := findBestSequence(inputString, true)
	return thrust
}

func findBestSequence(program string, feedback bool) ([]int, int) {
	var phases []int
	if feedback {
		phases = []int{5, 6, 7, 8, 9}
	} else {
		phases = []int{0, 1, 2, 3, 4}
	}
	bestSeq := []int{}
	bestThrust := 0
	perms := utils.PermuteInts(phases)
	for _, phase := range perms {
		var thrust int
		if feedback {
			thrust = testPhaseSequenceWithFeedback(phase, program)
		} else {
			thrust = testPhaseSequence(phase, program)
		}
		if thrust > bestThrust {
			bestThrust = thrust
			bestSeq = phase
		}
	}
	return bestSeq, bestThrust
}

func testPhaseSequence(phases []int, program string) int {
	signal := 0
	output := make(chan int, 1)
	for _, phase := range phases {
		inputs := make(chan int, 2)
		inputs <- phase
		inputs <- signal
		tape := newMachine(program, inputs, output)
		tape.run()
		close(inputs)
		signal = <-output
	}
	close(output)
	return signal
}

func testPhaseSequenceWithFeedback(phases []int, program string) int {
	thrust := 0
	wires := make([](chan int), len(phases))
	for id := range phases {
		wires[id] = make(chan int, 2)
	}
	machines := make([]machine, len(phases))
	for id, phase := range phases {
		wires[id] <- phase
		machines[id] = newMachine(program, wires[id], wires[(id+1)%len(wires)])
	}
	wires[0] <- 0
	wg := sync.WaitGroup{}
	for _, m := range machines {
		wg.Add(1)
		go func(m machine) {
			m.run()
			wg.Done()
		}(m)
	}
	wg.Wait()
	thrust = <-wires[0]
	return thrust
}

func loadInputString() string {
	inputRaw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(inputRaw)

}

type machine struct {
	headPos int
	values  map[int]int
	inputs  <-chan int
	outputs chan<- int
}

func newMachine(initial string, inputs <-chan int, output chan<- int) machine {
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
		outputs: output,
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
	switch oper {
	case 1:
		// Add
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		t.values[p3] = p1 + p2
	case 2:
		// Mult
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		t.values[p3] = p1 * p2
	case 3:
		// Input
		params := t.getParams(paramModes, 1, true)
		p := params[0]

		nextInput := <-t.inputs
		t.values[p] = nextInput
	case 4:
		// Output
		params := t.getParams(paramModes, 1, false)
		p := params[0]

		t.outputs <- p
	case 5:
		// JNZ
		params := t.getParams(paramModes, 2, false)
		p1 := params[0]
		p2 := params[1]

		if p1 != 0 {
			t.headPos = p2
		}
	case 6:
		// JEZ
		params := t.getParams(paramModes, 2, false)
		p1 := params[0]
		p2 := params[1]

		if p1 == 0 {
			t.headPos = p2
		}
	case 7:
		// CLT
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		if p1 < p2 {
			t.values[p3] = 1
		} else {
			t.values[p3] = 0
		}
	case 8:
		// CMP
		params := t.getParams(paramModes, 3, true)
		p1 := params[0]
		p2 := params[1]
		p3 := params[2]

		if p1 == p2 {
			t.values[p3] = 1
		} else {
			t.values[p3] = 0
		}
	case 99:
		// HCF
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
			if paramMode(paramModes, param) == 0 {
				p = t.getVal(p)
			}
		}
		params = append(params, p)
	}

	t.headPos = t.headPos + numParams + 1
	return params
}

func paramMode(modes, pos int) int {
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
