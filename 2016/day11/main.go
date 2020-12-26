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
	state.rtgs = append(state.rtgs, 1, 1)   // 2 new rtgs on floor 1
	state.chips = append(state.chips, 1, 1) // 2 corresponding chips on floor 1
	turns, solved := allToFourth(state)
	if !solved {
		fmt.Println("UNSOLVED")
	}
	return turns
}

func allToFourth(initialState facilityState) (int, bool) {
	initialHash := initialState.hash()
	seenStates := map[facilityHash]int{initialHash: 0}
	openStates := map[facilityHash]facilityState{initialHash: initialState}
	for turn := 0; turn < 60; turn++ {
		nextOpenStates := make(map[facilityHash]facilityState, len(openStates))
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
				hash := nextState.hash()
				if _, found := seenStates[hash]; !found {
					seenStates[hash] = turn
					nextOpenStates[hash] = nextState
				}
			}
		}
		openStates = nextOpenStates
		// fmt.Println(turn, "->", len(openStates))
		if len(openStates) == 0 {
			return turn, false
		}
	}
	fmt.Println(len(openStates))
	return -1, false
}

type facilityState struct {
	elevatorFloor int
	rtgs          []int
	chips         []int
}

func (f facilityState) next() []facilityState {
	nextStates := []facilityState{}
	rtgsOnFloor := []int{}
	chipsOnFloor := []int{}
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
		rtgCombs := [][]int{}
		if moveRTGs == 0 {
			rtgCombs = append(rtgCombs, []int{})
		} else {
			utils.IterateCombinations(len(rtgsOnFloor), moveRTGs, func(iter []int) {
				choice := []int{}
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
			chipCombs := [][]int{}
			if moveChips == 0 {
				chipCombs = append(chipCombs, []int{})
			} else {
				utils.IterateCombinations(len(chipsOnFloor), moveChips, func(iter []int) {
					choice := []int{}
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
		rtgs:          make([]int, len(f.rtgs)),
		chips:         make([]int, len(f.chips)),
	}
	copy(clone.rtgs, f.rtgs)
	copy(clone.chips, f.chips)
	return clone
}

type facilityHash uint64

func (f facilityState) hash() facilityHash {
	hash := facilityHash(0)
	for rtg, floor := range f.rtgs {
		itemHash := facilityHash(floor-1) << (rtg * 2)
		hash |= itemHash
	}
	offset := len(f.chips) + 2
	hash <<= offset * 2
	for chip, floor := range f.chips {
		itemHash := facilityHash(floor-1) << (chip * 2)
		hash |= itemHash
	}
	hash <<= 2
	hash += facilityHash(f.elevatorFloor - 1)
	return hash
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
		rtgs:          make([]int, 10),
		chips:         make([]int, 10),
	}
	lines := utils.ReadInputLines(filename)
	elements := map[string]int{}
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
			elementID := 0
			if id, found := elements[parts[0]]; found {
				elementID = id
			} else {
				elementID = len(elements)
				elements[parts[0]] = elementID
			}
			switch parts[1] {
			case "generator":
				facility.rtgs[elementID] = floor
			case "microchip":
				facility.chips[elementID] = floor
			}
		}
	}
	facility.rtgs = facility.rtgs[:len(elements)]
	facility.chips = facility.chips[:len(elements)]
	return facility
}

var benchmark = false
