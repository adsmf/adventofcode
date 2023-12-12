package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve() (int, int) {
	totalP1, totalP2 := 0, 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		contitions, report := parseLine(line)
		totalP1 += countArrangements(append(contitions, condOperational), report)
		p2cond := make([]condition, 0, len(contitions)*5+4)
		p2report := make([]int, 0, len(report)*5)
		for i := 0; i < 5; i++ {
			p2cond = append(p2cond, contitions...)
			if i < 4 {
				p2cond = append(p2cond, condUnknown)
			}
			p2report = append(p2report, report...)
		}
		totalP2 += countArrangements(p2cond, p2report)
		return false
	})
	return totalP1, totalP2
}

func countArrangements(conditions []condition, report []int) int {
	cache := make(map[knownEntry]int, len(conditions)*10)
	return countSub(conditions, report, cache, 0, 0, 0)
}

func countSub(conditions []condition, report []int, known map[knownEntry]int, position int, groupOffset int, complete int) int {
	k := knownEntry{position, groupOffset, complete}
	if count, found := known[k]; found {
		return count
	}
	save := func(value int) int {
		known[k] = value
		return value
	}
	callSub := func(position int, groupOffset int, complete int) int {
		return countSub(conditions, report, known, position, groupOffset, complete)
	}
	if position == len(conditions) {
		if len(report) == complete {
			return save(1)
		}
		if len(report) == complete+1 && report[complete] == groupOffset {
			return save(1)
		}
		return save(0)
	}
	switch conditions[position] {
	case condOperational:
		if groupOffset == 0 {
			return save(callSub(position+1, 0, complete))
		} else if complete < len(report) && groupOffset == report[complete] {
			return save(callSub(position+1, 0, complete+1))
		}
		return save(0)
	case condDamaged:
		if complete >= len(report) {
			return save(0)
		}
		if groupOffset > report[complete] {
			return save(0)
		}
		return save(callSub(position+1, groupOffset+1, complete))
	default:
		assumeOperational, assumeDamaged := 0, 0
		// Assume damaged
		if complete < len(report) && groupOffset < report[complete] {
			assumeDamaged = callSub(position+1, groupOffset+1, complete)
		}

		// Assume operational
		if groupOffset == 0 {
			assumeOperational = callSub(position+1, 0, complete)
		} else if complete < len(report) && groupOffset == report[complete] {
			assumeOperational = callSub(position+1, 0, complete+1)
		}
		return save(assumeDamaged + assumeOperational)
	}
}

type knownEntry struct {
	position    int
	groupOffset int
	complete    int
}

func parseLine(line string) ([]condition, []int) {
	conditions := []condition{}
	report := []int{}
	accumulator := 0
	for _, ch := range line {
		switch ch {
		case '?', '.', '#':
			conditions = append(conditions, condition(ch))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			accumulator = accumulator*10 + int(ch-'0')
		case ',':
			report = append(report, accumulator)
			accumulator = 0
		}
	}
	report = append(report, accumulator)
	return conditions, report
}

type condition byte

const (
	condUnknown     condition = '?'
	condOperational condition = '.'
	condDamaged     condition = '#'
)

var benchmark = false
