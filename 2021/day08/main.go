package main

import (
	_ "embed"
	"fmt"
	"sort"
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
			switch len(attempt) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	return count
}

func part2(board patchboard) int {
	total := 0
	idealUses := countSegmentUses(idealDigitSegments)
	segmentLookup := map[string]int{}
	for i, segments := range idealDigitSegments {
		segmentLookup[segments] = i
	}
	for _, display := range board.displays {
		total += solveDisplay(idealUses, segmentLookup, display)
	}
	return total
}

type segmentUseCount map[rune]useHash
type segmentUseMap map[useHash]rune
type useHash string

func countSegmentUses(attempts []string) segmentUseMap {
	charUses := make(segmentUseCount, 10)
	for ch := 'a'; ch <= 'g'; ch++ {
		uses := strings.Builder{}
		for _, attempt := range attempts {
			for _, attCh := range attempt {
				if attCh == ch {
					uses.WriteByte(byte(len(attempt) + '0'))
				}
			}
		}
		charUses[ch] = useHash(sortString(uses.String()))
	}

	useMap := make(segmentUseMap, 10)
	for ch, uses := range charUses {
		useMap[uses] = ch
	}

	return useMap
}

func solveDisplay(idealUses segmentUseMap, segmentLookup map[string]int, disp display) int {
	displayUseCount := countSegmentUses(disp.attempts)
	segmentMap := map[rune]rune{}
	for uses, ch := range displayUseCount {
		segmentMap[ch] = idealUses[uses]
	}

	displayNum := 0
	for _, digit := range disp.rendered {
		decodedSB := strings.Builder{}
		for _, ch := range digit {
			decodedSB.WriteByte(byte(segmentMap[ch]))
		}
		decoded := sortString(decodedSB.String())
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
		display := display{attempts: attempts, rendered: rendered}
		displays = append(displays, display)
	}
	return patchboard{
		displays: displays,
	}
}

func sortString(input string) string {
	stringSlice := []byte(input)
	sort.Sort(sortByteSlice(stringSlice))
	return string(stringSlice)
}

type sortByteSlice []byte

func (s sortByteSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByteSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s sortByteSlice) Len() int           { return len(s) }

type patchboard struct {
	displays []display
}

type display struct {
	attempts []string
	rendered []string
}

var idealDigitSegments = []string{
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
