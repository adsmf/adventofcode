package main

import (
	"fmt"
	"strings"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	p1 := part1()
	p2 := part2()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	state := load("input.txt")
	turns, solved := allToFourth(state)
	if !solved {
		fmt.Println("UNSOLVED")
	}
	return turns
}

func part2() int {
	state := load("input.txt")
	state.chips["elerium"] = 1
	state.chips["dilithium"] = 1
	state.rtgs["elerium"] = 1
	state.rtgs["dilithium"] = 1
	turns, solved := allToFourth(state)
	if !solved {
		fmt.Println("UNSOLVED")
	}
	return turns
}

func allToFourth(initialState facilityState) (int, bool) {
	genIndices(initialState)
	initialHash := initialState.hashIndexed()
	seenStates := map[int]int{initialHash: 0}
	openStates := map[int]facilityState{initialHash: initialState}
	for turn := 0; ; turn++ {
		nextOpenStates := map[int]facilityState{}
		for _, state := range openStates {
			done := true
			for _, floor := range state.rtgs {
				if floor != 4 {
					done = false
					break
				}
			}
			for _, floor := range state.chips {
				if floor != 4 {
					done = false
					break
				}
			}
			if done {
				return turn, true
			}
			for _, nextState := range state.next() {
				hash := nextState.hashIndexed()
				if _, found := seenStates[hash]; !found {
					seenStates[hash] = turn
					nextOpenStates[hash] = nextState
				}
			}
		}
		openStates = nextOpenStates
		if len(openStates) == 0 {
			return turn, false
		}
	}
}

type facilityState struct {
	elevatorFloor int
	rtgs          map[string]int
	chips         map[string]int
}

func (f facilityState) next() []facilityState {
	nextStates := []facilityState{}
	rtgsOnFloor := []string{}
	chipsOnFloor := []string{}
	for rtg, floor := range f.rtgs {
		if floor == f.elevatorFloor {
			rtgsOnFloor = append(rtgsOnFloor, rtg)
		}
	}
	for chip, floor := range f.chips {
		if floor == f.elevatorFloor {
			chipsOnFloor = append(chipsOnFloor, chip)
		}
	}
	elevatorDirections := []int{}
	if f.elevatorFloor > 1 {
		elevatorDirections = append(elevatorDirections, -1)
	}
	if f.elevatorFloor < 4 {
		elevatorDirections = append(elevatorDirections, 1)
	}
	for moveRTGs := 0; moveRTGs <= len(rtgsOnFloor) && moveRTGs <= 2; moveRTGs++ {
		rtgCombs := [][]string{}
		if moveRTGs == 0 {
			rtgCombs = append(rtgCombs, []string{})
		} else {
			utils.IterateCombinations(len(rtgsOnFloor), moveRTGs, func(iter []int) {
				choice := []string{}
				for _, i := range iter {
					choice = append(choice, rtgsOnFloor[i])
				}
				rtgCombs = append(rtgCombs, choice)
			})
		}
		for moveChips := 0; moveChips <= len(chipsOnFloor) && moveChips <= 2-moveRTGs; moveChips++ {
			if moveChips+moveRTGs == 0 {
				continue
			}
			chipCombs := [][]string{}
			if moveChips == 0 {
				chipCombs = append(chipCombs, []string{})
			} else {
				utils.IterateCombinations(len(chipsOnFloor), moveChips, func(iter []int) {
					choice := []string{}
					for _, i := range iter {
						choice = append(choice, chipsOnFloor[i])
					}
					chipCombs = append(chipCombs, choice)
				})
			}
			for r := 0; r < len(rtgCombs); r++ {
				for c := 0; c < len(chipCombs); c++ {
					for _, dir := range elevatorDirections {
						nextState := f.clone()
						nextState.elevatorFloor += dir
						for _, rtg := range rtgCombs[r] {
							nextState.rtgs[rtg] = nextState.elevatorFloor
						}
						for _, chip := range chipCombs[c] {
							nextState.chips[chip] = nextState.elevatorFloor
						}
						if nextState.valid() {
							nextStates = append(nextStates, nextState)
						}
					}
				}
			}
		}
	}
	return nextStates
}

func (f facilityState) clone() facilityState {
	clone := facilityState{
		elevatorFloor: f.elevatorFloor,
		rtgs:          make(map[string]int, len(f.rtgs)),
		chips:         make(map[string]int, len(f.chips)),
	}
	for rtg, floor := range f.rtgs {
		clone.rtgs[rtg] = floor
	}
	for chip, floor := range f.chips {
		clone.chips[chip] = floor
	}
	return clone
}

func (f facilityState) hashIndexed() int {
	hash := f.elevatorFloor * 10000000000
	for rtg, floor := range f.rtgs {
		hash += (indices[rtg] * 10 * floor)
	}
	for chip, floor := range f.chips {
		hash += (indices[chip] * 10 * floor) * 100000
	}
	return hash
}

var indices = map[string]int{}

func genIndices(state facilityState) {
	idx := 1
	indices = map[string]int{}
	for rtg := range state.rtgs {
		indices[rtg] = idx
		idx++
	}
}

func (f facilityState) valid() bool {
	floorsWithRTGs := map[int]bool{}
	for _, floor := range f.rtgs {
		floorsWithRTGs[floor] = true
	}
	for chip, chipFloor := range f.chips {
		if floorsWithRTGs[chipFloor] && f.rtgs[chip] != chipFloor {
			return false
		}
	}
	return true
}

func load(filename string) facilityState {
	facility := facilityState{
		elevatorFloor: 1,
		rtgs:          map[string]int{},
		chips:         map[string]int{},
	}
	lines := utils.ReadInputLines(filename)
	for floor := 1; floor < 4; floor++ {
		line := lines[floor-1]
		parts := strings.Split(line, "contains")
		contentString := strings.Trim(parts[1], " .")
		contentString = strings.ReplaceAll(contentString, ", and", ",")
		contentString = strings.ReplaceAll(contentString, " and ", ", ")
		contentString = strings.ReplaceAll(contentString, "-compatible", "")
		contents := strings.Split(contentString, ", ")
		for _, content := range contents {
			content = content[2:]
			parts := strings.Split(content, " ")
			switch parts[1] {
			case "generator":
				facility.rtgs[parts[0]] = floor
			case "microchip":
				facility.chips[parts[0]] = floor
			}
		}
	}
	return facility
}

var benchmark = false
