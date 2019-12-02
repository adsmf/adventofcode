package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Day 1: %d\n", day1())
	fmt.Printf("Day 2: %d\n", day2())
}

func day1() int {
	return runInput(12, 2)
}

func day2() int {
	target := 19690720
	for input1 := 0; input1 <= 99; input1++ {
		for input2 := 0; input2 <= 99; input2++ {
			result := runInput(input1, input2)
			if result == target {
				return input1*100 + input2
			}
		}
	}
	return 0
}

func runInput(input1, input2 int) int {
	inputString := loadInputString()
	tape := newMachine(inputString)
	tape.values[1] = input1
	tape.values[2] = input2
	tape.run()
	return tape.values[0]
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
	values  []int
}

func newMachine(initial string) machine {
	initialValueStrings := strings.Split(strings.TrimSpace(initial), ",")
	initialValues := []int{}
	for _, valString := range initialValueStrings {
		val, err := strconv.Atoi(valString)
		if err != nil {
			panic(err)
		}
		initialValues = append(initialValues, val)
	}
	mach := machine{
		values:  initialValues,
		headPos: 0,
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
	switch oper {
	case 1:
		// Add
		pos1 := t.getVal(t.headPos + 1)
		pos2 := t.getVal(t.headPos + 2)
		pos3 := t.getVal(t.headPos + 3)
		t.values[pos3] = t.getVal(pos1) + t.getVal(pos2)
		t.headPos = t.headPos + 4
	case 2:
		// Mult
		pos1 := t.getVal(t.headPos + 1)
		pos2 := t.getVal(t.headPos + 2)
		pos3 := t.getVal(t.headPos + 3)
		t.values[pos3] = t.getVal(pos1) * t.getVal(pos2)
		t.headPos = t.headPos + 4
	case 99:
		return true
	default:
		panic(fmt.Errorf("Invalid opcode %d at position %d: %#v", oper, t.headPos, t))
	}
	return false
}

func (t *machine) getVal(pos int) int {
	if pos >= len(t.values) {
		return 0
	}
	return t.values[pos]
}

func (t *machine) String() string {
	valueStrings := []string{}
	for _, val := range t.values {
		valueStrings = append(valueStrings, strconv.Itoa(val))
	}
	return strings.Join(valueStrings, ",")
}
