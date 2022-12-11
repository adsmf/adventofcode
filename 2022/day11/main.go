package main

import (
	_ "embed"
	"fmt"
	"sort"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input []byte

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	monkeys := loadMonkeys()
	inspectCount := make([]int, 8)
	for round := 0; round < 20; round++ {
		for m := 0; m < 8; m++ {
			monkey := monkeys[m]
			toThrow := monkey.items
			monkeys[m].items = monkey.items[:0]
			for i := 0; i < len(toThrow); i++ {
				inspectCount[m]++
				item := toThrow[i]
				switch monkey.op.code {
				case opAdd:
					item += monkey.op.value
				case opTimes:
					item *= monkey.op.value
				case opTimesSelf:
					item *= item
				}
				item /= 3
				if (item % monkey.testVal) == 0 {
					monkeys[monkey.throwTrue].items = append(monkeys[monkey.throwTrue].items, item)
				} else {
					monkeys[monkey.throwFalse].items = append(monkeys[monkey.throwFalse].items, item)
				}
			}
		}
	}
	sort.Ints(inspectCount)
	worry := inspectCount[6] * inspectCount[7]
	return worry
}

func part2() int {
	monkeys := loadMonkeys()
	inspectCount := make([]int, 8)
	monkeyTests := make([]int, 0, 8)
	for _, monkey := range monkeys {
		monkeyTests = append(monkeyTests, monkey.testVal)
	}
	monkeyLCM := utils.LowestCommonMultipleInt(monkeyTests...)
	for round := 0; round < 10000; round++ {
		for m := 0; m < 8; m++ {
			monkey := monkeys[m]
			toThrow := monkey.items
			monkeys[m].items = monkey.items[:0]
			for i := 0; i < len(toThrow); i++ {
				inspectCount[m]++
				item := toThrow[i]
				switch monkey.op.code {
				case opAdd:
					item += monkey.op.value
				case opTimes:
					item *= monkey.op.value
				case opTimesSelf:
					item *= item
				}
				item %= monkeyLCM
				if (item % monkey.testVal) == 0 {
					monkeys[monkey.throwTrue].items = append(monkeys[monkey.throwTrue].items, item)
				} else {
					monkeys[monkey.throwFalse].items = append(monkeys[monkey.throwFalse].items, item)
				}
			}
		}
	}
	sort.Ints(inspectCount)
	worry := inspectCount[6] * inspectCount[7]
	return worry
}

func loadMonkeys() [8]monkeyInfo {

	monkeys := [8]monkeyInfo{}

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

type monkeyInfo struct {
	items      monkeyItems
	op         inspectOperation
	testVal    int
	throwTrue  byte
	throwFalse byte
}

func (m *monkeyInfo) parseItems(input []byte, pos int) int {
	accumulator := 0
	for ; input[pos] != '\n'; pos++ {
		ch := input[pos]
		if ch >= '0' {
			accumulator *= 10
			accumulator += int(ch & 0xf)
			continue
		}
		m.items = append(m.items, accumulator)
		accumulator = 0
		pos++
	}
	m.items = append(m.items, accumulator)
	return pos + 1
}

func getInt(in []byte, pos int) (int, int) {
	accumulator := 0
	negative := false
	if in[pos] == '-' {
		negative = true
		pos++
	}
	for ; in[pos] != '\n'; pos++ {
		accumulator *= 10
		accumulator += int(in[pos] & 0xf)
	}
	if negative {
		accumulator = -accumulator
	}
	return accumulator, pos
}

type monkeyItems []int

type inspectOperation struct {
	code  opCode
	value int
}

type opCode int

const (
	opTimes opCode = iota
	opAdd
	opTimesSelf
)

var benchmark = false
