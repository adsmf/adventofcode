package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	monkeys := loadMonkeys()
	monkeyLCM := monkeys[0].testVal
	for i := 1; i < len(monkeys); i++ {
		monkeyLCM *= monkeys[i].testVal
	}
	p1 := keepAway(monkeys, 20, func(i monkeyItem) monkeyItem { return i / 3 })
	p2 := keepAway(monkeys, 10000, func(i monkeyItem) monkeyItem { return i % monkeyLCM })
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func keepAway(monkeys tribe, rounds int, reduce transform) int {
	inspectCount := [8]int{}
	inspections := [...]transform{
		monkeys[0].inspectOp(),
		monkeys[1].inspectOp(),
		monkeys[2].inspectOp(),
		monkeys[3].inspectOp(),
		monkeys[4].inspectOp(),
		monkeys[5].inspectOp(),
		monkeys[6].inspectOp(),
		monkeys[7].inspectOp(),
	}
	var target *monkeyInfo
	for round := 0; round < rounds; round++ {
		for m := 0; m < 8; m++ {
			monkey := &monkeys[m]
			for i := 0; i < monkey.numItems; i++ {
				inspectCount[m]++
				item := reduce(inspections[m](monkey.items[i]))
				if (item % monkey.testVal) == 0 {
					target = &monkeys[monkey.throwTrue]
				} else {
					target = &monkeys[monkey.throwFalse]
				}
				target.items[target.numItems] = item
				target.numItems++
			}
			monkey.numItems = 0
		}
	}
	w1, w2 := 0, 0
	for i := 0; i < 8; i++ {
		if inspectCount[i] > w1 {
			w1 = inspectCount[i]
			if w1 > w2 {
				w1, w2 = w2, w1
			}
		}
	}
	return int(w1 * w2)
}

func loadMonkeys() tribe {
	monkeys := tribe{}
	for pos := 7; pos < len(input); {
		monkeyID := input[pos] - '0'
		monkey := &monkeys[monkeyID]
		pos = monkeys[monkeyID].parseItems(input, pos+21)
		switch input[pos+23] {
		case '+':
			monkey.op.code = opAdd
			monkey.op.value, pos = getInt(input, pos+25)
		case '*':
			if input[pos+25] == 'o' {
				monkey.op.code = opTimesSelf
				pos += 28
				break
			}
			monkey.op.code = opTimes
			monkey.op.value, pos = getInt(input, pos+25)
		}
		monkey.testVal, pos = getInt(input, pos+21)
		monkey.throwTrue = input[pos+30] & 0xf
		monkey.throwFalse = input[pos+62] & 0xf
		pos += 72
	}
	return monkeys
}

type transform func(monkeyItem) monkeyItem
type tribe [8]monkeyInfo

type monkeyInfo struct {
	items      monkeyItems
	numItems   int
	op         inspectOperation
	testVal    monkeyItem
	throwTrue  byte
	throwFalse byte
}

func (m monkeyInfo) inspectOp() transform {
	switch m.op.code {
	case opAdd:
		val := m.op.value
		return func(item monkeyItem) monkeyItem { return item + val }
	case opTimes:
		val := m.op.value
		return func(item monkeyItem) monkeyItem { return item * val }
	case opTimesSelf:
		return func(item monkeyItem) monkeyItem { return item * item }
	}
	return nil
}

func (m *monkeyInfo) parseItems(input []byte, pos int) int {
	accumulator := monkeyItem(0)
	addItem := func() {
		m.items[m.numItems] = accumulator
		m.numItems++
		accumulator = 0
	}
	for ; input[pos] != '\n'; pos++ {
		ch := input[pos]
		if ch >= '0' {
			accumulator *= 10
			accumulator += monkeyItem(ch & 0xf)
			continue
		}
		addItem()
		pos++
	}
	addItem()
	return pos + 1
}

func getInt(in []byte, pos int) (monkeyItem, int) {
	accumulator := monkeyItem(0)
	for ; in[pos] != '\n'; pos++ {
		accumulator *= 10
		accumulator += monkeyItem(in[pos] & 0xf)
	}
	return accumulator, pos
}

type monkeyItem uint
type monkeyItems [30]monkeyItem

type inspectOperation struct {
	code  opCode
	value monkeyItem
}

type opCode int

const (
	opTimes opCode = iota
	opAdd
	opTimesSelf
)

var benchmark = false
