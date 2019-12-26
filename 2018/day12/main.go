package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/adsmf/adventofcode/utils"
)

func main() {
	initial, mapping := loadData("input.txt")
	sum := runSim(initial, mapping, 50000000000)
	fmt.Printf("Sum: %d\n", sum)
}

type potState struct {
	value    int
	occupied uint
}

func runSim(potString string, mapping []bool, ticks int) int {
	var value int
	potsA := list.New()
	potsB := list.New()
	for idx, pot := range potString {
		potVal, err := strconv.Atoi(string(pot))
		if err != nil {
			panic(err)
		}
		newPot := potState{
			value:    idx,
			occupied: uint(potVal),
		}
		potsA.PushBack(newPot)
		potsB.PushBack(newPot)
	}

	potsCurrent := potsA
	potsNext := potsB

	start := time.Now()

	values := []int{}
	for tick := 0; tick < ticks; tick++ {
		// Swap buffers
		potsTemp := potsCurrent
		potsCurrent = potsNext
		potsNext = potsTemp

		// Pad stuff
		firstValue := potsCurrent.Front().Value.(potState).value
		lastValue := potsCurrent.Back().Value.(potState).value
		for padIdx := 0; padIdx < 2; padIdx++ {
			newFront := potState{
				value:    firstValue - padIdx - 1,
				occupied: 0,
			}
			newBack := potState{
				value:    lastValue + padIdx + 1,
				occupied: 0,
			}
			potsCurrent.PushBack(newBack)
			potsCurrent.PushFront(newFront)
			potsNext.PushBack(newBack)
			potsNext.PushFront(newFront)
		}

		var inPosM1, inPosM2 *list.Element
		var inPosP1, inPosP2 *list.Element
		value = 0
		for inPos, outPos := potsCurrent.Front(), potsNext.Front(); inPos != nil; inPos, outPos = inPos.Next(), outPos.Next() {

			var ctx uint
			ctx = inPos.Value.(potState).occupied << 2

			inPosM1 = inPos.Prev()
			if inPosM1 != nil {
				ctx += inPosM1.Value.(potState).occupied << 3
				inPosM2 = inPosM1.Prev()
				if inPosM2 != nil {
					ctx += inPosM2.Value.(potState).occupied << 4
				}
			}

			inPosP1 = inPos.Next()
			if inPosP1 != nil {
				ctx += inPosP1.Value.(potState).occupied << 1
				inPosP2 = inPosP1.Next()
				if inPosP2 != nil {
					ctx += inPosP2.Value.(potState).occupied
				}
			}

			potValue := inPos.Value.(potState).value
			if mapping[ctx] {
				outPos.Value = potState{
					value:    potValue,
					occupied: 1,
				}
				value += potValue
			} else {
				outPos.Value = potState{
					value:    potValue,
					occupied: 0,
				}
			}
		}
		values = append(values, value)

		if tick > 0 && tick%1000 == 0 {
			percentComplete := float64(tick) / float64(ticks)
			elapsed := time.Since(start)
			fmt.Printf("%f%% -- current %d @ tick %d -- elapsed %v \n", percentComplete, value, tick, elapsed)
		}

		for cycleSize := 2; cycleSize <= int(len(values)/2); cycleSize++ {
			isCycle := true
			checkVal := values[tick] - values[tick-cycleSize]
			for dt := 0; dt < cycleSize*2; dt++ {
				if values[tick-dt]-values[tick-cycleSize-dt] != checkVal {
					isCycle = false
					break
				}
			}
			if isCycle {
				fmt.Printf("Cycle detected at tick %d, delta %d every %d!\n", tick, checkVal, cycleSize)
				for dt := 0; dt < cycleSize; dt++ {
					baseTick := (tick - dt)
					remainingTicks := ticks - baseTick
					if remainingTicks%cycleSize == 0 {
						baseValue := values[baseTick-1]
						remainingCycles := remainingTicks / cycleSize
						predictedValue := remainingCycles*checkVal + baseValue
						fmt.Printf("\tBase: %d from tick %d\n", baseValue, baseTick)
						fmt.Printf("\tRemaining cycles: (%d - %d)/%d => %d\n", ticks, baseTick, cycleSize, remainingCycles)
						fmt.Printf("\tPredicted end: %d + %d*%d => %d\n", baseValue, remainingCycles, checkVal, predictedValue)
						return predictedValue
					}
				}
			}
		}
	}
	if ticks == 0 {
		value = 0
		for pot := potsNext.Front(); pot != nil; pot = pot.Next() {
			state := pot.Value.(potState)
			if state.occupied > 0 {
				value += state.value
			}
		}
	}
	return value
}

func loadData(filename string) (string, []bool) {
	lines := utils.ReadInputLines(filename)

	initial := strings.TrimPrefix(lines[0], "initial state: ")
	initial = strings.Replace(initial, ".", "0", -1)
	initial = strings.Replace(initial, "#", "1", -1)
	lines = lines[2:]

	patterns := make([]bool, 32)
	for _, line := range lines {
		line = strings.Replace(line, ".", "0", -1)
		line = strings.Replace(line, "#", "1", -1)
		lineParts := strings.Split(line, " => ")
		pattern, _ := strconv.ParseUint(lineParts[0], 2, 5)
		result, _ := strconv.Atoi(lineParts[1])
		patterns[byte(pattern)] = (result == 1)
	}

	return initial, patterns
}
