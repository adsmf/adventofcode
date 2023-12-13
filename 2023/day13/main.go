package main

import (
	_ "embed"
	"fmt"
	"math/bits"

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
	lineBuffer := make([]string, 0, 18)
	utils.EachSectionMB(input, "\n\n", func(sectionNum int, section string) (done bool) {
		mirror := findReflection(section, lineBuffer, 0)
		totalP1 += mirror
		copy := []byte(section)
		flip := func(offset int) bool {
			switch copy[offset] {
			case '.':
				copy[offset] = '#'
				return true
			case '#':
				copy[offset] = '.'
				return true
			}
			return false
		}
		for i := 0; i < len(section); i++ {
			alt := 0
			if flip(i) {
				alt = findReflection(string(copy), lineBuffer, mirror)

				if alt > 0 {
					totalP2 += alt
					break
				}
				flip(i)
			}
		}
		return false
	})
	return totalP1, totalP2
}

func findReflection(section string, lineBuffer []string, ignore int) int {
	sectionLines := lineBuffer[0:0]
	utils.EachLine(section, func(index int, line string) (done bool) {
		sectionLines = append(sectionLines, line)
		return false
	})
	possMirrorCols := 0
	notMirrorCols := 0
	ignoreRow := 0
	if ignore < 100 {
		notMirrorCols = 1 << ignore
	} else {
		ignoreRow = ignore / 100
	}
	for row, line := range sectionLines {
		for x := 1; x < len(line); x++ {
			if notMirrorCols&(1<<x) > 0 {
				continue
			}
			match := true
			for i := 0; i < x && (i+x) < len(line); i++ {
				if line[x+i] != line[x-i-1] {
					match = false
					break
				}
			}
			if match {
				possMirrorCols |= 1 << x
			} else {
				notMirrorCols |= 1 << x
			}
		}
		if row == 0 || row == ignoreRow {
			continue
		}
		valid := true
		for i := 0; i < row && (row+i) < len(sectionLines); i++ {
			if sectionLines[row+i] != sectionLines[row-i-1] {
				valid = false
				break
			}
		}
		if valid {
			return 100 * row
		}
	}
	mirrorCols := possMirrorCols & (^notMirrorCols)
	if mirrorCols > 0 {
		return bits.Len(uint(mirrorCols)) - 1
	}
	return 0
}

var benchmark = false
