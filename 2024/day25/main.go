package main

import (
	_ "embed"
	"fmt"

	"github.com/adsmf/adventofcode/utils"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func part1() int {
	locks := make([][5]int, 0, 250)
	keys := make([][5]int, 0, 250)
	utils.EachSectionMB(input, "\n\n", func(item int, section string) (done bool) {
		checkChar := section[0]
		heights := [5]int{}
		utils.EachLine(section, func(y int, line string) (done bool) {
			for x, ch := range line {
				if ch == rune(checkChar) {
					heights[x] = y
				}
			}
			return
		})
		if checkChar == '#' {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
		return
	})
	p1 := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for pin, height := range lock {
				if key[pin] < height {
					fits = false
					break
				}
			}
			if fits {
				p1++
			}
		}
	}
	return p1
}

var benchmark = false
