package main

import (
	_ "embed"
	"fmt"
	"math/bits"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	board := loadData()
	p1 := part1(board)
	p2 := part2(board)
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1(board patchboard) int {
	count := 0
	for _, display := range board.displays {
		for _, attempt := range display.rendered {
			switch bits.OnesCount(uint(attempt)) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	return count
}

func part2(board patchboard) int {
	total := 0
	idealSevenSegs := make([]sevenSeg, 10)
	for i := 0; i < 10; i++ {
		idealSevenSegs[i] = toSevenSeg(idealDigitSegmentStrings[i])
	}
	idealUses := countSegmentUses(idealSevenSegs)
	segmentLookup := make(map[int]int, 10)
	for i, segments := range idealDigitSegmentStrings {
		segmentLookup[int(toSevenSeg(segments))] = i
	}
	for _, display := range board.displays {
		total += solveDisplay(idealUses, segmentLookup, display)
	}
	return total
}

type segmentUseCount map[int]useHash
type segmentUseMap map[useHash]int
type useHash uint16

func countSegmentUses(attempts []sevenSeg) segmentUseMap {
	segUses := make(segmentUseCount, 7)
	for seg := 0; seg < 7; seg++ {
		segMask := sevenSeg(1 << seg)
		counts := [8]int{}
		for _, attempt := range attempts {
			if attempt&segMask > 0 {
				counts[bits.OnesCount(uint(attempt))]++
			}
		}
		uses := useHash(0)
		for count := 2; count < 8; count++ {
			uses |= useHash(counts[count] << (count * 2))
		}
		segUses[seg] = uses
	}

	useMap := make(segmentUseMap, 7)
	for ch, uses := range segUses {
		useMap[uses] = ch
	}

	return useMap
}

func solveDisplay(idealUses segmentUseMap, segmentLookup map[int]int, disp display) int {
	displayUseCount := countSegmentUses(disp.attempts)
	segmentMap := make(map[int]int, 7)
	for uses, ch := range displayUseCount {
		segmentMap[ch] = idealUses[uses]
	}

	displayNum := 0
	for _, digit := range disp.rendered {
		decoded := 0
		for ch := 0; ch < 10; ch++ {
			if (1<<ch)&digit > 0 {
				decoded |= 1 << segmentMap[ch]
			}
		}
		displayNum = displayNum*10 + segmentLookup[decoded]
	}
	return displayNum
}

func loadData() patchboard {
	lines := strings.Split(input, "\n")
	displays := make([]display, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		attempts := strings.Split(parts[0], " ")
		rendered := strings.Split(parts[1], " ")
		attemptsSeg := make([]sevenSeg, len(attempts))
		for i := 0; i < len(attempts); i++ {
			attemptsSeg[i] = toSevenSeg(attempts[i])
		}
		renderedSeg := make([]sevenSeg, len(rendered))
		for i := 0; i < len(rendered); i++ {
			renderedSeg[i] = toSevenSeg(rendered[i])
		}
		display := display{attempts: attemptsSeg, rendered: renderedSeg}
		displays = append(displays, display)
	}
	return patchboard{
		displays: displays,
	}
}

type patchboard struct {
	displays []display
}

type display struct {
	attempts []sevenSeg
	rendered []sevenSeg
}

type sevenSeg uint8

func toSevenSeg(input string) sevenSeg {
	res := sevenSeg(0)
	for _, char := range []byte(input) {
		res |= 1 << (sevenSeg(char) - 'a')
	}
	return res
}

var idealDigitSegmentStrings = []string{
	"abcefg",  // "abc efg", // 0
	"cf",      // "  c  f ", // 1
	"acdeg",   // "a cde g", // 2
	"acdfg",   // "a cd fg", // 3
	"bcdf",    // " bcd f ", // 4
	"abdfg",   // "ab d fg", // 5
	"abdefg",  // "ab defg", // 6
	"acf",     // "a c  f ", // 7
	"abcdefg", // "abcdefg", // 8
	"abcdfg",  // "abcd fg", // 9
}

var benchmark = false
