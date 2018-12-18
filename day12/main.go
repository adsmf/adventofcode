package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/adsmf/adventofcode2018/utils"
)

func main() {
	initial, mapping := loadData("input.txt")
	_, sum := runSim(initial, mapping, 20)
	fmt.Printf("Sum: %d\n", sum)
}

func runSim(pots string, mapping []bool, ticks int) (string, int) {
	padding := "00"
	padLen := len(padding)

	initialLen := len(pots)
	for tick := 0; tick < ticks; tick++ {
		pots = fmt.Sprintf("%s%s%s", padding, pots, padding)
		paddedPots := fmt.Sprintf("%s%s%s", padding, pots, padding)
		next := ""
		for curPos := 0; curPos < len(pots); curPos++ {

			part := paddedPots[curPos : curPos+5]
			partInt, _ := strconv.ParseUint(part, 2, 5)
			if mapping[partInt] {
				next = fmt.Sprintf("%s1", next)
			} else {
				next = fmt.Sprintf("%s0", next)
			}
		}
		next = fmt.Sprintf("%s%s", next, padding)
		pots = next
	}
	value := 0
	for curPot := 0; curPot < len(pots); curPot++ {
		if pots[curPot] == byte("1"[0]) {
			value += curPot - ticks*padLen
		}
	}
	return pots[ticks*padLen : ticks*padLen+initialLen], value
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
