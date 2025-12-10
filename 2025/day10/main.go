package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	totalPresses := 0
	utils.EachLine(input, func(index int, line string) (done bool) {
		targetState, buttons, _ := parseLine(line)
		open, next := []stateInfo{{}}, []stateInfo{}
		seen := map[lightState]bool{}
		for len(open) > 0 {
			for _, state := range open {
				for idx, button := range buttons {
					option := state.nextState(idx, button)
					if targetState == option.lights {
						totalPresses += option.presses
						return false
					}
					if seen[option.lights] {
						continue
					}
					seen[option.lights] = true
					next = append(next, option)
				}
			}
			open, next = next, open[0:0]
		}
		panic("No solution")
	})
	return totalPresses
}

func part2() int {
	return -1
}

type stateInfo struct {
	lights  lightState
	buttons []int
	presses int
}

func (s stateInfo) String() string {
	return fmt.Sprintf("%08b (%d: %v)", s.lights, s.presses, s.buttons)
}

func (s stateInfo) nextState(buttonIdx int, toggle lightState) stateInfo {
	next := stateInfo{
		lights:  s.lights ^ toggle,
		presses: s.presses + 1,
		buttons: append(s.buttons, buttonIdx),
	}
	return next
}

type lightState uint16

func parseLine(line string) (lightState, []lightState, []int) {
	lineParts := strings.Split(line, " ")
	reqRaw := lineParts[0]
	buttonsRaw := lineParts[1 : len(lineParts)-1]
	joltages := utils.GetInts(lineParts[len(lineParts)-1])
	var targetState lightState
	for i := range len(reqRaw) - 2 {
		if reqRaw[i+1] == '#' {
			targetState |= 1 << i
		}
	}
	buttons := make([]lightState, len(buttonsRaw))
	for i, raw := range buttonsRaw {
		vals := utils.GetInts(raw)
		var toggle lightState
		for _, v := range vals {
			toggle |= 1 << v
		}
		buttons[i] = toggle
	}
	return targetState, buttons, joltages
}

var benchmark = false
