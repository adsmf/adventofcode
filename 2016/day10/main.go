package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1, p2 := run()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type preparedAssignment struct {
	value, targetBot int
}

func run() (int, int) {
	bots := map[int]*bot{}
	outputs := map[int][]int{}

	valuesToAssign := []preparedAssignment{}

	for _, line := range utils.ReadInputLines("input.txt") {
		if strings.HasPrefix(line, "value") {
			parts := utils.GetInts(line)
			valuesToAssign = append(valuesToAssign, preparedAssignment{parts[0], parts[1]})
			if _, found := bots[parts[1]]; !found {
				bots[parts[1]] = &bot{id: parts[1]}
			}
		} else {
			matcher := regexp.MustCompile(`^bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)$`)
			matches := matcher.FindStringSubmatch(line)
			botId, _ := strconv.Atoi(matches[1])
			if _, found := bots[botId]; !found {
				bots[botId] = &bot{id: botId}
			}
			lowID, _ := strconv.Atoi(matches[3])
			bots[botId].outputLow = wiring{
				objectType: matches[2],
				ordinal:    lowID,
			}

			highID, _ := strconv.Atoi(matches[5])
			bots[botId].outputHigh = wiring{
				objectType: matches[4],
				ordinal:    highID,
			}
		}
	}
	p1 := -1
	for _, assignment := range valuesToAssign {
		trigger := bots[assignment.targetBot].give(assignment.value, bots, outputs)
		if trigger >= 0 {
			p1 = trigger
		}
	}
	return p1, outputs[0][0] * outputs[1][0] * outputs[2][0]
}

type bot struct {
	id         int
	chips      []int
	outputLow  wiring
	outputHigh wiring
}

func (b *bot) give(newChip int, bots map[int]*bot, outputs map[int][]int) int {
	result := -1
	if b.chips == nil {
		b.chips = []int{}
	}
	b.chips = append(b.chips, newChip)
	if len(b.chips) == 2 {
		high, low := b.chips[0], b.chips[1]
		if low > high {
			high, low = low, high
		}
		if high == 61 && low == 17 {
			result = b.id
		}
		b.chips = []int{}

		targetHigh, targetLow := b.outputHigh.ordinal, b.outputLow.ordinal

		switch b.outputHigh.objectType {
		case "bot":
			trigger := bots[targetHigh].give(high, bots, outputs)
			if trigger >= 0 {
				result = trigger
			}
		case "output":
			outputs[targetHigh] = append(outputs[targetHigh], high)
		}
		switch b.outputLow.objectType {
		case "bot":
			trigger := bots[targetLow].give(low, bots, outputs)
			if trigger >= 0 {
				result = trigger
			}
		case "output":
			outputs[targetLow] = append(outputs[targetLow], low)
		}
		b.chips = []int{}
	}
	return result
}

type wiring struct {
	objectType string
	ordinal    int
}

var benchmark = false
