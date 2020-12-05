package main

import (
	"fmt"
	"io/ioutil"
)

var benchmark = false

func main() {
	part1, part2 := loadBitwise("input.txt")

	if !benchmark {
		fmt.Printf("Part 1: %d\n", part1)
		fmt.Printf("Part 2: %d\n", part2)
	}
}

func loadBitwise(filename string) (int, int) {
	inputBytes, _ := ioutil.ReadFile(filename)
	min, max, total := 1<<17, 0, 0
	for cursor := 0; cursor <= len(inputBytes)-11; cursor += 11 {
		line := inputBytes[cursor : cursor+10]
		pass := 0
		for i := 0; i <= 9; i++ {
			pass |= (int(^line[i]&0x4) >> 2) << (9 - i)
		}
		total += pass
		if pass < min {
			min = pass
		}
		if pass > max {
			max = pass
		}
	}
	total -= (max * (max + 1) / 2) - (min * (min - 1) / 2)
	return max, total * -1
}
